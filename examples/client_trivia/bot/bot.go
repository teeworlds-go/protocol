package bot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

type Bot struct {
	sf      ServerFactory
	qf      QuestionFactory
	backoff BackoffFunc
}

func NewBot() *Bot {
	return &Bot{
		qf:      NewQuestionFactory(),
		sf:      NewServerFactory(),
		backoff: newDefaultBackoffPolicy(time.Second, 5*time.Minute),
	}
}

func (b *Bot) Run(ctx context.Context) error {

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// continue
		}

		_ = b.retry(ctx, func() (cont bool, err error) {
			server, err := b.sf.Next(ctx)
			if err != nil {
				return true, err
			}
			address := server.Addresses[0]
			parts := strings.SplitN(address, ":", 3)
			if len(parts) != 3 {
				return false, fmt.Errorf("invalid address: %s", address)
			}

			ip := strings.TrimLeft(parts[1], "/")
			port, err := strconv.ParseInt(parts[2], 10, 16)
			if err != nil {
				return false, fmt.Errorf("invalid port: %s", address)
			}

			b.start(ctx, ip, int(port))

			return false, nil
		})
	}
}

func (b *Bot) start(ctx context.Context, ip string, port int) error {
	ctx, cancelCause := context.WithCancelCause(ctx)
	defer cancelCause(nil)

	client := teeworlds7.NewClient()

	client.OnUnknown(func(msg *messages7.Unknown, defaultAction teeworlds7.DefaultAction) error {
		return nil
	})
	client.OnMotd(func(msg *messages7.SvMotd, defaultAction teeworlds7.DefaultAction) error {
		return nil
	})

	state := NewState(&b.qf)
	client.OnChat(b.statefulOnChat(ctx, client, state))
	client.OnChat(func(msg *messages7.SvChat, defaultAction teeworlds7.DefaultAction) error {
		if msg.ClientId < 0 {
			return nil
		}

		switch msg.Message {
		case "!leave":
			cancelCause(errors.New("leave command !leave received"))
			return nil
		case "!help", "!h":
			return client.SendChat("Available commands: !trivia, !top, !score and !leave")
		}
		if msg.Message == "!leave" {
			cancelCause(errors.New("leave command !leave received"))
		}

		return nil
	})

	client.OnDisconnect(func(msg *messages7.CtrlClose, defaultAction teeworlds7.DefaultAction) error {
		cancelCause(errors.New("disconnected"))
		return defaultAction()
	})

	client.OnClientInfo(func(msg *messages7.SvClientInfo, defaultAction teeworlds7.DefaultAction) error {
		client.Game.Players[msg.ClientId].Info = *msg
		if msg.Local {
			client.LocalClientId = msg.ClientId
		}
		return nil
	})

	client.OnServerInfo(func(msg *messages7.ServerInfo, defaultAction teeworlds7.DefaultAction) error {
		return nil
	})

	client.OnBroadcast(func(msg *messages7.SvBroadcast, defaultAction teeworlds7.DefaultAction) error {
		return nil
	})

	slog.Info(fmt.Sprintf("connecting to server: %s:%d", ip, port), "ip", ip, "port", port)
	defer func() {
		slog.Info(fmt.Sprintf("disconnected from server: %s:%d", ip, port), "ip", ip, "port", port)
	}()
	return client.ConnectContext(ctx, ip, int(port))
}

func (b *Bot) statefulOnChat(ctx context.Context, client *teeworlds7.Client, state *State) func(msg *messages7.SvChat, defaultAction teeworlds7.DefaultAction) error {
	return func(msg *messages7.SvChat, _ teeworlds7.DefaultAction) error {
		if msg.ClientId < 0 {
			// world chat, skip that
			return nil
		}

		if msg.Message == "!trivia" {
			chatLine, err := state.Start(ctx)
			if err != nil {
				return client.SendChat(fmt.Sprintf("Failed to start trivia: %v", err))
			}
			err = client.SendChat(chatLine)
			if err != nil {
				return err
			}

			return nil
		} else if msg.Message == "!top" {
			return client.SendChat(state.Top())
		} else if msg.Message == "!score" {
			playerName := client.Game.Players[msg.ClientId].Info.Name
			return client.SendChat(state.Score(playerName))
		}

		if !state.Running() {
			return nil
		}

		playerName := client.Game.Players[msg.ClientId].Info.Name
		correctAnswer, correct := state.Answer(playerName, msg.Message)
		if !correct {
			return nil
		}
		// player answered correctly
		return client.SendChat(fmt.Sprintf("%s answered correctly: %s", playerName, correctAnswer))
	}
}

func (b *Bot) retry(ctx context.Context, f func() (cont bool, err error)) error {
	// fast path
	cont, err := f()
	if err != nil {
		return err
	}
	if !cont {
		return nil
	}

	// continue second try with backoff timer overhead
	var (
		retry   = 1
		timer   = time.NewTimer(b.backoff(retry))
		drained = false
	)
	defer closeTimer(timer, &drained)

	for {

		select {
		case <-timer.C:
			// at this point we know that the timer channel has been drained
			drained = true

			// try again
			cont, err = f()
			if err != nil {
				return err
			}
			if !cont {
				return nil
			}

			retry++
			resetTimer(timer, b.backoff(retry), &drained)

		case <-ctx.Done():
			return errors.Join(err, ctx.Err())
		}
	}
}

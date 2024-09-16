package teeworlds7

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
)

const (
	maxPacksize = 1400
)

var (
	// ErrProcessPacketFailed is returned when processing a packet failed
	// This error can be checked in your error handling callback with errors.Is
	ErrProcessPacketFailed = errors.New("failed to process packet")

	// ErrUnpackPacketFailed is returned when unpacking a packet failed
	// This error can be checked in your error handling callback with errors.Is
	ErrUnpackPacketFailed = errors.New("failed to unpack packet")

	// ErrProcessGameTickFailed is returned when processing a game tick failed
	// This error can be checked in your error handling callback with errors.Is
	ErrProcessGameTickFailed = errors.New("failed to process game tick")
)

func readNetwork(ctx context.Context, cancelCause context.CancelCauseFunc, wg *sync.WaitGroup, ch chan<- []byte, conn net.Conn) {
	defer wg.Done()
	slog.Debug("starting reader goroutine...")
	defer slog.Debug("reader goroutine stopped")

	buf := make([]byte, maxPacksize)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			cancelCause(err)
			return
		}

		if n == 0 {
			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}

		packet := make([]byte, n)
		copy(packet, buf[:n])
		select {
		case ch <- packet:
		case <-ctx.Done():
			return
		}
	}
}

func (client *Client) Connect(serverIp string, serverPort int) error {
	return client.ConnectContext(context.Background(), serverIp, serverPort)
}

func (client *Client) ConnectContext(ctx context.Context, serverIp string, serverPort int) (err error) {
	ctx, cancelCause := context.WithCancelCause(ctx)
	defer cancelCause(nil) // always cancel

	// wait for the reader goroutine to finish execution, before leaving this function scope
	var wg sync.WaitGroup
	defer wg.Wait()

	ch := make(chan []byte, maxPacksize)
	var d net.Dialer
	conn, err := d.DialContext(ctx, "udp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		return fmt.Errorf("failed to connect to server: %s:%d: %w", serverIp, serverPort, err)
	}
	client.Conn = conn
	defer func() {
		// only the first cancelation cause is relevant
		// subsequent cancelations will be ignored
		// this one might be a subsequent cancelation
		cancelCause(err)

		ctxErr := context.Cause(ctx)
		if ctxErr != nil {
			err = ctxErr
			return
		}

		ctxErr = ctx.Err()
		if ctxErr != nil && !errors.Is(ctxErr, context.Canceled) {
			err = ctxErr
			return
		}

		// close connection after error handling in order not to
		// hide the actual cause of the error
		closeErr := client.Conn.Close()
		if closeErr != nil {
			slog.Error("failed to close connection", "error", closeErr)
		}
	}()

	client.Session = protocol7.NewSession()
	client.Game.Players = make([]Player, network7.MaxClients)

	wg.Add(1)
	go readNetwork(ctx, cancelCause, &wg, ch, conn)

	err = client.SendPacket(client.Session.CtrlToken())
	if err != nil {
		return fmt.Errorf("failed to send token: %w", err)
	}

	// TODO: do we really need a non blocking network read?
	//       if not remove the channel, the sleep and the select statement
	//       if yes also offer an OnTick callback to the user, and also do keepalives and resends

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				panic("processing channel closed unexpectedly")
			}
			packet := &protocol7.Packet{}
			err := packet.Unpack(msg)
			if err != nil {
				err = client.handleInternalError(fmt.Errorf("%w: %w", ErrUnpackPacketFailed, err))
				if err != nil {
					return err
				}
				// there was an actual unpacking error, which is why we cannot proceed the execution
				// processing is not possible
				continue
			}
			err = client.processPacket(packet)
			if err != nil {
				err = client.handleInternalError(fmt.Errorf("%w: %w", ErrProcessPacketFailed, err))
				if err != nil {
					return err
				}
			}
		case <-ticker.C:
			err = client.gameTick()
			if err != nil {
				err = client.handleInternalError(fmt.Errorf("%w: %w", ErrProcessGameTickFailed, err))
				if err != nil {
					return err
				}
			}
		case <-ctx.Done():
			return nil
		}
	}

}

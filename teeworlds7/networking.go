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
	defer func() {
		// can be called only once
		// in case this cause is second, it will not override
		// the initial cancelation cause
		cancelCause(nil)

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
	}()

	ch := make(chan []byte, maxPacksize)
	var d net.Dialer
	conn, err := d.DialContext(ctx, "udp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		return fmt.Errorf("failed to connect to server: %s:%d: %w", serverIp, serverPort, err)
	}
	client.Conn = conn
	defer func() {
		closeErr := client.Conn.Close()
		if closeErr != nil {
			slog.Error("failed to close connection", "error", closeErr)
		}
	}()

	client.Session = protocol7.NewSession()
	client.Game.Players = make([]Player, network7.MaxClients)

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait() // wait for the reader goroutine to finish execution, before leaving this function scope

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
				return fmt.Errorf("failed to unpack packet: %w", err)
			}
			err = client.processPacket(packet)
			if err != nil {
				return fmt.Errorf("failed to process packet: %w", err)
			}
		case <-ticker.C:
			err = client.gameTick()
			if err != nil {
				return fmt.Errorf("failed to process game tick: %w", err)
			}
		case <-ctx.Done():
			return nil
		}
	}

}

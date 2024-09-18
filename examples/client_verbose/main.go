package main

// TODO: split this up into multiple examples
//       the verbose mode should not have a render loop
//       the verbose mode should not implement disconnect

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	client := teeworlds7.NewClient()
	client.Name = "nameless tee"

	client.OnAccept(func(msg *messages7.CtrlAccept, defaultAction teeworlds7.DefaultAction) error {
		// respond with the next message to establish a connection
		err := defaultAction()
		if err != nil {
			return err
		}

		fmt.Println("got accept message")
		return nil
	})

	// read incoming traffic
	// you can also alter packet here before it will be processed by the internal state machine
	//
	// return false to drop the packet
	client.OnPacket(func(packet *protocol7.Packet) bool {
		fmt.Printf("got packet with %d messages\n", len(packet.Messages))
		return true
	})

	// inspect outgoing traffic
	// you can also alter packet here before it will be sent to the server
	//
	// return false to drop the packet
	client.OnSendPacket(func(packet *protocol7.Packet) bool {
		fmt.Printf("sending packet with %d messages\n", len(packet.Messages))
		return true
	})

	client.OnChat(func(msg *messages7.SvChat, defaultAction teeworlds7.DefaultAction) error {
		// the default action prints the chat message to the console
		// if this is not called and you don't print it your self the chat will not be visible
		err := defaultAction()
		if err != nil {
			return err
		}

		// additional custom chat print
		fmt.Printf("%d %s\n", msg.ClientId, msg.Message)
		return nil
	})

	// this is matching the default behavior
	client.OnDisconnect(func(msg *messages7.CtrlClose, defaultAction teeworlds7.DefaultAction) error {
		fmt.Printf("disconnected (%s)\n", msg.Reason)

		// cancel client context on being kicked
		// this will tell the client to
		// gracefully shutdown
		client.CancelCause(teeworlds7.DisconnectError{Reason: msg.Reason})
		return nil
	})

	// if you do not implement OnError it will throw on error
	client.OnError(func(err error) error {
		fmt.Println(err)

		// return nil to stop the error propagation
		// return the error in order to quit the client execution
		return nil
	})

	err := client.ConnectContext(ctx, "127.0.0.1", 8303)
	if err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println(err)
		os.Exit(1)
	}

	// channel was closed, no error
	fmt.Println("graceful shutdown")
}

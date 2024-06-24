package main

// TODO: split this up into multiple examples
//       the verbose mode should not have a render loop
//       the verbose mode should not implement disconnect

import (
	"fmt"
	"os"
	"time"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
	"github.com/teeworlds-go/go-teeworlds-protocol/teeworlds7"
)

func main() {
	client := teeworlds7.Client{
		Name: "nameless tee",
	}

	client.OnAccept(func(msg *messages7.CtrlAccept, defaultAction teeworlds7.DefaultAction) {
		// respond with the next message to establish a connection
		defaultAction()

		fmt.Println("got accept message")
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
	client.OnSend(func(packet *protocol7.Packet) bool {
		fmt.Printf("sending packet with %d messages\n", len(packet.Messages))
		return true
	})

	client.OnChat(func(msg *messages7.SvChat, defaultAction teeworlds7.DefaultAction) {
		// the default action prints the chat message to the console
		// if this is not called and you don't print it your self the chat will not be visible
		defaultAction()

		// additional custom chat print
		fmt.Printf("%d %s\n", msg.ClientId, msg.Message)
	})

	// this is matching the default behavior
	client.OnDisconnect(func(msg *messages7.CtrlClose, defaultAction teeworlds7.DefaultAction) {
		fmt.Printf("disconnected (%s)\n", msg.Reason)
		os.Exit(0)
	})

	// if you do not implement OnError it will throw on error
	client.OnError(func(err error) {
		fmt.Print(err)
	})

	go func() {
		client.Connect("127.0.0.1", 8303)
	}()

	for {
		// render loop
		time.Sleep(100_000_000)
	}
}

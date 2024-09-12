package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	client := teeworlds7.NewClient()
	client.Name = "nameless tee"

	client.OnPacket(func(packet *protocol7.Packet) bool {
		fmt.Printf("received packet with %d message(s)\n", len(packet.Messages))

		for idx, msg := range packet.Messages {
			fmt.Printf("%-3d %-8s %d\n", idx, msg.MsgType(), msg.MsgId())

		}
		return true
	})

	err := client.ConnectContext(ctx, "127.0.0.1", 8303)
	if err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println("failed to connect:", err)
		os.Exit(1)
	}

	fmt.Println("shutdown")
}

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	client := teeworlds7.NewClient()

	client.OnChat(func(msg *messages7.SvChat, _ teeworlds7.DefaultAction) error {
		if msg.Mode != network7.ChatAll || msg.ClientId < 0 {
			// ignore if not from users and not in public chat
			return nil
		}

		// echo user messages
		return client.SendChat(msg.Message)
	})

	err := client.ConnectContext(ctx, "127.0.0.1", 8303)
	if err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("graceful shutdown")
}

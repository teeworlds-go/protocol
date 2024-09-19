package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/teeworlds-go/protocol/examples/client_trivia/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	const (
		connections = 1 // number of concurrent instances
	)

	var (
		bot = bot.NewBot()
		wg  sync.WaitGroup
	)

	for i := 0; i < connections; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := bot.Run(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				slog.Error(fmt.Sprintf("bot routine %d stopped: %v", id, err))
			} else {
				slog.Info(fmt.Sprintf("bot routine %d stopped", id))
			}
		}(i)
	}

	wg.Wait()
	err := context.Cause(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		os.Exit(1)
	}
}

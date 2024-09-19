package bot_test

import (
	"context"
	"sync"
	"testing"

	"github.com/teeworlds-go/protocol/examples/client_trivia/bot"
)

func TestServerFactory(t *testing.T) {
	sf := bot.NewServerFactory()
	ctx := context.Background()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := sf.Next(ctx)
			if err != nil {
				t.Error(err)
			}

		}()
	}

	wg.Wait()
}

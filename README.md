# teeworlds 0.7 protocol library for go

A client side network protocol implementation of the game teeworlds.

## high level api for ease of use

The package **teeworlds7** implements a high level client library. Designed for ease of use.

```go
package main

import (
	"fmt"
	"os"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/snapshot7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

func main() {
	client := teeworlds7.NewClient()
	client.Name = "nameless tee"

	// Register your callback for incoming chat messages
	// For a full list of all callbacks see: https://github.com/teeworlds-go/protocol/tree/master/teeworlds7/user_hooks.go
	client.OnChat(func(msg *messages7.SvChat, defaultAction teeworlds7.DefaultAction) error {
		// the default action prints the chat message to the console
		// if this is not called and you don't print it your self the chat will not be visible
		err := defaultAction()
		if err != nil {
			return err
		}

		if msg.Message == "!ping" {
			// Send reply in chat using the SendChat() action
			// For a full list of all actions see: https://github.com/teeworlds-go/protocol/tree/master/teeworlds7/user_actions.go
			return client.SendChat("pong")
		}
		return nil
	})

	client.OnSnapshot(func(snap *snapshot7.Snapshot, defaultAction teeworlds7.DefaultAction) error {
		fmt.Printf("got snap with %d items\n", len(snap.Items))

		for _, character := range client.Game.Snap.Characters {
			fmt.Printf("  got tee at %.2f %.2f\n", float32(character.X)/32.0, float32(character.Y)/32.0)
		}

		char, found, err := client.SnapFindCharacter(client.LocalClientId)
		if err != nil {
			return err
		}
		if !found {
			return nil
		}
		fmt.Printf("  we are at %d %d\n", char.X/32, char.Y/32)
		client.Right()
		return nil
	})

	err := client.Connect("127.0.0.1", 8303)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

Example usages:

- [client_verbose](./examples/client_verbose/) a verbose client show casing the easy to use high level api
- [client_simple](./examples/client_simple/) a simple client showing the basic use of the high level api

## low level api for power users

The packages **chunk7, messages7, network7, packer, protocol7** Implement the low level 0.7 teeworlds protocol. Use them if you want to build something advanced such as a custom proxy.

## projects using protocol

- [MITM teeworlds proxy](https://github.com/teeworlds-go/proxy)
- [goofworlds gui client](https://github.com/teeworlds-go/goofworlds)

# go-teeworlds-protocol

# WARNING! NOT READY TO BE USED YET! Apis might change. Packages and repository might be renamed!

A client side network protocol implementation of the game teeworlds.

## high level api for ease of use

The package **teeworlds7** implements a high level client library. Designed for ease of use.

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/teeworlds7"
)

func main() {
	client := teeworlds7.Client{Name: "nameless tee"}
	
	// Register your callback for incoming chat messages
	// For a full list of all callbacks see: https://github.com/teeworlds-go/go-teeworlds-protocol/tree/master/teeworlds7/user_hooks.go
	client.OnChat(func(msg *messages7.SvChat, defaultAction teeworlds7.DefaultAction) {
		// the default action prints the chat message to the console
		// if this is not called and you don't print it your self the chat will not be visible
		defaultAction()

		if msg.Message == "!ping" {
			// Send reply in chat using the SendChat() action
			// For a full list of all actions see: https://github.com/teeworlds-go/go-teeworlds-protocol/tree/master/teeworlds7/user_actions.go
			client.SendChat("pong")
		}
	})

	client.Connect("127.0.0.1", 8303)
}
```


Example usages:
- [client_verbose](./examples/client_verbose/) a verbose client show casing the easy to use high level api

## low level api for power users

The packages **chunk7, messages7, network7, packer, protocol7** Implement the low level 0.7 teeworlds protocol. Use them if you want to build something advanced such as a custom proxy.

## projects using go-teeworlds-protocol

- [MITM teeworlds proxy](https://github.com/teeworlds-go/proxy)


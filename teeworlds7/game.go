package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

func (client *Client) processGame(netMsg messages7.NetMessage, response *protocol7.Packet) bool {
	switch msg := netMsg.(type) {
	case *messages7.SvMotd:
		defaultAction := func() {
			if msg.Message != "" {
				fmt.Printf("[motd] %s\n", msg.Message)
			}
		}
		if client.Callbacks.GameSvMotd == nil {
			defaultAction()
		} else {
			client.Callbacks.GameSvMotd(msg, defaultAction)
		}
	case *messages7.SvBroadcast:
		defaultAction := func() {
			fmt.Printf("[broadcast] %s\n", msg.Message)
		}
		if client.Callbacks.GameSvBroadcast == nil {
			defaultAction()
		} else {
			client.Callbacks.GameSvBroadcast(msg, defaultAction)
		}
	case *messages7.SvChat:
		defaultAction := func() {
			if msg.ClientId < 0 || msg.ClientId > network7.MaxClients {
				fmt.Printf("[chat] *** %s\n", msg.Message)
				return
			}
			name := client.Game.Players[msg.ClientId].Info.Name
			fmt.Printf("[chat] <%s> %s\n", name, msg.Message)
		}
		if client.Callbacks.GameSvChat == nil {
			defaultAction()
		} else {
			client.Callbacks.GameSvChat(msg, defaultAction)
		}
	case *messages7.SvClientInfo:
		defaultAction := func() {
			client.Game.Players[msg.ClientId].Info = *msg
			fmt.Printf("got client info id=%d name=%s\n", msg.ClientId, msg.Name)
		}
		if client.Callbacks.GameSvClientInfo == nil {
			defaultAction()
		} else {
			client.Callbacks.GameSvClientInfo(msg, defaultAction)
		}
	case *messages7.ReadyToEnter:
		defaultAction := func() {
			fmt.Println("got ready to enter")
			response.Messages = append(response.Messages, &messages7.EnterGame{})
		}
		if client.Callbacks.GameReadyToEnter == nil {
			defaultAction()
		} else {
			client.Callbacks.GameReadyToEnter(msg, defaultAction)
		}
	case *messages7.Unknown:
		defaultAction := func() {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown game")
		}
		if client.Callbacks.MsgUnknown == nil {
			defaultAction()
		} else {
			client.Callbacks.MsgUnknown(msg, defaultAction)
		}
	default:
		printUnknownMessage(netMsg, "unprocessed game")
		return false
	}
	return true
}

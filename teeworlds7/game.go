package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
)

func (client *Client) processGame(netMsg messages7.NetMessage, response *protocol7.Packet) bool {
	switch msg := netMsg.(type) {
	case *messages7.SvMotd:
		userMsgCallback(client.Callbacks.GameSvMotd, msg, func() {
			if msg.Message != "" {
				fmt.Printf("[motd] %s\n", msg.Message)
			}
		})
	case *messages7.SvBroadcast:
		userMsgCallback(client.Callbacks.GameSvBroadcast, msg, func() {
			fmt.Printf("[broadcast] %s\n", msg.Message)
		})
	case *messages7.SvChat:
		userMsgCallback(client.Callbacks.GameSvChat, msg, func() {
			if msg.ClientId < 0 || msg.ClientId > network7.MaxClients {
				fmt.Printf("[chat] *** %s\n", msg.Message)
				return
			}
			name := client.Game.Players[msg.ClientId].Info.Name
			fmt.Printf("[chat] <%s> %s\n", name, msg.Message)
		})
	case *messages7.SvClientInfo:
		userMsgCallback(client.Callbacks.GameSvClientInfo, msg, func() {
			client.Game.Players[msg.ClientId].Info = *msg
			fmt.Printf("got client info id=%d name=%s\n", msg.ClientId, msg.Name)
		})
	case *messages7.SvReadyToEnter:
		userMsgCallback(client.Callbacks.GameSvReadyToEnter, msg, func() {
			fmt.Println("got ready to enter")
			response.Messages = append(response.Messages, &messages7.EnterGame{})
		})
	case *messages7.Unknown:
		userMsgCallback(client.Callbacks.MsgUnknown, msg, func() {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown game")
		})
	default:
		printUnknownMessage(netMsg, "unprocessed game")
		return false
	}
	return true
}

package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
)

func (client *Client) processGame(netMsg messages7.NetMessage, response *protocol7.Packet) (process bool, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to process game message: %w", err)
		}
	}()

	switch msg := netMsg.(type) {
	case *messages7.SvMotd:
		err = userMsgCallback(client.Callbacks.GameSvMotd, msg, func() error {
			if msg.Message != "" {
				fmt.Printf("[motd] %s\n", msg.Message)
			}
			return nil
		})
	case *messages7.SvBroadcast:
		err = userMsgCallback(client.Callbacks.GameSvBroadcast, msg, func() error {
			fmt.Printf("[broadcast] %s\n", msg.Message)
			return nil
		})
	case *messages7.SvChat:
		err = userMsgCallback(client.Callbacks.GameSvChat, msg, func() error {
			if msg.ClientId < 0 || msg.ClientId > network7.MaxClients {
				fmt.Printf("[chat] *** %s\n", msg.Message)
				return nil
			}
			name := client.Game.Players[msg.ClientId].Info.Name
			fmt.Printf("[chat] <%s> %s\n", name, msg.Message)
			return nil
		})
	case *messages7.SvClientInfo:
		err = userMsgCallback(client.Callbacks.GameSvClientInfo, msg, func() error {
			client.Game.Players[msg.ClientId].Info = *msg
			if msg.Local {
				client.LocalClientId = msg.ClientId
			}
			fmt.Printf("got client info id=%d name=%s\n", msg.ClientId, msg.Name)
			return nil
		})
	case *messages7.SvReadyToEnter:
		err = userMsgCallback(client.Callbacks.GameSvReadyToEnter, msg, func() error {
			fmt.Println("got ready to enter")
			response.Messages = append(response.Messages, &messages7.EnterGame{})
			return nil
		})
	case *messages7.Unknown:
		err = userMsgCallback(client.Callbacks.MsgUnknown, msg, func() error {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown game")
			return nil
		})
	default:
		printUnknownMessage(netMsg, "unprocessed game")
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

func (client *Client) processSystem(netMsg messages7.NetMessage, response *protocol7.Packet) bool {
	switch msg := netMsg.(type) {
	case *messages7.MapChange:
		userMsgCallback(client.Callbacks.SysMapChange, msg, func() {
			fmt.Println("got map change")
			response.Messages = append(response.Messages, &messages7.Ready{})
		})
	case *messages7.MapData:
		userMsgCallback(client.Callbacks.SysMapData, msg, func() {
			fmt.Printf("got map chunk %x\n", msg.Data)
		})
	case *messages7.ServerInfo:
		userMsgCallback(client.Callbacks.SysServerInfo, msg, func() {
			fmt.Printf("connected to server with name '%s'\n", msg.Name)
		})
	case *messages7.ConReady:
		userMsgCallback(client.Callbacks.SysConReady, msg, func() {
			fmt.Println("connected, sending info")
			info := &messages7.ClStartInfo{
				Name:                  client.Name,
				Clan:                  client.Clan,
				Country:               client.Country,
				Body:                  "greensward",
				Marking:               "duodonny",
				Decoration:            "",
				Hands:                 "standard",
				Feet:                  "standard",
				Eyes:                  "standard",
				CustomColorBody:       false,
				CustomColorMarking:    false,
				CustomColorDecoration: false,
				CustomColorHands:      false,
				CustomColorFeet:       false,
				CustomColorEyes:       false,
				ColorBody:             0,
				ColorMarking:          0,
				ColorDecoration:       0,
				ColorHands:            0,
				ColorFeet:             0,
				ColorEyes:             0,
			}
			response.Messages = append(response.Messages, info)
		})
	case *messages7.Snap:
		userMsgCallback(client.Callbacks.SysSnap, msg, func() {
			response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
		})
	case *messages7.SnapSingle:
		userMsgCallback(client.Callbacks.SysSnapSingle, msg, func() {
			response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
		})
	case *messages7.SnapEmpty:
		userMsgCallback(client.Callbacks.SysSnapEmpty, msg, func() {
			response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
		})
	case *messages7.InputTiming:
		userMsgCallback(client.Callbacks.SysInputTiming, msg, func() {
			fmt.Printf("timing time left=%d\n", msg.TimeLeft)
		})
	case *messages7.RconAuthOn:
		userMsgCallback(client.Callbacks.SysRconAuthOn, msg, func() {
			fmt.Println("you are now authenticated in rcon")
		})
	case *messages7.RconAuthOff:
		userMsgCallback(client.Callbacks.SysRconAuthOff, msg, func() {
			fmt.Println("you are no longer authenticated in rcon")
		})
	case *messages7.RconLine:
		userMsgCallback(client.Callbacks.SysRconLine, msg, func() {
			fmt.Printf("[rcon] %s\n", msg.Line)
		})
	case *messages7.RconCmdAdd:
		userMsgCallback(client.Callbacks.SysRconCmdAdd, msg, func() {
			fmt.Printf("got rcon cmd=%s %s %s\n", msg.Name, msg.Params, msg.Help)
		})
	case *messages7.RconCmdRem:
		userMsgCallback(client.Callbacks.SysRconCmdRem, msg, func() {
			fmt.Printf("removed cmd=%s\n", msg.Name)
		})
	case *messages7.Unknown:
		userMsgCallback(client.Callbacks.MsgUnknown, msg, func() {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown system")
		})
	default:
		printUnknownMessage(netMsg, "unprocessed system")
		return false
	}
	return true
}

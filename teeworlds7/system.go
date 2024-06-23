package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

func (client *Client) processSystem(netMsg messages7.NetMessage, response *protocol7.Packet) bool {
	switch msg := netMsg.(type) {
	case *messages7.MapChange:
		defaultAction := func() {
			fmt.Println("got map change")
			response.Messages = append(response.Messages, &messages7.Ready{})
		}
		if client.Callbacks.SysMapChange == nil {
			defaultAction()
		} else {
			client.Callbacks.SysMapChange(msg, defaultAction)
		}
	case *messages7.MapData:
		defaultAction := func() {
			fmt.Printf("got map chunk %x\n", msg.Data)
		}
		if client.Callbacks.SysMapData == nil {
			defaultAction()
		} else {
			client.Callbacks.SysMapData(msg, defaultAction)
		}
	case *messages7.ServerInfo:
		defaultAction := func() {
			fmt.Printf("connected to server with name '%s'\n", msg.Name)
		}
		if client.Callbacks.SysServerInfo == nil {
			defaultAction()
		} else {
			client.Callbacks.SysServerInfo(msg, defaultAction)
		}
	case *messages7.ConReady:
		defaultAction := func() {
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
		}
		if client.Callbacks.SysConReady == nil {
			defaultAction()
		} else {
			client.Callbacks.SysConReady(msg, defaultAction)
		}
	case *messages7.Snap:
		// fmt.Printf("got snap tick=%d\n", msg.GameTick)
		response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
	case *messages7.SnapSingle:
		// fmt.Printf("got snap single tick=%d\n", msg.GameTick)
		response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
	case *messages7.SnapEmpty:
		// fmt.Printf("got snap empty tick=%d\n", msg.GameTick)
		response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
	case *messages7.InputTiming:
		defaultAction := func() {
			fmt.Printf("timing time left=%d\n", msg.TimeLeft)
		}
		if client.Callbacks.SysInputTiming == nil {
			defaultAction()
		} else {
			client.Callbacks.SysInputTiming(msg, defaultAction)
		}
	case *messages7.RconAuthOn:
		defaultAction := func() {
			fmt.Println("you are now authenticated in rcon")
		}
		if client.Callbacks.SysRconAuthOn == nil {
			defaultAction()
		} else {
			client.Callbacks.SysRconAuthOn(msg, defaultAction)
		}
	case *messages7.RconAuthOff:
		defaultAction := func() {
			fmt.Println("you are no longer authenticated in rcon")
		}
		if client.Callbacks.SysRconAuthOff == nil {
			defaultAction()
		} else {
			client.Callbacks.SysRconAuthOff(msg, defaultAction)
		}
	case *messages7.RconLine:
		defaultAction := func() {
			fmt.Printf("[rcon] %s\n", msg.Line)
		}
		if client.Callbacks.SysRconLine == nil {
			defaultAction()
		} else {
			client.Callbacks.SysRconLine(msg, defaultAction)
		}
	case *messages7.RconCmdAdd:
		defaultAction := func() {
			fmt.Printf("got rcon cmd=%s %s %s\n", msg.Name, msg.Params, msg.Help)
		}
		if client.Callbacks.SysRconCmdAdd == nil {
			defaultAction()
		} else {
			client.Callbacks.SysRconCmdAdd(msg, defaultAction)
		}
	case *messages7.RconCmdRem:
		defaultAction := func() {
			fmt.Printf("removed cmd=%s\n", msg.Name)
		}
		if client.Callbacks.SysRconCmdRem == nil {
			defaultAction()
		} else {
			client.Callbacks.SysRconCmdRem(msg, defaultAction)
		}
	case *messages7.Unknown:
		defaultAction := func() {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown system")
		}
		if client.Callbacks.MsgUnknown == nil {
			defaultAction()
		} else {
			client.Callbacks.MsgUnknown(msg, defaultAction)
		}
	default:
		printUnknownMessage(netMsg, "unprocessed system")
		return false
	}
	return true
}

package protocol7

import (
	"bytes"
	"fmt"
	"os"

	"github.com/teeworlds-go/teeworlds/messages7"
	"github.com/teeworlds-go/teeworlds/network7"
)

type Player struct {
	Info messages7.SvClientInfo
}

type Connection struct {
	ClientToken [4]byte
	ServerToken [4]byte

	// The amount of vital chunks received
	Ack int

	// The amount of vital chunks sent
	Sequence int

	// The amount of vital chunks acknowledged by the peer
	PeerAck int

	Players []Player
}

func (connection *Connection) BuildResponse() *Packet {
	return &Packet{
		Header: PacketHeader{
			Flags: PacketFlags{
				Connless:    false,
				Compression: false,
				Resend:      false,
				Control:     false,
			},
			Ack:       connection.Ack,
			NumChunks: 0, // will be set in Packet.Pack()
			Token:     connection.ServerToken,
		},
	}
}

func (connection *Connection) CtrlToken() *Packet {
	response := connection.BuildResponse()
	response.Header.Flags.Control = true
	response.Messages = append(
		response.Messages,
		&messages7.CtrlToken{
			Token: connection.ClientToken,
		},
	)

	return response
}

func (client *Connection) MsgStartInfo() *messages7.ClStartInfo {
	return &messages7.ClStartInfo{
		Name:                  "gopher",
		Clan:                  "",
		Country:               0,
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
}

func byteSliceToString(s []byte) string {
	n := bytes.IndexByte(s, 0)
	if n >= 0 {
		s = s[:n]
	}
	return string(s)
}

func (connection *Connection) printUnknownMessage(msg messages7.NetMessage, msgType string) {
	fmt.Printf("%s message id=%d\n", msgType, msg.MsgId())
	if msg.Header() == nil {
		fmt.Println("  header: nil")
	} else {
		fmt.Printf("  header: %x\n", msg.Header().Pack())
	}
	fmt.Printf("  payload: %x\n", msg.Pack())
	if msg.Header() != nil {
		fmt.Printf("  full msg: %x%x\n", msg.Header().Pack(), msg.Pack())
	}
}

func (connection *Connection) OnSystemMsg(msg messages7.NetMessage, response *Packet) bool {
	// TODO: is this shadow nasty?
	switch msg := msg.(type) {
	case *messages7.MapChange:
		fmt.Println("got map change")
		response.Messages = append(response.Messages, &messages7.Ready{})
	case *messages7.ConReady:
		fmt.Println("got ready")
		response.Messages = append(response.Messages, connection.MsgStartInfo())
	case *messages7.SnapSingle:
		// fmt.Printf("got snap single tick=%d\n", msg.GameTick)
		response.Messages = append(response.Messages, &messages7.CtrlKeepAlive{})
	case *messages7.SnapEmpty:
		// fmt.Printf("got snap empty tick=%d\n", msg.GameTick)
	case *messages7.InputTiming:
		// fmt.Printf("timing time left=%d\n", msg.TimeLeft)
	case *messages7.Unknown:
		// TODO: msg id of unknown messages should not be -1
		fmt.Println("TODO: why is the msg id -1???")
		connection.printUnknownMessage(msg, "unknown system")
	default:
		connection.printUnknownMessage(msg, "unhandled system")
		return false
	}
	return true
}

func (client *Connection) OnChatMessage(msg *messages7.SvChat) {
	if msg.ClientId < 0 || msg.ClientId > network7.MaxClients {
		fmt.Printf("[chat] *** %s\n", msg.Message)
		return
	}
	name := client.Players[msg.ClientId].Info.Name
	fmt.Printf("[chat] <%s> %s\n", name, msg.Message)
}

func (connection *Connection) OnGameMsg(msg messages7.NetMessage, response *Packet) bool {
	// TODO: is this shadow nasty?
	switch msg := msg.(type) {
	case *messages7.ReadyToEnter:
		fmt.Println("got ready to enter")
		response.Messages = append(response.Messages, &messages7.EnterGame{})
	case *messages7.SvMotd:
		if msg.Message != "" {
			fmt.Printf("[motd] %s\n", msg.Message)
		}
	case *messages7.SvChat:
		connection.OnChatMessage(msg)
	case *messages7.SvClientInfo:
		connection.Players[msg.ClientId].Info = *msg
		fmt.Printf("got client info id=%d name=%s\n", msg.ClientId, msg.Name)
	case *messages7.Unknown:
		connection.printUnknownMessage(msg, "unknown game")
	default:
		connection.printUnknownMessage(msg, "unhandled game")
		return false
	}
	return true
}

func (connection *Connection) OnMessage(msg messages7.NetMessage, response *Packet) bool {
	if msg.Header() == nil {
		// this is probably an unknown message
		fmt.Printf("warning ignoring msgId=%d because header is nil\n", msg.MsgId())
		return false
	}
	if msg.Header().Flags.Vital {
		connection.Ack++
	}

	if msg.System() {
		return connection.OnSystemMsg(msg, response)
	}
	return connection.OnGameMsg(msg, response)
}

// Takes a full teeworlds packet as argument
// And returns the response packet from the clients perspective
func (connection *Connection) OnPacket(packet *Packet) *Packet {
	response := connection.BuildResponse()

	if packet.Header.Flags.Control {
		msg := packet.Messages[0]
		fmt.Printf("got ctrl msg %d\n", msg.MsgId())
		// TODO: is this shadow nasty?
		switch msg := msg.(type) {
		case *messages7.CtrlToken:
			fmt.Printf("got server token %x\n", msg.Token)
			connection.ServerToken = msg.Token
			response.Header.Token = msg.Token
			response.Messages = append(
				response.Messages,
				&messages7.CtrlConnect{
					Token: connection.ClientToken,
				},
			)
		case *messages7.CtrlAccept:
			fmt.Println("got accept")
			// TODO: don't hardcode info
			response.Messages = append(response.Messages, &messages7.Info{})
		case *messages7.CtrlClose:
			fmt.Printf("disconnected (%s)\n", msg.Reason)
			os.Exit(0)
		default:
			fmt.Printf("unknown control message: %d\n", msg.MsgId())
		}
		return response
	}

	for _, msg := range packet.Messages {
		connection.OnMessage(msg, response)
	}

	return response
}

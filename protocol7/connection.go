package protocol7

import (
	"bytes"
	"fmt"
	"os"

	"github.com/teeworlds-go/huffman"
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/messages7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
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

func (connection *Connection) OnSystemMsg(msg int, chunk chunk7.Chunk, u *packer.Unpacker, result *PacketResult) {
	if msg == network7.MsgSysMapChange {
		fmt.Println("got map change")
		result.Response.Messages = append(result.Response.Messages, &messages7.Ready{})
	} else if msg == network7.MsgSysConReady {
		fmt.Println("got ready")
		result.Response.Messages = append(result.Response.Messages, connection.MsgStartInfo())
	} else if msg == network7.MsgSysSnapSingle {
		// tick := u.GetInt()
		// fmt.Printf("got snap single tick=%d\n", tick)
		result.Response.Messages = append(result.Response.Messages, &messages7.CtrlKeepAlive{})
	} else {
		fmt.Printf("unknown system message id=%d data=%x\n", msg, chunk.Data)
	}
}

func (client *Connection) OnChatMessage(msg *messages7.SvChat) {
	if msg.ClientId < 0 || msg.ClientId > network7.MaxClients {
		fmt.Printf("[chat] *** %s\n", msg.Message)
		return
	}
	name := client.Players[msg.ClientId].Info.Name
	fmt.Printf("[chat] <%s> %s\n", name, msg.Message)
}

func (client *Connection) OnMotd(motd string) {
	fmt.Printf("[motd] %s\n", motd)
}

func (client *Connection) OnGameMsg(msg int, chunk chunk7.Chunk, u *packer.Unpacker, result *PacketResult) {
	if msg == network7.MsgGameReadyToEnter {
		fmt.Println("got ready to enter")
		result.Packet.Messages = append(result.Packet.Messages, &messages7.Ready{ChunkHeader: &chunk.Header})
		result.Response.Messages = append(result.Response.Messages, &messages7.EnterGame{})
	} else if msg == network7.MsgGameSvMotd {
		motd := u.GetString()
		if motd != "" {
			client.OnMotd(motd)
		}
	} else if msg == network7.MsgGameSvChat {
		chat := &messages7.SvChat{ChunkHeader: &chunk.Header}
		chat.Unpack(u)
		client.OnChatMessage(chat)
		result.Packet.Messages = append(result.Packet.Messages, chat)
	} else if msg == network7.MsgGameSvClientInfo {
		clientId := packer.UnpackInt(chunk.Data[1:])
		client.Players[clientId].Info.Unpack(u)

		fmt.Printf("got client info id=%d name=%s\n", clientId, client.Players[clientId].Info.Name)
	} else {
		fmt.Printf("unknown game message id=%d data=%x\n", msg, chunk.Data)
	}
}

func (client *Connection) OnMessage(chunk chunk7.Chunk, result *PacketResult) {
	// fmt.Printf("got chunk size=%d data=%v\n", chunk.Header.Size, chunk.Data)

	if chunk.Header.Flags.Vital {
		client.Ack++
	}

	u := packer.Unpacker{}
	u.Reset(chunk.Data)

	msg := u.GetInt()

	sys := msg&1 != 0
	msg >>= 1

	if sys {
		client.OnSystemMsg(msg, chunk, &u, result)
	} else {
		client.OnGameMsg(msg, chunk, &u, result)
	}
}

func (connection *Connection) OnPacketPayload(data []byte, result *PacketResult) (*PacketResult, error) {
	chunks := chunk7.UnpackChunks(data)

	for _, c := range chunks {
		connection.OnMessage(c, result)
	}
	return result, nil

}

type PacketResult struct {
	// Suggested response that should be sent to the server
	// Will be *nil* if no response should be sent
	Response *Packet

	// Incoming traffic from the server parsed into a Packet struct
	Packet *Packet
}

func (connection *Connection) OnPacket(data []byte) (*PacketResult, error) {
	result := &PacketResult{
		Response: connection.BuildResponse(),
		Packet:   &Packet{},
	}
	result.Packet.Header.Unpack(data[:7])
	payload := data[7:]

	if result.Packet.Header.Flags.Control {
		ctrlMsg := int(payload[0])
		fmt.Printf("got ctrl msg %d\n", ctrlMsg)
		if ctrlMsg == network7.MsgCtrlToken {
			copy(connection.ServerToken[:], payload[1:5])
			result.Response.Header.Token = connection.ServerToken
			fmt.Printf("got server token %x\n", connection.ServerToken)
			result.Packet.Messages = append(result.Packet.Messages, &messages7.CtrlToken{Token: connection.ServerToken})
			result.Response.Messages = append(
				result.Response.Messages,
				&messages7.CtrlConnect{
					Token: connection.ClientToken,
				},
			)
		} else if ctrlMsg == network7.MsgCtrlAccept {
			fmt.Println("got accept")
			result.Packet.Messages = append(result.Packet.Messages, &messages7.CtrlAccept{})
			// TODO: don't hardcode info
			result.Response.Messages = append(result.Response.Messages, &messages7.Info{})
		} else if ctrlMsg == network7.MsgCtrlClose {
			// TODO: get length from packet header to determine if a reason is set or not
			// len(data) -> is 1400 (maxPacketLen)

			reason := byteSliceToString(payload)
			fmt.Printf("disconnected (%s)\n", reason)

			os.Exit(0)
		} else {
			fmt.Printf("unknown control message: %x\n", data)
		}

		if len(result.Response.Messages) == 0 {
			return nil, nil
		}

		return result, nil
	}

	if result.Packet.Header.Flags.Compression {
		huff := huffman.Huffman{}
		var err error
		payload, err = huff.Decompress(payload)
		if err != nil {
			fmt.Printf("huffman error: %v\n", err)
			return nil, nil
		}
	}

	result, err := connection.OnPacketPayload(payload, result)
	return result, err
}

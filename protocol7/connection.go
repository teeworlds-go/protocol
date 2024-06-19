package protocol7

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"slices"

	"github.com/teeworlds-go/huffman"
	"github.com/teeworlds-go/teeworlds/chunk"
	message "github.com/teeworlds-go/teeworlds/messages"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
	"github.com/teeworlds-go/teeworlds/packet"
)

type Player struct {
	Info message.SvClientInfo
}

type Connection struct {
	ClientToken [4]byte
	ServerToken [4]byte
	Conn        net.Conn

	// The amount of vital chunks received
	Ack int

	// The amount of vital chunks sent
	Sequence int

	// The amount of vital chunks acknowledged by the peer
	PeerAck int

	Players []Player
}

func (c *Connection) SendCtrlToken(myToken []byte) {
	header := []byte{0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff}
	ctrlToken := append([]byte{0x05}, myToken...)
	zeros := []byte{512: 0}
	data := slices.Concat(header, ctrlToken, zeros)
	c.Conn.Write(data)
}

func (c *Connection) SendCtrlMsg(data []byte) {
	header := packet.PacketHeader{
		Flags: packet.PacketFlags{
			Connless:    false,
			Compression: false,
			Resend:      false,
			Control:     true,
		},
		Ack:       c.Ack,
		NumChunks: 0,
		Token:     c.ServerToken,
	}

	packet := slices.Concat(header.Pack(), data)
	c.Conn.Write(packet)
}

func (c *Connection) SendKeepAlive() {
	c.SendCtrlMsg([]byte{byte(network7.MsgCtrlKeepAlive)})
}

func (c *Connection) SendReady() {
	ready := []byte{0x40, 0x01, 0x02, 0x25}

	c.Sequence++
	c.SendPacket(ready, 1)
}

type ChunkArgs struct {
	MsgId   network7.NetMsg
	System  bool
	Flags   chunk.ChunkFlags
	Payload []byte
}

func (client *Connection) PackChunk(c ChunkArgs) []byte {
	c.MsgId <<= 1
	if c.System {
		c.MsgId |= 1
	}

	client.Sequence++
	msgAndSys := packer.PackMsg(c.MsgId)

	chunkHeader := chunk.ChunkHeader{
		Flags: c.Flags,
		Size:  len(msgAndSys) + len(c.Payload),
		Seq:   client.Sequence,
	}

	data := slices.Concat(
		chunkHeader.Pack(),
		msgAndSys,
		c.Payload,
	)

	fmt.Printf("packed chunk: %x\n", data)

	return data
}

func (client *Connection) SendPacket(payload []byte, numChunks int) {
	header := packet.PacketHeader{
		Flags: packet.PacketFlags{
			Connless:    false,
			Compression: false,
			Resend:      false,
			Control:     false,
		},
		Ack:       client.Ack,
		NumChunks: numChunks,
		Token:     client.ServerToken,
	}

	packet := slices.Concat(header.Pack(), payload)
	client.Conn.Write(packet)
}

func (client *Connection) SendInfo() {
	info := []byte{0x40, 0x28, 0x01, 0x03, 0x30, 0x2E, 0x37, 0x20, 0x38, 0x30, 0x32, 0x66,
		0x31, 0x62, 0x65, 0x36, 0x30, 0x61, 0x30, 0x35, 0x36, 0x36, 0x35, 0x66,
		0x00, 0x6D, 0x79, 0x5F, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6F, 0x72, 0x64,
		0x5F, 0x31, 0x32, 0x33, 0x00, 0x85, 0x1C, 0x00}

	client.Sequence++
	client.SendPacket(info, 1)
}

func (client *Connection) SendStartInfo() {
	info := message.ClStartInfo{
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

	payload := client.PackChunk(ChunkArgs{
		MsgId: network7.MsgGameClStartInfo,
		Flags: chunk.ChunkFlags{
			Vital: true,
		},
		Payload: info.Pack(),
	})

	client.SendPacket(payload, 1)
}

func (client *Connection) SendEnterGame() {
	enter := []byte{
		0x40, 0x01, 0x04, 0x27,
	}

	client.SendPacket(enter, 1)
}

func byteSliceToString(s []byte) string {
	n := bytes.IndexByte(s, 0)
	if n >= 0 {
		s = s[:n]
	}
	return string(s)
}

func (client *Connection) OnSystemMsg(msg network7.NetMsg, chunk chunk.Chunk, u *packer.Unpacker) {
	if msg == network7.MsgSysMapChange {
		fmt.Println("got map change")
		client.SendReady()
	} else if msg == network7.MsgSysConReady {
		fmt.Println("got ready")
		client.SendStartInfo()
	} else if msg == network7.MsgSysSnapSingle {
		// tick := u.GetInt()
		// fmt.Printf("got snap single tick=%d\n", tick)
		client.SendKeepAlive()
	} else {
		fmt.Printf("unknown system message id=%d data=%x\n", msg, chunk.Data)
	}
}

func (client *Connection) OnChatMessage(mode int, clientId int, targetId int, message string) {
	name := client.Players[clientId].Info.Name
	fmt.Printf("[chat] <%s> %s\n", name, message)
}

func (client *Connection) OnMotd(motd string) {
	fmt.Printf("[motd] %s\n", motd)
}

func (client *Connection) OnGameMsg(msg network7.NetMsg, chunk chunk.Chunk, u *packer.Unpacker) {
	if msg == network7.MsgGameReadyToEnter {
		fmt.Println("got ready to enter")
		client.SendEnterGame()
	} else if msg == network7.MsgGameSvMotd {
		motd := u.GetString()
		if motd != "" {
			client.OnMotd(motd)
		}
	} else if msg == network7.MsgGameSvChat {
		mode := u.GetInt()
		clientId := u.GetInt()
		targetId := u.GetInt()
		message := u.GetString()
		client.OnChatMessage(mode, clientId, targetId, message)
	} else if msg == network7.MsgGameSvClientInfo {
		clientId := packer.UnpackInt(chunk.Data[1:])
		client.Players[clientId].Info.Unpack(u)

		fmt.Printf("got client info id=%d name=%s\n", clientId, client.Players[clientId].Info.Name)
	} else {
		fmt.Printf("unknown game message id=%d data=%x\n", msg, chunk.Data)
	}
}

func (client *Connection) OnMessage(chunk chunk.Chunk) {
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
		client.OnSystemMsg(network7.NetMsg(msg), chunk, &u)
	} else {
		client.OnGameMsg(network7.NetMsg(msg), chunk, &u)
	}
}

func (client *Connection) OnPacketPayload(header []byte, data []byte) {
	chunks := chunk.UnpackChunks(data)

	for _, c := range chunks {
		client.OnMessage(c)
	}
}

func (client *Connection) OnPacket(data []byte) {
	header := packet.PacketHeader{}
	headerRaw := data[:7]
	payload := data[7:]
	header.Unpack(headerRaw)

	if header.Flags.Control {
		ctrlMsg := network7.ControlMsg(payload[0])
		fmt.Printf("got ctrl msg %d\n", ctrlMsg)
		if ctrlMsg == network7.MsgCtrlToken {
			copy(client.ServerToken[:], payload[1:5])
			fmt.Printf("got server token %x\n", client.ServerToken)
			client.SendCtrlMsg(slices.Concat([]byte{byte(network7.MsgCtrlConnect)}, client.ClientToken[:]))
		} else if ctrlMsg == network7.MsgCtrlAccept {
			fmt.Println("got accept")
			client.SendInfo()
		} else if ctrlMsg == network7.MsgCtrlClose {
			// TODO: get length from packet header to determine if a reason is set or not
			// len(data) -> is 1400 (maxPacketLen)

			reason := byteSliceToString(payload)
			fmt.Printf("disconnected (%s)\n", reason)

			os.Exit(0)
		} else {
			fmt.Printf("unknown control message: %x\n", data)
		}
		return
	}

	if header.Flags.Compression {
		huff := huffman.Huffman{}
		var err error
		payload, err = huff.Decompress(payload)
		if err != nil {
			fmt.Printf("huffman error: %v\n", err)
			return
		}
	}

	client.OnPacketPayload(headerRaw, payload)
}

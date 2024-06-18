package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"slices"
	"time"

	"github.com/teeworlds-go/huffman"
	"github.com/teeworlds-go/teeworlds/chunk"
	"github.com/teeworlds-go/teeworlds/packer"
	"github.com/teeworlds-go/teeworlds/packet"
)

const (
	maxPacksize = 1400

	msgCtrlKeepAlive = 0x00
	msgCtrlConnect   = 0x01
	msgCtrlAccept    = 0x02
	msgCtrlToken     = 0x05
	msgCtrlClose     = 0x04

	msgSysMapChange  = 2
	msgSysConReady   = 5
	msgSysSnapSingle = 8

	msgGameMotd         = 1
	msgGameSvChat       = 3
	msgGameReadyToEnter = 8
)

func ctrlToken(myToken []byte) []byte {
	header := []byte{0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff}
	ctrlToken := append([]byte{0x05}, myToken...)
	zeros := []byte{512: 0}
	data := slices.Concat(header, ctrlToken, zeros)
	return data
}

func getConnection() (net.Conn, error) {
	conn, err := net.Dial("udp", "127.0.0.1:8303")
	if err != nil {
		fmt.Printf("Some error %v", err)
	}
	return conn, err
}

func readNetwork(ch chan<- []byte, conn net.Conn) {
	buf := make([]byte, maxPacksize)

	for {
		len, err := bufio.NewReader(conn).Read(buf)
		packet := make([]byte, len)
		copy(packet[:], buf[:])
		if err == nil {
			ch <- packet
		} else {
			fmt.Printf("Some error %v\n", err)
			break
		}
	}

	conn.Close()
}

type TeeworldsClient struct {
	clientToken [4]byte
	serverToken [4]byte
	conn        net.Conn

	Ack int
}

func (client TeeworldsClient) sendCtrlMsg(data []byte) {
	header := packet.PacketHeader{
		Flags: packet.PacketFlags{
			Connless:    false,
			Compression: false,
			Resend:      false,
			Control:     true,
		},
		Ack:       client.Ack,
		NumChunks: 0,
		Token:     client.serverToken,
	}

	packet := slices.Concat(header.Pack(), data)
	client.conn.Write(packet)
}

func (client TeeworldsClient) sendReady() {
	packet := slices.Concat(
		[]byte{0x00, 0x01, 0x01},
		client.serverToken[:],
		[]byte{0x40, 0x01, 0x02, 0x25},
	)
	client.conn.Write(packet)
}

func (client TeeworldsClient) sendKeepAlive() {
	client.sendCtrlMsg([]byte{msgCtrlKeepAlive})
}

func (client TeeworldsClient) sendInfo() {
	info := []byte{0x40, 0x28, 0x01, 0x03, 0x30, 0x2E, 0x37, 0x20, 0x38, 0x30, 0x32, 0x66,
		0x31, 0x62, 0x65, 0x36, 0x30, 0x61, 0x30, 0x35, 0x36, 0x36, 0x35, 0x66,
		0x00, 0x6D, 0x79, 0x5F, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6F, 0x72, 0x64,
		0x5F, 0x31, 0x32, 0x33, 0x00, 0x85, 0x1C, 0x00}

	packet := slices.Concat(
		[]byte{0x00, 0x00, 0x01},
		client.serverToken[:],
		info,
	)

	client.conn.Write(packet)
}

func (client TeeworldsClient) sendStartInfo() {
	info := []byte{
		0x41, 0x14, 0x03, 0x36, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x00, 0x00, 0x40, 0x67, 0x72, 0x65,
		0x65, 0x6e, 0x73, 0x77, 0x61, 0x72, 0x64, 0x00, 0x64, 0x75, 0x6f, 0x64, 0x6f, 0x6e, 0x6e, 0x79,
		0x00, 0x00, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73, 0x74, 0x61, 0x6e, 0x64,
		0x61, 0x72, 0x64, 0x00, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x01, 0x01, 0x00,
		0x00, 0x00, 0x00, 0x80, 0xfc, 0xaf, 0x05, 0xeb, 0x83, 0xd0, 0x0a, 0x80, 0xfe, 0x07, 0x80, 0xfe,
		0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07,
	}

	packet := slices.Concat(
		[]byte{0x00, 0x04, 0x01},
		client.serverToken[:],
		info,
	)

	client.conn.Write(packet)
}

func (client TeeworldsClient) sendEnterGame() {
	enter := []byte{
		0x40, 0x01, 0x04, 0x27,
	}

	packet := slices.Concat(
		[]byte{0x00, 0x07, 0x01},
		client.serverToken[:],
		enter,
	)

	client.conn.Write(packet)
}

func isCtrlMsg(data []byte) bool {
	return data[0] == 0x04
}

func isMapChange(data []byte) bool {
	// unsafe and trol
	return data[10] == msgSysMapChange
}

func isCompressed(data []byte) bool {
	// we don't talk about it
	return data[0] == 0x10
}

func numberOfChunks(data []byte) int {
	return int(data[2])
}

func isReadyToEnter(data []byte) bool {
	return !isCompressed(data) && numberOfChunks(data) == 3
}

func isConReady(data []byte) bool {
	// unsafe and yemDX level troling
	return isCompressed(data) // data[10] == msgSysMapChange
}

func byteSliceToString(s []byte) string {
	n := bytes.IndexByte(s, 0)
	if n >= 0 {
		s = s[:n]
	}
	return string(s)
}

func (client *TeeworldsClient) onSystemMsg(msg int, chunk chunk.Chunk, u *packer.Unpacker) {
	if msg == msgSysMapChange {
		fmt.Println("got map change")
		client.sendReady()
	} else if msg == msgSysConReady {
		fmt.Println("got ready")
		client.sendStartInfo()
	} else if msg == msgSysSnapSingle {
		// tick := u.GetInt()
		// fmt.Printf("got snap single tick=%d\n", tick)
		client.sendKeepAlive()
	} else {
		fmt.Printf("unknown system message id=%d data=%x\n", msg, chunk.Data)
	}
}

func onChatMessage(mode int, clientId int, targetId int, message string) {
	fmt.Printf("[chat] %d: %s\n", clientId, message)
}

func (client *TeeworldsClient) onGameMsg(msg int, chunk chunk.Chunk, u *packer.Unpacker) {
	if msg == msgGameReadyToEnter {
		fmt.Println("got ready to enter")
		client.sendEnterGame()
	} else if msg == msgGameSvChat {
		mode := u.GetInt()
		clientId := u.GetInt()
		targetId := u.GetInt()
		message := u.GetString()
		onChatMessage(mode, clientId, targetId, message)
	} else {
		fmt.Printf("unknown game message id=%d data=%x\n", msg, chunk.Data)
	}
}

func (client *TeeworldsClient) onMessage(chunk chunk.Chunk) {
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
		client.onSystemMsg(msg, chunk, &u)
	} else {
		client.onGameMsg(msg, chunk, &u)
	}
}

func (client *TeeworldsClient) onPacketPayload(header []byte, data []byte) {
	chunks := chunk.UnpackChunks(data)

	for _, c := range chunks {
		client.onMessage(c)
	}
}

func (client *TeeworldsClient) onPacket(data []byte) {
	header := packet.PacketHeader{}
	headerRaw := data[:7]
	payload := data[7:]
	header.Unpack(headerRaw)

	if header.Flags.Control {
		ctrlMsg := payload[0]
		fmt.Printf("got ctrl msg %d\n", ctrlMsg)
		if ctrlMsg == msgCtrlToken {
			copy(client.serverToken[:], payload[1:5])
			fmt.Printf("got server token %x\n", client.serverToken)
			client.sendCtrlMsg(slices.Concat([]byte{msgCtrlConnect}, client.clientToken[:]))
		} else if ctrlMsg == msgCtrlAccept {
			fmt.Println("got accept")
			client.sendInfo()
		} else if ctrlMsg == msgCtrlClose {
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

	client.onPacketPayload(headerRaw, payload)
}

func main() {
	ch := make(chan []byte, maxPacksize)

	conn, err := getConnection()
	if err != nil {
		fmt.Printf("error connecting %v\n", err)
		os.Exit(1)
	}

	client := &TeeworldsClient{
		clientToken: [4]byte{0x01, 0x02, 0x03, 0x04},
		serverToken: [4]byte{0xff, 0xff, 0xff, 0xff},
		conn:        conn,
	}

	go readNetwork(ch, client.conn)

	conn.Write(ctrlToken(client.clientToken[:]))

	for {
		time.Sleep(10_000_000)
		select {
		case msg := <-ch:
			client.onPacket(msg)
		default:
			// do nothing
		}
	}

}

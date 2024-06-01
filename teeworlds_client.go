package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"slices"
	"time"
)

const (
	maxPacksize = 1400

	msgCtrlConnect = 0x01
	msgCtrlAccept  = 0x02
	msgCtrlToken   = 0x05
	msgCtrlClose   = 0x04

	msgSysMapChange = 0x05
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
	packet := make([]byte, maxPacksize)

	for {
		_, err := bufio.NewReader(conn).Read(packet)
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
}

func (client TeeworldsClient) sendCtrlMsg(data []byte) {
	flags := []byte{0x04, 0x00, 0x00}
	packet := slices.Concat(flags, client.serverToken[:], data)

	// fmt.Printf("sending %v\n", packet)

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

func isCtrlMsg(data []byte) bool {
	return data[0] == 0x04
}

func isMapChange(data []byte) bool {
	// unsafe and trol
	return data[10] == msgSysMapChange
}

func byteSliceToString(s []byte) string {
	n := bytes.IndexByte(s, 0)
	if n >= 0 {
		s = s[:n]
	}
	return string(s)
}

func (client *TeeworldsClient) onMessage(data []byte) {
	if isCtrlMsg(data) {
		ctrlMsg := data[7]
		fmt.Printf("got ctrl msg %d\n", ctrlMsg)
		if ctrlMsg == msgCtrlToken {
			copy(client.serverToken[:], data[8:12])
			fmt.Printf("got server token %v\n", client.serverToken)
			client.sendCtrlMsg(slices.Concat([]byte{msgCtrlConnect}, client.clientToken[:]))
		} else if ctrlMsg == msgCtrlAccept {
			fmt.Println("got accept")
			client.sendInfo()
		} else if ctrlMsg == msgCtrlClose {
			// TODO: get length from packet header to determine if a reason is set or not
			// len(data) -> is 1400 (maxPacketLen)

			reason := byteSliceToString(data[8:])
			fmt.Printf("disconnected (%s)\n", reason)

			os.Exit(0)
		} else {
			fmt.Printf("unknown control message: %v\n", data)
		}
	} else if isMapChange(data) {
		fmt.Println("got map change")
		client.sendReady()

	} else {
		fmt.Printf("unknown message: %v\n", data)
	}
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
			client.onMessage(msg)
		default:
			// do nothing
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"net"
	"slices"
	"time"
)

const (
	maxPacksize  = 1400
	msgCtrlToken = 0x04
)

func ctrlToken(myToken []byte) []byte {
	header := []byte{0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff}
	ctrlToken := append([]byte{0x05}, myToken...)
	zeros := []byte{512: 0}
	data := slices.Concat(header, ctrlToken, zeros)
	return data
}

func pumpNetwork(ch chan []byte) {
	packet := make([]byte, maxPacksize)
	conn, err := net.Dial("udp", "127.0.0.1:8303")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	// myToken := []byte{0xfe, 0xed, 0xba, 0xbe}
	myToken := []byte{0x01, 0x02, 0x03, 0x04}
	conn.Write(ctrlToken(myToken))

	for {
		_, err = bufio.NewReader(conn).Read(packet)
		if err == nil {
			ch <- packet
		} else {
			fmt.Printf("Some error %v\n", err)
			break
		}
	}

	conn.Close()
}

func onMessage(data []byte) {
	if data[0] == msgCtrlToken {
		serverToken := data[8:12]
		fmt.Printf("got token %v\n", serverToken)
	} else {
		fmt.Println("unknown message")
	}
}

func main() {
	ch := make(chan []byte, maxPacksize)

	go pumpNetwork(ch)

	for {
		time.Sleep(10_000_000)
		select {
		case msg := <-ch:
			onMessage(msg)
		default:
			// do nothing
		}
	}

}

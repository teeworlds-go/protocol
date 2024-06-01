package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

func getConnection() (net.Conn, error) {
	conn, err := net.Dial("udp", "127.0.0.1:8303")
	if err != nil {
		fmt.Printf("Some error %v", err)
	}
	return conn, err
}

func readNetwork(ch chan []byte, conn net.Conn) {
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

func onMessage(data []byte, conn net.Conn) {
	if data[0] == msgCtrlToken {
		serverToken := data[8:12]
		fmt.Printf("got token %v\n", serverToken)
		conn.Write([]byte{0xff, 0xff, 0xff})
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

	go readNetwork(ch, conn)

	myToken := []byte{0x01, 0x02, 0x03, 0x04}
	conn.Write(ctrlToken(myToken))

	for {
		time.Sleep(10_000_000)
		select {
		case msg := <-ch:
			onMessage(msg, conn)
		default:
			// do nothing
		}
	}

}

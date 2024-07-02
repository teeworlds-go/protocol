package teeworlds7

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
)

const (
	maxPacksize = 1400
)

func getConnection(serverIp string, serverPort int) (net.Conn, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", serverIp, serverPort))
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

func (client *Client) Connect(serverIp string, serverPort int) {
	ch := make(chan []byte, maxPacksize)

	conn, err := getConnection(serverIp, serverPort)
	if err != nil {
		fmt.Printf("error connecting %v\n", err)
		return
	}

	client.Session = protocol7.Session{
		ClientToken: [4]byte{0x01, 0x02, 0x03, 0x04},
		ServerToken: [4]byte{0xff, 0xff, 0xff, 0xff},
		Ack:         0,
	}
	client.Game.Players = make([]Player, network7.MaxClients)
	client.Conn = conn

	go readNetwork(ch, conn)

	client.SendPacket(client.Session.CtrlToken())

	// TODO: do we really need a non blocking network read?
	//       if not remove the channel, the sleep and the select statement
	//       if yes also offer an OnTick callback to the user, and also do keepalives and resends
	for {
		time.Sleep(10_000_000)
		select {
		case msg := <-ch:
			packet := &protocol7.Packet{}
			err := packet.Unpack(msg)
			if err != nil {
				client.throwError(err)
			}
			err = client.processPacket(packet)
			if err != nil {
				client.throwError(err)
			}
		default:
			// do nothing
		}
	}
}

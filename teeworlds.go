package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/teeworlds-go/teeworlds/messages7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/protocol7"
)

const (
	maxPacksize = 1400
)

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

func main() {
	ch := make(chan []byte, maxPacksize)

	conn, err := getConnection()
	if err != nil {
		fmt.Printf("error connecting %v\n", err)
		os.Exit(1)
	}

	client := &protocol7.Connection{
		ClientToken: [4]byte{0x01, 0x02, 0x03, 0x04},
		ServerToken: [4]byte{0xff, 0xff, 0xff, 0xff},
		Ack:         0,
		Players:     make([]protocol7.Player, network7.MaxClients),
	}

	go readNetwork(ch, conn)

	tokenPacket := client.CtrlToken()
	conn.Write(tokenPacket.Pack(client))

	for {
		time.Sleep(10_000_000)
		select {
		case msg := <-ch:
			result, err := client.OnPacket(msg)
			if err != nil {
				panic(err)
			}
			if result.Response != nil {

				// example of inspecting incoming trafic
				for i, msg := range result.Packet.Messages {
					if msg.MsgId() == network7.MsgGameSvChat {
						var chat *messages7.SvChat
						var ok bool
						if chat, ok = result.Packet.Messages[i].(*messages7.SvChat); ok {
							fmt.Printf("got chat msg: %s\n", chat.Message)

							// modify chat if this was a proxy
							result.Packet.Messages[i] = chat
						}
					}
				}

				// example of modifying outgoing traffic
				for i, msg := range result.Response.Messages {
					if msg.MsgId() == network7.MsgCtrlConnect {
						var connect *messages7.CtrlConnect
						var ok bool
						if connect, ok = result.Response.Messages[i].(*messages7.CtrlConnect); ok {
							connect.Token = [4]byte{0xaa, 0xaa, 0xaa, 0xaa}
							result.Response.Messages[i] = connect
						}
					}
				}

				conn.Write(result.Response.Pack(client))
			}
		default:
			// do nothing
		}
	}
}

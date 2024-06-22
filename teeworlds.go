package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/teeworlds-go/teeworlds/messages7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/protocol7"
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

func main() {
	ch := make(chan []byte, maxPacksize)

	serverIp := "127.0.0.1"
	serverPort := 8303

	if len(os.Args) > 1 {
		if os.Args[1][0] == '-' {
			fmt.Println("usage: ./teeworlds [serverIp] [serverPort]")
			os.Exit(1)
		}
		serverIp = os.Args[1]
	}
	if len(os.Args) > 2 {
		var err error
		serverPort, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	conn, err := getConnection(serverIp, serverPort)
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
			packet := &protocol7.Packet{}
			err := packet.Unpack(msg)
			if err != nil {
				panic(err)
			}
			response := client.OnPacket(packet)
			if response != nil {
				// example of inspecting incoming trafic
				for _, msg := range packet.Messages {
					switch msg := msg.(type) {
					case *messages7.SvChat:
						// inspect incoming traffic
						fmt.Printf("got incoming chat msg: %s\n", msg.Message)
					default:
					}
				}

				// example of modifying outgoing traffic
				for i, msg := range response.Messages {
					switch msg := msg.(type) {
					case *messages7.SvChat:
						// inspect outgoing traffic
						fmt.Printf("got outgoing chat msg: %s\n", msg.Message)

						// change outgoing traffic
						msg.Message += " (edited by go)"
						packet.Messages[i] = msg
					default:
					}
				}

				if len(response.Messages) > 0 || response.Header.Flags.Resend {
					conn.Write(response.Pack(client))
				}
			}
		default:
			// do nothing
		}
	}
}

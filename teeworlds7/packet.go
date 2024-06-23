package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

func printUnknownMessage(msg messages7.NetMessage, msgType string) {
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

func (client *Client) processMessage(msg messages7.NetMessage, response *protocol7.Packet) bool {
	if msg.Header() == nil {
		// this is probably an unknown message
		fmt.Printf("warning ignoring msgId=%d because header is nil\n", msg.MsgId())
		return false
	}
	if msg.Header().Flags.Vital {
		client.Session.Ack++
	}

	if msg.System() {
		return client.processSystem(msg, response)
	}
	return client.processGame(msg, response)
}

func (client *Client) processPacket(packet *protocol7.Packet) error {
	if client.Callbacks.PacketIn != nil {
		client.Callbacks.PacketIn(packet)
	}

	response := client.Session.BuildResponse()

	if packet.Header.Flags.Control {
		if len(packet.Messages) != 1 {
			return fmt.Errorf("got control packet with %d messages.\n", len(packet.Messages))
		}

		msg := packet.Messages[0]
		// TODO: is this shadow nasty?
		switch msg := msg.(type) {
		case *messages7.CtrlKeepAlive:
			defaultAction := func() {
				fmt.Println("got keep alive")
			}
			if client.Callbacks.CtrlKeepAlive == nil {
				defaultAction()
			} else {
				client.Callbacks.CtrlKeepAlive(msg, defaultAction)
			}
		case *messages7.CtrlConnect:
			defaultAction := func() {
				fmt.Println("we got connect as a client. this should never happen lol.")
				fmt.Println("who is tryint to connect to us? We are not a server!")
			}
			if client.Callbacks.CtrlConnect == nil {
				defaultAction()
			} else {
				client.Callbacks.CtrlConnect(msg, defaultAction)
			}
		case *messages7.CtrlAccept:
			defaultAction := func() {
				fmt.Println("got accept")
				response.Messages = append(
					response.Messages,
					&messages7.Info{
						Version:       network7.NetVersion,
						Password:      "",
						ClientVersion: network7.ClientVersion,
					},
				)
				client.SendPacket(response)
			}
			if client.Callbacks.CtrlAccept == nil {
				defaultAction()
			} else {
				client.Callbacks.CtrlAccept(msg, defaultAction)
			}
		case *messages7.CtrlClose:
			defaultAction := func() {
				fmt.Printf("disconnected (%s)\n", msg.Reason)
			}
			if client.Callbacks.CtrlClose == nil {
				defaultAction()
			} else {
				client.Callbacks.CtrlClose(msg, defaultAction)
			}
		case *messages7.CtrlToken:
			defaultAction := func() {
				fmt.Printf("got server token %x\n", msg.Token)
				client.Session.ServerToken = msg.Token
				response.Header.Token = msg.Token
				response.Messages = append(
					response.Messages,
					&messages7.CtrlConnect{
						Token: client.Session.ClientToken,
					},
				)
				client.SendPacket(response)
			}
			if client.Callbacks.CtrlToken == nil {
				defaultAction()
			} else {
				client.Callbacks.CtrlToken(msg, defaultAction)
			}
		case *messages7.Unknown:
			defaultAction := func() {
				printUnknownMessage(msg, "unknown control")
			}
			if client.Callbacks.MsgUnknown == nil {
				defaultAction()
			} else {
				client.Callbacks.MsgUnknown(msg, defaultAction)
			}
			return fmt.Errorf("unknown control message: %d\n", msg.MsgId())
		default:
			return fmt.Errorf("unprocessed control message: %d\n", msg.MsgId())
		}
		return nil
	}

	for _, msg := range packet.Messages {
		client.processMessage(msg, response)
	}

	if len(response.Messages) > 0 || response.Header.Flags.Resend {
		client.SendPacket(response)
	}
	return nil
}

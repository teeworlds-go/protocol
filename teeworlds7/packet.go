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
	for _, callback := range client.Callbacks.PacketIn {
		if callback(packet) == false {
			return nil
		}
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
			userMsgCallback(client.Callbacks.CtrlKeepAlive, msg, func() {
				fmt.Println("got keep alive")
			})
		case *messages7.CtrlConnect:
			userMsgCallback(client.Callbacks.CtrlConnect, msg, func() {
				fmt.Println("we got connect as a client. this should never happen lol.")
				fmt.Println("who is tryint to connect to us? We are not a server!")
			})
		case *messages7.CtrlAccept:
			userMsgCallback(client.Callbacks.CtrlAccept, msg, func() {
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
			})
		case *messages7.CtrlClose:
			userMsgCallback(client.Callbacks.CtrlClose, msg, func() {
				fmt.Printf("disconnected (%s)\n", msg.Reason)
			})
		case *messages7.CtrlToken:
			userMsgCallback(client.Callbacks.CtrlToken, msg, func() {
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
			})
		case *messages7.Unknown:
			userMsgCallback(client.Callbacks.MsgUnknown, msg, func() {
				printUnknownMessage(msg, "unknown control")
			})
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

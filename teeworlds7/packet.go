package teeworlds7

import (
	"fmt"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
)

// This error is set as cancel cause in case that we are disconnected from the server
// This makes the error tangible for the user
// It can be checked with errors.As
type DisconnectError struct {
	// Reason is the reason why we were disconnected
	Reason string
}

func (e DisconnectError) Error() string {
	return fmt.Sprintf("disconnected: %s", e.Reason)
}

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

func (client *Client) processMessage(msg messages7.NetMessage, response *protocol7.Packet) (process bool, err error) {
	if msg.Header() == nil {
		// this is probably an unknown message
		fmt.Printf("warning ignoring msgId=%d because header is nil\n", msg.MsgId())
		return false, nil
	}
	if msg.Header().Flags.Vital {
		client.Session.Ack++
	}

	if msg.System() {
		return client.processSystem(msg, response)
	}
	return client.processGame(msg, response)
}

func (client *Client) processPacket(packet *protocol7.Packet) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to process packet: %w", err)
		}
	}()

	for _, callback := range client.Callbacks.PacketIn {
		if !callback(packet) {
			return nil
		}
	}

	response := client.Session.BuildResponse()

	if packet.Header.Flags.Control {
		if len(packet.Messages) != 1 {
			// TODO: implement resends for when 0 messages were received
			return fmt.Errorf("got control packet with %d messages", len(packet.Messages))
		}

		msg := packet.Messages[0]
		// TODO: is this shadow nasty?
		switch msg := msg.(type) {
		case *messages7.CtrlKeepAlive:
			err = userMsgCallback(client.Callbacks.CtrlKeepAlive, msg, func() error {
				fmt.Println("got keep alive")
				return nil
			})
		case *messages7.CtrlConnect:
			err = userMsgCallback(client.Callbacks.CtrlConnect, msg, func() error {
				fmt.Println("we got connect as a client. this should never happen lol.")
				fmt.Println("who is trying to connect to us? We are not a server!")
				return nil
			})
		case *messages7.CtrlAccept:
			err = userMsgCallback(client.Callbacks.CtrlAccept, msg, func() error {
				fmt.Println("got accept")
				response.Messages = append(
					response.Messages,
					&messages7.Info{
						Version:       network7.NetVersion,
						Password:      "",
						ClientVersion: network7.ClientVersion,
					},
				)
				return client.SendPacket(response)
			})
		case *messages7.CtrlClose:
			err = userMsgCallback(client.Callbacks.CtrlClose, msg, func() error {
				client.CancelCause(DisconnectError{Reason: msg.Reason})
				fmt.Printf("disconnected (%s)\n", msg.Reason)
				return nil
			})
		case *messages7.CtrlToken:
			err = userMsgCallback(client.Callbacks.CtrlToken, msg, func() error {
				fmt.Printf("got server token %x\n", msg.Token)
				client.Session.ServerToken = msg.Token
				response.Header.Token = msg.Token
				response.Messages = append(
					response.Messages,
					&messages7.CtrlConnect{
						Token: client.Session.ClientToken,
					},
				)
				return client.SendPacket(response)
			})
		case *messages7.Unknown:
			err = userMsgCallback(client.Callbacks.MsgUnknown, msg, func() error {
				printUnknownMessage(msg, "unknown control")
				return nil
			})
			if err != nil {
				return fmt.Errorf("unknown control message: %d: %w", msg.MsgId(), err)
			}

			return fmt.Errorf("unknown control message: %d", msg.MsgId())
		default:
			return fmt.Errorf("unprocessed control message: %d", msg.MsgId())
		}
		return err
	}

	for _, msg := range packet.Messages {
		_, err = client.processMessage(msg, response)
		if err != nil {
			return err
		}
	}

	if len(response.Messages) > 0 || response.Header.Flags.Resend {
		return client.SendPacket(response)
	}
	return nil
}

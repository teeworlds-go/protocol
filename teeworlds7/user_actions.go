package teeworlds7

import (
	"fmt"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
)

// ----------------------------
// low level access for experts
// ----------------------------

func (client *Client) SendPacket(packet *protocol7.Packet) error {
	if packet.Header.Flags.Resend == false && len(packet.Messages) == 0 && len(client.QueuedMessages) == 0 {
		return fmt.Errorf("Failed to send packet: payload is empty.")
	}

	gotNet := false
	numCtrlMsgs := 0

	for _, msg := range packet.Messages {
		if msg.MsgType() == network7.TypeControl {
			numCtrlMsgs++
		} else if msg.MsgType() == network7.TypeNet {
			gotNet = true
		} else {
			return fmt.Errorf("Failed to send packet: only game, system and control messages are supported.")
		}
	}

	if gotNet && numCtrlMsgs > 0 {
		return fmt.Errorf("Failed to send packet: can not mix control messages with others.")
	}

	if numCtrlMsgs > 1 {
		// TODO: should this automatically split it up into multiple packets?
		return fmt.Errorf("Failed to send packet: can only send one control message at a time.")
	}

	// If the user queued a game message and then sends a control message
	// before the queue got processed we send two packets in the correct order
	// For example in this case:
	//
	// client.SendChat("bye") // queue game chunk
	// client.Disconnect() // SendPacket(ctrl) -> first flush out the game chunk packet then send the control packet
	//
	if numCtrlMsgs > 0 && len(client.QueuedMessages) > 0 {
		// TODO: we could apply compression here
		// flushPacket.Header.Flags.Compression = true

		flushPacket := client.Session.BuildResponse()
		client.SendPacket(flushPacket)
	}

	for _, queuedChunk := range client.QueuedMessages {
		// TODO: check if we exceed packet size and only put in as many chunks as we can
		//       also use a more performant queue implementation then if we unshift it partially
		//       popping of one element from the queue should not reallocate the entire queued messages slice
		packet.Messages = append(packet.Messages, queuedChunk)
	}
	client.QueuedMessages = nil

	packet.Messages = client.registerMessagesCallbacks(packet.Messages)

	// slog.Info("after filter got messages", "len", len(packet.Messages))

	// for i, msg := range filteredMessages {
	// 	slog.Info("about to xxxx pack w msg", "i", i, "msg", msg)
	// }

	// for i, msg := range packet.Messages {
	// 	slog.Info("about to send pack w msg", "i", i, "msg", msg)
	// }

	for _, callback := range client.Callbacks.PacketOut {
		if callback(packet) == false {
			return nil
		}
	}

	client.LastSend = time.Now()
	client.Conn.Write(packet.Pack(&client.Session))
	return nil
}

// WARNING! this is does not send chat messages
// this sends a network chunk and is for expert users
//
// if you want to send a chat message use SendChat()
func (client *Client) SendMessage(msg messages7.NetMessage) {
	if msg.MsgType() == network7.TypeControl {
		packet := client.Session.BuildResponse()
		packet.Header.Flags.Control = true
		packet.Messages = append(packet.Messages, msg)
		client.SendPacket(packet)
		return
	}
	if msg.MsgType() == network7.TypeConnless {
		// TODO: connless
		panic("connless messages are not supported yet")
	}

	client.QueuedMessages = append(client.QueuedMessages, msg)
}

// ----------------------------
// high level actions
// ----------------------------

// Example of walking left
//
//	client.Game.Input.Direction = -1
//	client.SendInput()
//
// see also:
//
//	Right()
//	Left()
//	Stop()
//	Jump()
//	Fire()
//	Hook()
//	Aim(x, y)
func (client *Client) SendInput() {
	client.SendMessage(client.Game.Input)
}

func (client *Client) Right() {
	client.Game.Input.Direction = 1
	// client.SendInput()
}

func (client *Client) Left() {
	client.Game.Input.Direction = -1
	// client.SendInput()
}

func (client *Client) Stop() {
	client.Game.Input.Direction = 0
	// client.SendInput()
}

func (client *Client) Jump() {
	client.Game.Input.Jump = 1
	// client.SendInput()
}

func (client *Client) Hook() {
	client.Game.Input.Hook = 1
	// client.SendInput()
}

func (client *Client) Fire() {
	// TODO: fire is weird do we ever have to reset or mask it or something?
	client.Game.Input.Fire++
	// client.SendInput()
}

func (client *Client) Aim(x int, y int) {
	client.Game.Input.TargetX = x
	client.Game.Input.TargetY = y
	// client.SendInput()
}

// see also SendWhisper()
// see also SendChatTeam()
func (client *Client) SendChat(msg string) {
	client.SendMessage(
		&messages7.ClSay{
			Mode:     network7.ChatAll,
			Message:  msg,
			TargetId: -1,
		},
	)
}

// see also SendWhisper()
// see also SendChat()
func (client *Client) SendChatTeam(msg string) {
	client.SendMessage(
		&messages7.ClSay{
			Mode:     network7.ChatTeam,
			Message:  msg,
			TargetId: -1,
		},
	)
}

// see also SendChat()
// see also SendChatTeam()
func (client *Client) SendWhisper(targetId int, msg string) {
	client.SendMessage(
		&messages7.ClSay{
			Mode:     network7.ChatWhisper,
			Message:  msg,
			TargetId: targetId,
		},
	)
}

func (client *Client) SendKeepAlive() {
	client.SendMessage(&messages7.CtrlKeepAlive{})
}

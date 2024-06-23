package teeworlds7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

// ----------------------------
// low level access for experts
// ----------------------------

func (client *Client) SendPacket(packet *protocol7.Packet) {

	// TODO: append queued messages to packet messages here

	if client.Callbacks.PacketOut != nil {
		client.Callbacks.PacketOut(packet)
	}
	client.Conn.Write(packet.Pack(&client.Session))
}

// WARNING! this is does not send chat messages
// this sends a network chunk and is for expert users
//
// if you want to send a chat message use SendChat()
func (client *Client) SendMessage(msg messages7.NetMessage) {
	// TODO: set vital header and stuff
	client.QueuedMessages = append(client.QueuedMessages, msg)
}

// ----------------------------
// high level actions
// ----------------------------

// see also SendWhisper()
// see also SendChatTeam()
func (client *Client) SendChat(msg string) {
	client.SendMessage(
		&messages7.SvChat{
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
		&messages7.SvChat{
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
		&messages7.SvChat{
			Mode:     network7.ChatWhisper,
			Message:  msg,
			TargetId: targetId,
		},
	)
}

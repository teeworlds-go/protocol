package teeworlds7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
	"github.com/teeworlds-go/go-teeworlds-protocol/snapshot7"
)

// --------------------------------
// special cases
// --------------------------------

// if not implemented by the user the application might throw and exit
//
// return false to drop the error
// return true if you did not handle the error and it should be passed on (will crash unless someone else catches it)
func (client *Client) OnError(callback func(err error) bool) {
	client.Callbacks.InternalError = append(client.Callbacks.InternalError, callback)
}

// inspect outgoing traffic
// and alter it before it gets sent to the server
//
// return false to drop the packet
func (client *Client) OnSend(callback func(packet *protocol7.Packet) bool) {
	client.Callbacks.PacketOut = append(client.Callbacks.PacketOut, callback)
}

// read incoming traffic
// and alter it before it hits the internal state machine
//
// return false to drop the packet
func (client *Client) OnPacket(callback func(packet *protocol7.Packet) bool) {
	client.Callbacks.PacketIn = append(client.Callbacks.PacketIn, callback)
}

func (client *Client) OnUnknown(callback func(msg *messages7.Unknown, defaultAction DefaultAction)) {
	client.Callbacks.MsgUnknown = append(client.Callbacks.MsgUnknown, callback)
}

// will be called when a snap, snap single or empty snapshot is received
// if you want to know which type of snapshot was received look at OnMsgSnap(), OnMsgSnapEmpty(), OnMsgSnapSingle(), OnMsgSnapSmall()
func (client *Client) OnSnapshot(callback func(snap *snapshot7.Snapshot, defaultAction DefaultAction)) {
	client.Callbacks.Snapshot = append(client.Callbacks.Snapshot, callback)
}

// --------------------------------
// control messages
// --------------------------------

func (client *Client) OnKeepAlive(callback func(msg *messages7.CtrlKeepAlive, defaultAction DefaultAction)) {
	client.Callbacks.CtrlKeepAlive = append(client.Callbacks.CtrlKeepAlive, callback)
}

// This is just misleading. It should never be called. This message is only received by the server.
// func (client *Client) OnCtrlConnect(callback func(msg *messages7.CtrlConnect, defaultAction DefaultAction)) {
// 	client.Callbacks.CtrlConnect = append(// 	client.Callbacks, callback)
// }

func (client *Client) OnAccept(callback func(msg *messages7.CtrlAccept, defaultAction DefaultAction)) {
	client.Callbacks.CtrlAccept = append(client.Callbacks.CtrlAccept, callback)
}

func (client *Client) OnDisconnect(callback func(msg *messages7.CtrlClose, defaultAction DefaultAction)) {
	client.Callbacks.CtrlClose = append(client.Callbacks.CtrlClose, callback)
}

func (client *Client) OnToken(callback func(msg *messages7.CtrlToken, defaultAction DefaultAction)) {
	client.Callbacks.CtrlToken = append(client.Callbacks.CtrlToken, callback)
}

// --------------------------------
// game messages
// --------------------------------

func (client *Client) OnMotd(callback func(msg *messages7.SvMotd, defaultAction DefaultAction)) {
	client.Callbacks.GameSvMotd = append(client.Callbacks.GameSvMotd, callback)
}

func (client *Client) OnBroadcast(callback func(msg *messages7.SvBroadcast, defaultAction DefaultAction)) {
	client.Callbacks.GameSvBroadcast = append(client.Callbacks.GameSvBroadcast, callback)
}

func (client *Client) OnChat(callback func(msg *messages7.SvChat, defaultAction DefaultAction)) {
	client.Callbacks.GameSvChat = append(client.Callbacks.GameSvChat, callback)
}

func (client *Client) OnTeam(callback func(msg *messages7.SvTeam, defaultAction DefaultAction)) {
	client.Callbacks.GameSvTeam = append(client.Callbacks.GameSvTeam, callback)
}

// --------------------------------
// system messages
// --------------------------------

func (client *Client) OnMapChange(callback func(msg *messages7.MapChange, defaultAction DefaultAction)) {
	client.Callbacks.SysMapChange = append(client.Callbacks.SysMapChange, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnap(callback func(msg *messages7.Snap, defaultAction DefaultAction)) {
	client.Callbacks.SysSnap = append(client.Callbacks.SysSnap, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnapEmpty(callback func(msg *messages7.SnapEmpty, defaultAction DefaultAction)) {
	client.Callbacks.SysSnapEmpty = append(client.Callbacks.SysSnapEmpty, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnapSingle(callback func(msg *messages7.SnapSingle, defaultAction DefaultAction)) {
	client.Callbacks.SysSnapSingle = append(client.Callbacks.SysSnapSingle, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnapSmall(callback func(msg *messages7.SnapSmall, defaultAction DefaultAction)) {
	client.Callbacks.SysSnapSmall = append(client.Callbacks.SysSnapSmall, callback)
}

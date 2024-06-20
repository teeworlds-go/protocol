package messages7

// this message is shared between client and server
// but this implementation is assuming we are sending from a client

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type CtrlToken struct {
	Token [4]byte
}

func (msg CtrlToken) MsgId() int {
	return network7.MsgCtrlToken
}

func (msg CtrlToken) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg CtrlToken) System() bool {
	return false
}

func (msg CtrlToken) Vital() bool {
	return false
}

func (msg CtrlToken) Pack() []byte {
	return slices.Concat(
		[]byte{network7.MsgCtrlToken},
		msg.Token[:],
		[]byte{512: 0},
	)
}

// TODO: no idea if this works
func (msg *CtrlToken) Unpack(u *packer.Unpacker) {
	msg.Token = [4]byte(u.Data())
}

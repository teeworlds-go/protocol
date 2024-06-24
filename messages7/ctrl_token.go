package messages7

// this message is shared between client and server
// but this implementation is assuming we are sending from a client

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlToken struct {
	Token [4]byte
}

func (msg *CtrlToken) MsgId() int {
	return network7.MsgCtrlToken
}

func (msg *CtrlToken) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlToken) System() bool {
	return false
}

func (msg *CtrlToken) Vital() bool {
	return false
}

func (msg *CtrlToken) Pack() []byte {
	return slices.Concat(
		[]byte{network7.MsgCtrlToken},
		msg.Token[:],
		[]byte{512: 0},
	)
}

func (msg *CtrlToken) Unpack(u *packer.Unpacker) error {
	msg.Token = [4]byte(u.Rest())
	return nil
}

func (msg *CtrlToken) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlToken) SetHeader(header *chunk7.ChunkHeader) {
}

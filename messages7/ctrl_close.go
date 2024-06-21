package messages7

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type CtrlClose struct {
	Reason string
}

func (msg CtrlClose) MsgId() int {
	return network7.MsgCtrlClose
}

func (msg CtrlClose) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg CtrlClose) System() bool {
	return false
}

func (msg CtrlClose) Vital() bool {
	return false
}

func (msg CtrlClose) Pack() []byte {
	return slices.Concat(
		[]byte{network7.MsgCtrlClose},
		packer.PackStr(msg.Reason),
	)
}

func (msg *CtrlClose) Unpack(u *packer.Unpacker) {
	// TODO: sanitize
	msg.Reason = u.GetString()
}

func (msg *CtrlClose) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlClose) SetHeader(header *chunk7.ChunkHeader) {
}

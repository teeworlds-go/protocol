package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlConnect struct {
	Token [4]byte
}

func (msg *CtrlConnect) MsgId() int {
	return network7.MsgCtrlConnect
}

func (msg *CtrlConnect) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlConnect) System() bool {
	return false
}

func (msg *CtrlConnect) Vital() bool {
	return false
}

func (msg *CtrlConnect) Pack() []byte {
	return slices.Concat(
		[]byte{network7.MsgCtrlConnect},
		msg.Token[:],
		[]byte{512: 0},
	)
}

func (msg *CtrlConnect) Unpack(u *packer.Unpacker) error {
	msg.Token = [4]byte(u.Rest())
	return nil
}

func (msg *CtrlConnect) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlConnect) SetHeader(header *chunk7.ChunkHeader) {
}

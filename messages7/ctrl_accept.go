package messages7

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type CtrlAccept struct {
	Token [4]byte
}

func (msg CtrlAccept) MsgId() int {
	return network7.MsgCtrlAccept
}

func (msg CtrlAccept) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg CtrlAccept) System() bool {
	return false
}

func (msg CtrlAccept) Vital() bool {
	return false
}

func (msg CtrlAccept) Pack() []byte {
	return slices.Concat(
		[]byte{network7.MsgCtrlAccept},
		msg.Token[:],
		[]byte{512: 0},
	)
}

func (msg *CtrlAccept) Unpack(u *packer.Unpacker) {
}

func (msg *CtrlAccept) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlAccept) SetHeader(header *chunk7.ChunkHeader) {
}

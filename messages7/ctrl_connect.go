package messages7

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type CtrlConnect struct {
	Token [4]byte
}

func (msg CtrlConnect) MsgId() int {
	return network7.MsgCtrlConnect
}

func (msg CtrlConnect) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg CtrlConnect) System() bool {
	return false
}

func (msg CtrlConnect) Vital() bool {
	return false
}

func (msg CtrlConnect) Pack() []byte {
	return slices.Concat(
		[]byte{network7.MsgCtrlConnect},
		msg.Token[:],
		[]byte{512: 0},
	)
}

// TODO: no idea if this works
func (msg *CtrlConnect) Unpack(u *packer.Unpacker) {
	msg.Token = [4]byte(u.Data())
}

func (msg *CtrlConnect) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlConnect) SetHeader(header *chunk7.ChunkHeader) {
}

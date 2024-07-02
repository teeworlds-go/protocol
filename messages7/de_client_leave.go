package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

// only used for demos
// never send over the network
type DeClientLeave struct {
	ChunkHeader *chunk7.ChunkHeader

	Name     string
	ClientId int
	Reason   string
}

func (msg *DeClientLeave) MsgId() int {
	return network7.MsgGameDeClientLeave
}

func (msg *DeClientLeave) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *DeClientLeave) System() bool {
	return false
}

func (msg *DeClientLeave) Vital() bool {
	return true
}

func (msg *DeClientLeave) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
		packer.PackInt(msg.ClientId),
		packer.PackStr(msg.Reason),
	)
}

func (msg *DeClientLeave) Unpack(u *packer.Unpacker) error {
	var err error
	msg.Name, err = u.GetString()
	if err != nil {
		return err
	}
	msg.ClientId = u.GetInt()
	msg.Reason, err = u.GetString()
	if err != nil {
		return err
	}

	return nil
}

func (msg *DeClientLeave) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *DeClientLeave) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

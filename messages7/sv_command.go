package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type ClCommand struct {
	ChunkHeader *chunk7.ChunkHeader

	Name      string
	Arguments string
}

func (msg *ClCommand) MsgId() int {
	return network7.MsgGameClCommand
}

func (msg *ClCommand) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClCommand) System() bool {
	return false
}

func (msg *ClCommand) Vital() bool {
	return true
}

func (msg *ClCommand) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
		packer.PackStr(msg.Arguments),
	)
}

func (msg *ClCommand) Unpack(u *packer.Unpacker) error {
	var err error
	msg.Name, err = u.GetString()
	if err != nil {
		return err
	}
	msg.Arguments, err = u.GetString()
	if err != nil {
		return err
	}

	return nil
}

func (msg *ClCommand) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClCommand) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

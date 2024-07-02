package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type ClSay struct {
	ChunkHeader *chunk7.ChunkHeader

	Mode     network7.ChatMode
	TargetId int
	Message  string
}

func (msg *ClSay) MsgId() int {
	return network7.MsgGameClSay
}

func (msg *ClSay) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClSay) System() bool {
	return false
}

func (msg *ClSay) Vital() bool {
	return true
}

func (msg *ClSay) Pack() []byte {
	return slices.Concat(
		packer.PackInt(int(msg.Mode)),
		packer.PackInt(msg.TargetId),
		packer.PackStr(msg.Message),
	)
}

func (msg *ClSay) Unpack(u *packer.Unpacker) error {
	msg.Mode = network7.ChatMode(u.GetInt())
	msg.TargetId = u.GetInt()
	msg.Message, _ = u.GetString()

	return nil
}

func (msg *ClSay) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClSay) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

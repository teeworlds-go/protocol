package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvChat struct {
	ChunkHeader *chunk7.ChunkHeader

	Mode     network7.ChatMode
	ClientId int
	TargetId int
	Message  string
}

func (msg *SvChat) MsgId() int {
	return network7.MsgGameSvChat
}

func (msg *SvChat) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvChat) System() bool {
	return false
}

func (msg *SvChat) Vital() bool {
	return true
}

func (msg *SvChat) Pack() []byte {
	return slices.Concat(
		packer.PackInt(int(msg.Mode)),
		packer.PackInt(msg.ClientId),
		packer.PackInt(msg.TargetId),
		packer.PackStr(msg.Message),
	)
}

func (msg *SvChat) Unpack(u *packer.Unpacker) error {
	msg.Mode = network7.ChatMode(u.GetInt())
	msg.ClientId = u.GetInt()
	msg.TargetId = u.GetInt()
	msg.Message, _ = u.GetString()

	return nil
}

func (msg *SvChat) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvChat) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type SvChat struct {
	header *chunk7.ChunkHeader

	Mode     int
	ClientId int
	TargetId int
	Message  string
}

func (msg SvChat) MsgId() int {
	return network7.MsgGameSvChat
}

func (msg SvChat) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg SvChat) System() bool {
	return false
}

func (msg SvChat) Vital() bool {
	return true
}

func (msg SvChat) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.Mode),
		packer.PackInt(msg.ClientId),
		packer.PackInt(msg.TargetId),
		packer.PackStr(msg.Message),
	)
}

func (msg *SvChat) Unpack(u *packer.Unpacker) {
	msg.Mode = u.GetInt()
	msg.ClientId = u.GetInt()
	msg.TargetId = u.GetInt()
	msg.Message = u.GetString()
}

func (msg *SvChat) Header() *chunk7.ChunkHeader {
	return msg.header
}

func (msg *SvChat) SetHeader(header *chunk7.ChunkHeader) {
	msg.header = header
}

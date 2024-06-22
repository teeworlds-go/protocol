package messages7

import (
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type SvMotd struct {
	ChunkHeader *chunk7.ChunkHeader

	Message string
}

func (msg SvMotd) MsgId() int {
	return network7.MsgGameSvMotd
}

func (msg SvMotd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg SvMotd) System() bool {
	return false
}

func (msg SvMotd) Vital() bool {
	return true
}

func (msg SvMotd) Pack() []byte {
	return packer.PackStr(msg.Message)
}

func (msg *SvMotd) Unpack(u *packer.Unpacker) {
	msg.Message = u.GetString()
}

func (msg *SvMotd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvMotd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

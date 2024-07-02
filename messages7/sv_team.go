package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvTeam struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId     int
	Silent       bool
	CooldownTick int
}

func (msg *SvTeam) MsgId() int {
	return network7.MsgGameSvTeam
}

func (msg *SvTeam) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvTeam) System() bool {
	return false
}

func (msg *SvTeam) Vital() bool {
	return true
}

func (msg *SvTeam) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.ClientId),
		packer.PackBool(msg.Silent),
		packer.PackInt(msg.CooldownTick),
	)
}

func (msg *SvTeam) Unpack(u *packer.Unpacker) error {
	msg.ClientId = u.GetInt()
	msg.Silent = u.GetInt() != 0
	msg.CooldownTick = u.GetInt()
	return nil
}

func (msg *SvTeam) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvTeam) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

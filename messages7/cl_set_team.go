package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type ClSetTeam struct {
	ChunkHeader *chunk7.ChunkHeader

	Team network7.GameTeam
}

func (msg *ClSetTeam) MsgId() int {
	return network7.MsgGameClSetTeam
}

func (msg *ClSetTeam) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClSetTeam) System() bool {
	return false
}

func (msg *ClSetTeam) Vital() bool {
	return true
}

func (msg *ClSetTeam) Pack() []byte {
	return slices.Concat(
		packer.PackInt(int(msg.Team)),
	)
}

func (msg *ClSetTeam) Unpack(u *packer.Unpacker) error {
	msg.Team = network7.GameTeam(u.GetInt())

	return nil
}

func (msg *ClSetTeam) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClSetTeam) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

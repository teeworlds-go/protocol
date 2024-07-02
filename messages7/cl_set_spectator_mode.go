package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type ClSetSpectatorMode struct {
	ChunkHeader *chunk7.ChunkHeader

	Mode        network7.Spec
	SpectatorId int
}

func (msg *ClSetSpectatorMode) MsgId() int {
	return network7.MsgGameClSetSpectatorMode
}

func (msg *ClSetSpectatorMode) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClSetSpectatorMode) System() bool {
	return false
}

func (msg *ClSetSpectatorMode) Vital() bool {
	return true
}

func (msg *ClSetSpectatorMode) Pack() []byte {
	return slices.Concat(
		packer.PackInt(int(msg.Mode)),
		packer.PackInt(msg.SpectatorId),
	)
}

func (msg *ClSetSpectatorMode) Unpack(u *packer.Unpacker) error {
	msg.Mode = network7.Spec(u.GetInt())
	msg.SpectatorId = u.GetInt()
	return nil
}

func (msg *ClSetSpectatorMode) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClSetSpectatorMode) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

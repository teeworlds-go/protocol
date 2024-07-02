package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvServerSettings struct {
	ChunkHeader *chunk7.ChunkHeader

	KickVote    bool
	KickMin     int
	SpecVote    bool
	TeamLock    bool
	TeamBalance bool
	PlayerSlots int
}

func (msg *SvServerSettings) MsgId() int {
	return network7.MsgGameSvServerSettings
}

func (msg *SvServerSettings) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvServerSettings) System() bool {
	return false
}

func (msg *SvServerSettings) Vital() bool {
	return true
}

func (msg *SvServerSettings) Pack() []byte {
	return slices.Concat(
		packer.PackBool(msg.KickVote),
		packer.PackInt(msg.KickMin),
		packer.PackBool(msg.SpecVote),
		packer.PackBool(msg.TeamLock),
		packer.PackBool(msg.TeamBalance),
		packer.PackInt(msg.PlayerSlots),
	)
}

func (msg *SvServerSettings) Unpack(u *packer.Unpacker) error {
	msg.KickVote = u.GetInt() != 0
	msg.KickMin = u.GetInt()
	msg.SpecVote = u.GetInt() != 0
	msg.TeamLock = u.GetInt() != 0
	msg.TeamBalance = u.GetInt() != 0
	msg.PlayerSlots = u.GetInt()

	return nil
}

func (msg *SvServerSettings) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvServerSettings) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

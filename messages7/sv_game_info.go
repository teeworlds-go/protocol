package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvGameInfo struct {
	ChunkHeader *chunk7.ChunkHeader

	// Gameflags can have any combination of those flag bits set:
	//    1 - GAMEFLAG_TEAMS
	//    2 - GAMEFLAG_FLAGS
	//    4 - GAMEFLAG_SURVIVAL
	//    8 - GAMEFLAG_RACE
	GameFlags    int
	ScoreLimit   int
	TimeLimit    int
	MatchNum     int
	MatchCurrent int
}

func (msg *SvGameInfo) MsgId() int {
	return network7.MsgGameSvGameInfo
}

func (msg *SvGameInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvGameInfo) System() bool {
	return false
}

func (msg *SvGameInfo) Vital() bool {
	return true
}

func (msg *SvGameInfo) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.GameFlags),
		packer.PackInt(msg.ScoreLimit),
		packer.PackInt(msg.TimeLimit),
		packer.PackInt(msg.MatchNum),
		packer.PackInt(msg.MatchCurrent),
	)
}

func (msg *SvGameInfo) Unpack(u *packer.Unpacker) error {
	msg.GameFlags = u.GetInt()
	msg.ScoreLimit = u.GetInt()
	msg.TimeLimit = u.GetInt()
	msg.MatchNum = u.GetInt()
	msg.MatchCurrent = u.GetInt()

	return nil
}

func (msg *SvGameInfo) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvGameInfo) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

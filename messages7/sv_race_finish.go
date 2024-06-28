package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvRaceFinish struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId       int
	Time           int
	Diff           int
	RecordPersonal int
	RecordServer   int
}

func (msg *SvRaceFinish) MsgId() int {
	return network7.MsgGameSvRaceFinish
}

func (msg *SvRaceFinish) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvRaceFinish) System() bool {
	return false
}

func (msg *SvRaceFinish) Vital() bool {
	return true
}

func (msg *SvRaceFinish) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.ClientId),
		packer.PackInt(msg.Time),
		packer.PackInt(msg.Diff),
		packer.PackInt(msg.RecordPersonal),
		packer.PackInt(msg.RecordServer),
	)
}

func (msg *SvRaceFinish) Unpack(u *packer.Unpacker) error {
	msg.ClientId = u.GetInt()
	msg.Time = u.GetInt()
	msg.Diff = u.GetInt()
	msg.RecordPersonal = u.GetInt()
	msg.RecordServer = u.GetInt()

	return nil
}

func (msg *SvRaceFinish) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvRaceFinish) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

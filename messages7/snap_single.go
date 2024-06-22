package messages7

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type SnapSingle struct {
	ChunkHeader *chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
	Crc       int
	PartSize  int
	Data      []byte
}

func (msg SnapSingle) MsgId() int {
	return network7.MsgSysSnapSingle
}

func (msg SnapSingle) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg SnapSingle) System() bool {
	return true
}

func (msg SnapSingle) Vital() bool {
	return false
}

func (msg SnapSingle) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.GameTick),
		packer.PackInt(msg.DeltaTick),
		packer.PackInt(msg.Crc),
		packer.PackInt(msg.PartSize),
		msg.Data[:],
	)
}

func (msg *SnapSingle) Unpack(u *packer.Unpacker) {
	msg.GameTick = u.GetInt()
	msg.DeltaTick = u.GetInt()
	msg.Crc = u.GetInt()
	msg.PartSize = u.GetInt()
	msg.Data = u.Rest()
}

func (msg *SnapSingle) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SnapSingle) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Snap struct {
	ChunkHeader *chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
	NumParts  int
	Part      int
	Crc       int
	PartSize  int
	Data      []byte
}

func (msg *Snap) MsgId() int {
	return network7.MsgSysSnap
}

func (msg *Snap) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *Snap) System() bool {
	return true
}

func (msg *Snap) Vital() bool {
	return false
}

func (msg *Snap) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.GameTick),
		packer.PackInt(msg.DeltaTick),
		packer.PackInt(msg.NumParts),
		packer.PackInt(msg.Part),
		packer.PackInt(msg.Crc),
		packer.PackInt(msg.PartSize),
		msg.Data,
	)
}

func (msg *Snap) Unpack(u *packer.Unpacker) error {
	msg.GameTick = u.GetInt()
	msg.DeltaTick = u.GetInt()
	msg.NumParts = u.GetInt()
	msg.Part = u.GetInt()
	msg.Crc = u.GetInt()
	msg.PartSize = u.GetInt()
	msg.Data = u.Rest()
	return nil
}

func (msg *Snap) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *Snap) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

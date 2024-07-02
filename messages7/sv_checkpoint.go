package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvCheckpoint struct {
	ChunkHeader *chunk7.ChunkHeader

	Diff int
}

func (msg *SvCheckpoint) MsgId() int {
	return network7.MsgGameSvCheckpoint
}

func (msg *SvCheckpoint) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvCheckpoint) System() bool {
	return false
}

func (msg *SvCheckpoint) Vital() bool {
	return true
}

func (msg *SvCheckpoint) Pack() []byte {
	return packer.PackInt(msg.Diff)
}

func (msg *SvCheckpoint) Unpack(u *packer.Unpacker) error {
	msg.Diff = u.GetInt()

	return nil
}

func (msg *SvCheckpoint) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvCheckpoint) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

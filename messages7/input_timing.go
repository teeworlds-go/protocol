package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type InputTiming struct {
	ChunkHeader *chunk7.ChunkHeader

	IntendedPredTick int
	TimeLeft         int
}

func (msg InputTiming) MsgId() int {
	return network7.MsgSysInputTiming
}

func (msg InputTiming) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg InputTiming) System() bool {
	return true
}

func (msg InputTiming) Vital() bool {
	return false
}

func (msg InputTiming) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.IntendedPredTick),
		packer.PackInt(msg.TimeLeft),
	)
}

func (msg *InputTiming) Unpack(u *packer.Unpacker) {
	msg.IntendedPredTick = u.GetInt()
	msg.TimeLeft = u.GetInt()
}

func (msg *InputTiming) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *InputTiming) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

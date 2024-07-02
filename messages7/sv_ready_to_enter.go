package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvReadyToEnter struct {
	ChunkHeader *chunk7.ChunkHeader
}

func (msg *SvReadyToEnter) MsgId() int {
	return network7.MsgGameSvReadyToEnter
}

func (msg *SvReadyToEnter) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvReadyToEnter) System() bool {
	return false
}

func (msg *SvReadyToEnter) Vital() bool {
	return true
}

func (msg *SvReadyToEnter) Pack() []byte {
	return []byte{}
}

func (msg *SvReadyToEnter) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *SvReadyToEnter) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvReadyToEnter) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type SvExtraProjectile struct {
	ChunkHeader *chunk7.ChunkHeader
}

func (msg *SvExtraProjectile) MsgId() int {
	return network7.MsgGameSvExtraProjectile
}

func (msg *SvExtraProjectile) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvExtraProjectile) System() bool {
	return false
}

func (msg *SvExtraProjectile) Vital() bool {
	return true
}

func (msg *SvExtraProjectile) Pack() []byte {
	return []byte{}
}

func (msg *SvExtraProjectile) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *SvExtraProjectile) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvExtraProjectile) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type ClKill struct {
	ChunkHeader *chunk7.ChunkHeader
}

func (msg *ClKill) MsgId() int {
	return network7.MsgGameClKill
}

func (msg *ClKill) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClKill) System() bool {
	return false
}

func (msg *ClKill) Vital() bool {
	return true
}

func (msg *ClKill) Pack() []byte {
	return []byte{}
}

func (msg *ClKill) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *ClKill) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClKill) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

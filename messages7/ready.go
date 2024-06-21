package messages7

import (
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type Ready struct {
	header *chunk7.ChunkHeader
}

func (msg Ready) MsgId() int {
	return network7.MsgSysReady
}

func (msg Ready) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg Ready) System() bool {
	return true
}

func (msg Ready) Vital() bool {
	return true
}

func (msg Ready) Pack() []byte {
	return []byte{}
}

func (msg *Ready) Unpack(u *packer.Unpacker) {
}

func (msg *Ready) Header() *chunk7.ChunkHeader {
	return msg.header
}

func (msg *Ready) SetHeader(header *chunk7.ChunkHeader) {
	msg.header = header
}

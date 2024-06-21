package messages7

import (
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type Info struct {
	header *chunk7.ChunkHeader
}

func (msg Info) MsgId() int {
	return network7.MsgSysInfo
}

func (info Info) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg Info) System() bool {
	return true
}

func (msg Info) Vital() bool {
	return true
}

func (msg Info) Pack() []byte {
	return []byte{
		0x30, 0x2E, 0x37, 0x20, 0x38, 0x30, 0x32, 0x66,
		0x31, 0x62, 0x65, 0x36, 0x30, 0x61, 0x30, 0x35, 0x36, 0x36, 0x35, 0x66,
		0x00, 0x6D, 0x79, 0x5F, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6F, 0x72, 0x64,
		0x5F, 0x31, 0x32, 0x33, 0x00, 0x85, 0x1C, 0x00,
	}
}

func (msg *Info) Unpack(u *packer.Unpacker) {
	// TODO: implement
	panic("not implemented")
}

func (msg *Info) Header() *chunk7.ChunkHeader {
	return msg.header
}

func (msg *Info) SetHeader(header *chunk7.ChunkHeader) {
	msg.header = header
}

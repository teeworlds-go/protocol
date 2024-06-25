package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type MaplistEntryAdd struct {
	ChunkHeader *chunk7.ChunkHeader

	MapName string
}

func (msg *MaplistEntryAdd) MsgId() int {
	return network7.MsgSysMaplistEntryAdd
}

func (msg *MaplistEntryAdd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *MaplistEntryAdd) System() bool {
	return true
}

func (msg *MaplistEntryAdd) Vital() bool {
	return true
}

func (msg *MaplistEntryAdd) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.MapName),
	)
}

func (msg *MaplistEntryAdd) Unpack(u *packer.Unpacker) error {
	msg.MapName, _ = u.GetString()
	return nil
}

func (msg *MaplistEntryAdd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *MaplistEntryAdd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

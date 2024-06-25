package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type MapChange struct {
	ChunkHeader *chunk7.ChunkHeader

	Name                        string
	Crc                         int
	Size                        int
	NumResponseChunksPerRequest int
	ChunkSize                   int
	Sha256                      [32]byte
}

func (msg *MapChange) MsgId() int {
	return network7.MsgSysMapChange
}

func (msg *MapChange) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *MapChange) System() bool {
	return true
}

func (msg *MapChange) Vital() bool {
	return true
}

func (msg *MapChange) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
		packer.PackInt(msg.Crc),
		packer.PackInt(msg.Size),
		packer.PackInt(msg.NumResponseChunksPerRequest),
		packer.PackInt(msg.ChunkSize),
		msg.Sha256[:],
	)
}

func (msg *MapChange) Unpack(u *packer.Unpacker) error {
	msg.Name, _ = u.GetString()
	msg.Crc = u.GetInt()
	msg.Size = u.GetInt()
	msg.NumResponseChunksPerRequest = u.GetInt()
	msg.ChunkSize = u.GetInt()
	msg.Sha256 = [32]byte(u.Rest())
	return nil
}

func (msg *MapChange) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *MapChange) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

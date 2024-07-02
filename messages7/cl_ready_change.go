package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

// message to request game pause and unpause
// will pause the game for everyone if one person sends it and if the server allows it
// will unpause the game if everyone send it
type ClReadyChange struct {
	ChunkHeader *chunk7.ChunkHeader
}

func (msg *ClReadyChange) MsgId() int {
	return network7.MsgGameClReadyChange
}

func (msg *ClReadyChange) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClReadyChange) System() bool {
	return false
}

func (msg *ClReadyChange) Vital() bool {
	return true
}

func (msg *ClReadyChange) Pack() []byte {
	return []byte{}
}

func (msg *ClReadyChange) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *ClReadyChange) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClReadyChange) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type ClEmoticon struct {
	ChunkHeader *chunk7.ChunkHeader

	Emoticon network7.Emoticon
}

func (msg *ClEmoticon) MsgId() int {
	return network7.MsgGameClEmoticon
}

func (msg *ClEmoticon) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClEmoticon) System() bool {
	return false
}

func (msg *ClEmoticon) Vital() bool {
	return true
}

func (msg *ClEmoticon) Pack() []byte {
	return packer.PackInt(int(msg.Emoticon))
}

func (msg *ClEmoticon) Unpack(u *packer.Unpacker) error {
	msg.Emoticon = network7.Emoticon(u.GetInt())

	return nil
}

func (msg *ClEmoticon) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClEmoticon) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

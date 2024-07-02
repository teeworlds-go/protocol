package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvEmoticon struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId int
	Emoticon network7.Emoticon
}

func (msg *SvEmoticon) MsgId() int {
	return network7.MsgGameSvEmoticon
}

func (msg *SvEmoticon) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvEmoticon) System() bool {
	return false
}

func (msg *SvEmoticon) Vital() bool {
	return true
}

func (msg *SvEmoticon) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.ClientId),
		packer.PackInt(int(msg.Emoticon)),
	)
}

func (msg *SvEmoticon) Unpack(u *packer.Unpacker) error {
	msg.ClientId = u.GetInt()
	msg.Emoticon = network7.Emoticon(u.GetInt())

	return nil
}

func (msg *SvEmoticon) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvEmoticon) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

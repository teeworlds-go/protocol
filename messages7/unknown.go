package messages7

import (
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type Unknown struct {
	ChunkHeader *chunk7.ChunkHeader

	// contains entire raw message
	// including message id and chunk header
	// can either be a control message or a game/system message
	Data []byte
	Type network7.MsgType
}

func (msg Unknown) MsgId() int {
	msgId := packer.UnpackInt(msg.Data)
	if msg.Type == network7.TypeControl {
		return msgId
	}
	msgId >>= 1
	return msgId
}

func (msg Unknown) MsgType() network7.MsgType {
	return msg.Type
}

func (msg Unknown) System() bool {
	msgId := packer.UnpackInt(msg.Data)
	if msg.Type == network7.TypeControl {
		return false
	}
	sys := msgId&1 != 0
	return sys
}

func (msg Unknown) Vital() bool {
	// TODO: check is not ctrl and then unpack Data
	panic("not implemented yet")
}

func (msg Unknown) Pack() []byte {
	return msg.Data
}

func (msg *Unknown) Unpack(u *packer.Unpacker) {
	msg.Data = u.Rest()
}

func (msg *Unknown) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *Unknown) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvClientDrop struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId int
	Reason   string
	Silent   bool
}

func (msg *SvClientDrop) MsgId() int {
	return network7.MsgGameSvClientDrop
}

func (msg *SvClientDrop) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvClientDrop) System() bool {
	return false
}

func (msg *SvClientDrop) Vital() bool {
	return true
}

func (msg *SvClientDrop) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.ClientId),
		packer.PackStr(msg.Reason),
		packer.PackBool(msg.Silent),
	)
}

func (msg *SvClientDrop) Unpack(u *packer.Unpacker) error {
	var err error
	msg.ClientId = u.GetInt()
	msg.Reason, err = u.GetString()
	if err != nil {
		return err
	}
	msg.Silent = u.GetInt() != 0

	return nil
}

func (msg *SvClientDrop) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvClientDrop) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

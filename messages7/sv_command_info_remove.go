package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvCommandInfoRemove struct {
	ChunkHeader *chunk7.ChunkHeader

	Name string
}

func (msg *SvCommandInfoRemove) MsgId() int {
	return network7.MsgGameSvCommandInfoRemove
}

func (msg *SvCommandInfoRemove) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvCommandInfoRemove) System() bool {
	return false
}

func (msg *SvCommandInfoRemove) Vital() bool {
	return true
}

func (msg *SvCommandInfoRemove) Pack() []byte {
	return packer.PackStr(msg.Name)
}

func (msg *SvCommandInfoRemove) Unpack(u *packer.Unpacker) error {
	var err error
	msg.Name, err = u.GetString()
	if err != nil {
		return err
	}

	return nil
}

func (msg *SvCommandInfoRemove) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvCommandInfoRemove) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvCommandInfo struct {
	ChunkHeader *chunk7.ChunkHeader

	Name       string
	ArgsFormat string
	HelpText   string
}

func (msg *SvCommandInfo) MsgId() int {
	return network7.MsgGameSvCommandInfo
}

func (msg *SvCommandInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvCommandInfo) System() bool {
	return false
}

func (msg *SvCommandInfo) Vital() bool {
	return true
}

func (msg *SvCommandInfo) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
		packer.PackStr(msg.ArgsFormat),
		packer.PackStr(msg.HelpText),
	)
}

func (msg *SvCommandInfo) Unpack(u *packer.Unpacker) error {
	var err error
	msg.Name, err = u.GetString()
	if err != nil {
		return err
	}
	msg.ArgsFormat, err = u.GetString()
	if err != nil {
		return err
	}
	msg.HelpText, err = u.GetString()
	if err != nil {
		return err
	}

	return nil
}

func (msg *SvCommandInfo) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvCommandInfo) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type RconCmdAdd struct {
	ChunkHeader *chunk7.ChunkHeader

	Name   string
	Help   string
	Params string
}

func (msg *RconCmdAdd) MsgId() int {
	return network7.MsgSysRconCmdAdd
}

func (msg *RconCmdAdd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconCmdAdd) System() bool {
	return true
}

func (msg *RconCmdAdd) Vital() bool {
	return true
}

func (msg *RconCmdAdd) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
		packer.PackStr(msg.Help),
		packer.PackStr(msg.Params),
	)
}

func (msg *RconCmdAdd) Unpack(u *packer.Unpacker) error {
	msg.Name, _ = u.GetString()
	msg.Help, _ = u.GetString()
	msg.Params, _ = u.GetString()
	return nil
}

func (msg *RconCmdAdd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *RconCmdAdd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

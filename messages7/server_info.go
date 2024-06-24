package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type ServerInfo struct {
	ChunkHeader *chunk7.ChunkHeader

	Version     string
	Name        string
	Hostname    string
	MapName     string
	GameType    string
	Flags       int
	SkillLevel  int
	PlayerCount int
	PlayerSlots int
	ClientCount int
	MaxClients  int
}

func (msg *ServerInfo) MsgId() int {
	return network7.MsgSysServerInfo
}

func (msg *ServerInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ServerInfo) System() bool {
	return true
}

func (msg *ServerInfo) Vital() bool {
	return true
}

func (msg *ServerInfo) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Version),
		packer.PackStr(msg.Name),
		packer.PackStr(msg.Hostname),
		packer.PackStr(msg.MapName),
		packer.PackStr(msg.GameType),
		packer.PackInt(msg.Flags),
		packer.PackInt(msg.SkillLevel),
		packer.PackInt(msg.PlayerCount),
		packer.PackInt(msg.PlayerSlots),
		packer.PackInt(msg.ClientCount),
		packer.PackInt(msg.MaxClients),
	)
}

func (msg *ServerInfo) Unpack(u *packer.Unpacker) error {
	msg.Version = u.GetString()
	msg.Name = u.GetString()
	msg.Hostname = u.GetString()
	msg.MapName = u.GetString()
	msg.GameType = u.GetString()
	msg.Flags = u.GetInt()
	msg.SkillLevel = u.GetInt()
	msg.PlayerCount = u.GetInt()
	msg.PlayerSlots = u.GetInt()
	msg.ClientCount = u.GetInt()
	msg.MaxClients = u.GetInt()
	return nil
}

func (msg *ServerInfo) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ServerInfo) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

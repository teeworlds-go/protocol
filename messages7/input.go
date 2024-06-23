package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Input struct {
	ChunkHeader *chunk7.ChunkHeader

	AckGameTick    int
	PredictionTick int
	Size           int

	Direction    int
	TargetX      int
	TargetY      int
	Jump         int
	Fire         int
	Hook         int
	PlayerFlags  int
	WantedWeapon network7.Weapon
	NextWeapon   network7.Weapon
	PrevWeapon   network7.Weapon
}

func (msg Input) MsgId() int {
	return network7.MsgSysInput
}

func (msg Input) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg Input) System() bool {
	return true
}

func (msg Input) Vital() bool {
	return false
}

func (msg Input) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.Direction),
		packer.PackInt(msg.TargetX),
		packer.PackInt(msg.TargetY),
		packer.PackInt(msg.Jump),
		packer.PackInt(msg.Fire),
		packer.PackInt(msg.Hook),
		packer.PackInt(msg.PlayerFlags),
		packer.PackInt(int(msg.WantedWeapon)),
		packer.PackInt(int(msg.NextWeapon)),
		packer.PackInt(int(msg.PrevWeapon)),
	)
}

func (msg *Input) Unpack(u *packer.Unpacker) {
	msg.Direction = u.GetInt()
	msg.TargetX = u.GetInt()
	msg.TargetY = u.GetInt()
	msg.Jump = u.GetInt()
	msg.Fire = u.GetInt()
	msg.Hook = u.GetInt()
	msg.PlayerFlags = u.GetInt()
	msg.WantedWeapon = network7.Weapon(u.GetInt())
	msg.NextWeapon = network7.Weapon(u.GetInt())
	msg.PrevWeapon = network7.Weapon(u.GetInt())
}

func (msg *Input) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *Input) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

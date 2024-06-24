package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvWeaponPickup struct {
	ChunkHeader *chunk7.ChunkHeader

	Weapon network7.Weapon
}

func (msg *SvWeaponPickup) MsgId() int {
	return network7.MsgGameSvWeaponPickup
}

func (msg *SvWeaponPickup) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvWeaponPickup) System() bool {
	return false
}

func (msg *SvWeaponPickup) Vital() bool {
	return true
}

func (msg *SvWeaponPickup) Pack() []byte {
	return slices.Concat(
		packer.PackInt(int(msg.Weapon)),
	)
}

func (msg *SvWeaponPickup) Unpack(u *packer.Unpacker) error {
	msg.Weapon = network7.Weapon(u.GetInt())

	return nil
}

func (msg *SvWeaponPickup) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvWeaponPickup) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

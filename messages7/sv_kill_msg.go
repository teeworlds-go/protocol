package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvKillMsg struct {
	ChunkHeader *chunk7.ChunkHeader

	// Client ID of the killer.
	// Can be the same as the Victim.
	// For example on a selfkill with grenade but also when a tee dies in a spike (death tile) or falls out of the world.
	KillerId int

	// Client ID of the killed.
	VictimId int

	// Weapon the tee was killed with. Can be one of those:
	//
	// -3 network7.WeaponGame (team switching etc)
	// -2 network7.WeaponSelf (console kill command)
	// -1 network7.WeaponWorld (death tiles etc)
	//  0 network7.WeaponHammer
	//  1 network7.WeaponGun
	//  2 network7.WeaponShotgun
	//  3 network7.WeaponGrenade
	//  4 network7.WeaponLase
	//  5 network7.WeaponNinja
	Weapon network7.Weapon

	// For CTF, if the guy is carrying a flag for example.
	// Only when the sv_gametype is ctf this mode is non zero.
	// It is set in ctf.cpp when a flag is involved on death.
	ModeSpecial int
}

func (msg *SvKillMsg) MsgId() int {
	return network7.MsgGameSvKillMsg
}

func (msg *SvKillMsg) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvKillMsg) System() bool {
	return false
}

func (msg *SvKillMsg) Vital() bool {
	return true
}

func (msg *SvKillMsg) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.KillerId),
		packer.PackInt(msg.VictimId),
		packer.PackInt(int(msg.Weapon)),
		packer.PackInt(msg.ModeSpecial),
	)
}

func (msg *SvKillMsg) Unpack(u *packer.Unpacker) error {
	msg.KillerId = u.GetInt()
	msg.VictimId = u.GetInt()
	msg.Weapon = network7.Weapon(u.GetInt())
	msg.ModeSpecial = u.GetInt()

	return nil
}

func (msg *SvKillMsg) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvKillMsg) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Character struct {
	ItemId int

	Tick            int
	X               int
	Y               int
	VelX            int
	VelY            int
	Angle           int
	Direction       int
	Jumped          int
	HookedPlayer    int
	HookState       int
	HookTick        int
	HookX           int
	HookY           int
	HookDx          int
	HookDy          int
	Health          int
	Armor           int
	AmmoCount       int
	Weapon          network7.Weapon
	Emote           network7.EyeEmote
	AttackTick      int
	TriggeredEvents int
}

func (o *Character) Id() int {
	return o.ItemId
}

func (o *Character) TypeId() int {
	return network7.ObjCharacter
}

func (o *Character) Size() int {
	return reflect.TypeOf(Character{}).NumField() - 1
}

func (o *Character) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.Tick),
		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(o.VelX),
		packer.PackInt(o.VelY),
		packer.PackInt(o.Angle),
		packer.PackInt(o.Direction),
		packer.PackInt(o.Jumped),
		packer.PackInt(o.HookedPlayer),
		packer.PackInt(o.HookState),
		packer.PackInt(o.HookTick),
		packer.PackInt(o.HookX),
		packer.PackInt(o.HookY),
		packer.PackInt(o.HookDx),
		packer.PackInt(o.HookDy),
		packer.PackInt(o.Health),
		packer.PackInt(o.Armor),
		packer.PackInt(o.AmmoCount),
		packer.PackInt(int(o.Weapon)),
		packer.PackInt(int(o.Emote)),
		packer.PackInt(o.AttackTick),
		packer.PackInt(o.TriggeredEvents),
	)
}

func (o *Character) Unpack(u *packer.Unpacker) error {
	o.Tick = u.GetInt()
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.VelX = u.GetInt()
	o.VelY = u.GetInt()
	o.Angle = u.GetInt()
	o.Direction = u.GetInt()
	o.Jumped = u.GetInt()
	o.HookedPlayer = u.GetInt()
	o.HookState = u.GetInt()
	o.HookTick = u.GetInt()
	o.HookX = u.GetInt()
	o.HookY = u.GetInt()
	o.HookDx = u.GetInt()
	o.HookDy = u.GetInt()
	o.Health = u.GetInt()
	o.Armor = u.GetInt()
	o.AmmoCount = u.GetInt()
	o.Weapon = network7.Weapon(u.GetInt())
	o.Emote = network7.EyeEmote(u.GetInt())
	o.AttackTick = u.GetInt()
	o.TriggeredEvents = u.GetInt()

	return nil
}

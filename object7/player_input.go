package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this object is never included in the snap
// the same order of ints is sent in the system message input
//
// but technically this item is unused
type PlayerInput struct {
	ItemId int

	Direction    int
	TargetX      int
	TargetY      int
	Jump         int
	Fire         int
	Hook         int
	PlayerFlags  int
	WantedWeapon int
	NextWeapon   int
	PrevWeapon   int
}

func (o *PlayerInput) Id() int {
	return o.ItemId
}

func (o *PlayerInput) TypeId() int {
	return network7.ObjPlayerInput
}

func (o *PlayerInput) Size() int {
	return reflect.TypeOf(PlayerInput{}).NumField() - 1
}

func (o *PlayerInput) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.Direction),
		packer.PackInt(o.TargetX),
		packer.PackInt(o.TargetY),
		packer.PackInt(o.Jump),
		packer.PackInt(o.Fire),
		packer.PackInt(o.Hook),
		packer.PackInt(o.PlayerFlags),
		packer.PackInt(o.WantedWeapon),
		packer.PackInt(o.NextWeapon),
		packer.PackInt(o.PrevWeapon),
	)
}

func (o *PlayerInput) Unpack(u *packer.Unpacker) error {
	o.Direction = u.GetInt()
	o.TargetX = u.GetInt()
	o.TargetY = u.GetInt()
	o.Jump = u.GetInt()
	o.Fire = u.GetInt()
	o.Hook = u.GetInt()
	o.PlayerFlags = u.GetInt()
	o.WantedWeapon = u.GetInt()
	o.NextWeapon = u.GetInt()
	o.PrevWeapon = u.GetInt()

	return nil
}

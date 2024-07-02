package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type Flag struct {
	ItemId int

	X    int
	Y    int
	Team network7.GameTeam
}

func (o *Flag) Id() int {
	return o.ItemId
}

func (o *Flag) TypeId() int {
	return network7.ObjFlag
}

func (o *Flag) Size() int {
	return reflect.TypeOf(Flag{}).NumField() - 1
}

func (o *Flag) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(int(o.Team)),
	)
}

func (o *Flag) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.Team = network7.GameTeam(u.GetInt())

	return nil
}

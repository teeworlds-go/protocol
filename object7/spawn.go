package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type Spawn struct {
	ItemId int

	X int
	Y int
}

func (o *Spawn) Id() int {
	return o.ItemId
}

func (o *Spawn) TypeId() int {
	return network7.ObjSpawn
}

func (o *Spawn) Size() int {
	return reflect.TypeOf(Spawn{}).NumField() - 1
}

func (o *Spawn) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
	)
}

func (o *Spawn) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()

	return nil
}

package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type Pickup struct {
	ItemId int

	X    int
	Y    int
	Type network7.Pickup
}

func (o *Pickup) Id() int {
	return o.ItemId
}

func (o *Pickup) TypeId() int {
	return network7.ObjPickup
}

func (o *Pickup) Size() int {
	return reflect.TypeOf(Pickup{}).NumField() - 1
}

func (o *Pickup) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(int(o.Type)),
	)
}

func (o *Pickup) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.Type = network7.Pickup(u.GetInt())

	return nil
}

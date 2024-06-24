package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Pickup struct {
	ItemId int

	X          int
	Y          int
	PickupType network7.Pickup
}

func (o *Pickup) Id() int {
	return o.ItemId
}

func (o *Pickup) Type() int {
	return network7.ObjPickup
}

func (o *Pickup) Size() int {
	return reflect.TypeOf(Pickup{}).NumField() - 1
}

func (o *Pickup) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.Type()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(int(o.PickupType)),
	)
}

func (o *Pickup) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.PickupType = network7.Pickup(u.GetInt())

	return nil
}

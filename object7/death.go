package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type Death struct {
	ItemId int

	X        int
	Y        int
	ClientId int
}

func (o *Death) Id() int {
	return o.ItemId
}

func (o *Death) TypeId() int {
	return network7.ObjDeath
}

func (o *Death) Size() int {
	return reflect.TypeOf(Death{}).NumField() - 1
}

func (o *Death) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(o.ClientId),
	)
}

func (o *Death) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.ClientId = u.GetInt()

	return nil
}

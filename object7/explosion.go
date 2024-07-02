package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type Explosion struct {
	ItemId int

	X int
	Y int
}

func (o *Explosion) Id() int {
	return o.ItemId
}

func (o *Explosion) TypeId() int {
	return network7.ObjExplosion
}

func (o *Explosion) Size() int {
	return reflect.TypeOf(Explosion{}).NumField() - 1
}

func (o *Explosion) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
	)
}

func (o *Explosion) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()

	return nil
}

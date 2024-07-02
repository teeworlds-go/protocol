package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type HammerHit struct {
	ItemId int

	X int
	Y int
}

func (o *HammerHit) Id() int {
	return o.ItemId
}

func (o *HammerHit) TypeId() int {
	return network7.ObjHammerHit
}

func (o *HammerHit) Size() int {
	return reflect.TypeOf(HammerHit{}).NumField() - 1
}

func (o *HammerHit) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
	)
}

func (o *HammerHit) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()

	return nil
}

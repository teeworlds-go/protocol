package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Laser struct {
	ItemId int

	X         int
	Y         int
	FromX     int
	FromY     int
	StartTick int
}

func (o *Laser) Id() int {
	return o.ItemId
}

func (o *Laser) TypeId() int {
	return network7.ObjLaser
}

func (o *Laser) Size() int {
	return reflect.TypeOf(Laser{}).NumField() - 1
}

func (o *Laser) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(o.FromX),
		packer.PackInt(o.FromY),
		packer.PackInt(o.StartTick),
	)
}

func (o *Laser) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.FromX = u.GetInt()
	o.FromY = u.GetInt()
	o.StartTick = u.GetInt()

	return nil
}

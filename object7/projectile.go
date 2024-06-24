package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Projectile struct {
	ItemId int

	X         int
	Y         int
	VelX      int
	VelY      int
	Type      int
	StartTick int
}

func (o *Projectile) Id() int {
	return o.ItemId
}

func (o *Projectile) TypeId() int {
	return network7.ObjProjectile
}

func (o *Projectile) Size() int {
	return reflect.TypeOf(Projectile{}).NumField() - 1
}

func (o *Projectile) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(o.VelX),
		packer.PackInt(o.VelY),
		packer.PackInt(o.Type),
		packer.PackInt(o.StartTick),
	)
}

func (o *Projectile) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.VelX = u.GetInt()
	o.VelY = u.GetInt()
	o.Type = u.GetInt()
	o.StartTick = u.GetInt()

	return nil
}

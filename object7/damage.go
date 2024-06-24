package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// damage indicicator
// displayed as yellow stars around the tee receiving damage
type Damage struct {
	ItemId int

	X int
	Y int
	// affected player receiving damage
	ClientId     int
	Angle        int
	HealthAmount int
	ArmorAmount  int
	// true if the damage receiver the damage dealer
	Self bool
}

func (o *Damage) Id() int {
	return o.ItemId
}

func (o *Damage) TypeId() int {
	return network7.ObjDamage
}

func (o *Damage) Size() int {
	return reflect.TypeOf(Damage{}).NumField() - 1
}

func (o *Damage) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(o.ClientId),
		packer.PackInt(o.Angle),
		packer.PackInt(o.HealthAmount),
		packer.PackInt(o.ArmorAmount),
		packer.PackBool(o.Self),
	)
}

func (o *Damage) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.ClientId = u.GetInt()
	o.Angle = u.GetInt()
	o.HealthAmount = u.GetInt()
	o.ArmorAmount = u.GetInt()
	o.Self = u.GetInt() != 0

	return nil
}

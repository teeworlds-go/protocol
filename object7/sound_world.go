package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SoundWorld struct {
	ItemId int

	X       int
	Y       int
	SoundId int
}

func (o *SoundWorld) Id() int {
	return o.ItemId
}

func (o *SoundWorld) TypeId() int {
	return network7.ObjSoundWorld
}

func (o *SoundWorld) Size() int {
	return reflect.TypeOf(SoundWorld{}).NumField() - 1
}

func (o *SoundWorld) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.X),
		packer.PackInt(o.Y),
		packer.PackInt(o.SoundId),
	)
}

func (o *SoundWorld) Unpack(u *packer.Unpacker) error {
	o.X = u.GetInt()
	o.Y = u.GetInt()
	o.SoundId = u.GetInt()

	return nil
}

package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SpectatorInfo struct {
	ItemId int

	SpecMode    network7.Spec
	SpectatorID int
	X           int
	Y           int
}

func (o *SpectatorInfo) Id() int {
	return o.ItemId
}

func (o *SpectatorInfo) TypeId() int {
	return network7.ObjSpectatorInfo
}

func (o *SpectatorInfo) Size() int {
	return reflect.TypeOf(SpectatorInfo{}).NumField() - 1
}

func (o *SpectatorInfo) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(int(o.SpecMode)),
		packer.PackInt(o.SpectatorID),
		packer.PackInt(o.X),
		packer.PackInt(o.Y),
	)
}

func (o *SpectatorInfo) Unpack(u *packer.Unpacker) error {
	o.SpecMode = network7.Spec(u.GetInt())
	o.SpectatorID = u.GetInt()
	o.X = u.GetInt()
	o.Y = u.GetInt()

	return nil
}

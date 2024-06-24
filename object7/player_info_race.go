package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this is a new snap item that was added after the 0.7 release
// so for backwards compability it includes a size field
// and older clients ignore it
//
// this message is not used by official servers
// and is part of an effort to support community made race modifications
type PlayerInfoRace struct {
	ItemId int

	RaceStartTick int
}

func (o *PlayerInfoRace) Id() int {
	return o.ItemId
}

func (o *PlayerInfoRace) TypeId() int {
	return network7.ObjPlayerInfoRace
}

func (o *PlayerInfoRace) Size() int {
	// TODO: is this correct? is this just payload size or does it contain the size field as well?
	return reflect.TypeOf(PlayerInfoRace{}).NumField() - 1
}

func (o *PlayerInfoRace) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),
		packer.PackInt(o.Size()),

		packer.PackInt(o.RaceStartTick),
	)
}

func (o *PlayerInfoRace) Unpack(u *packer.Unpacker) error {
	o.RaceStartTick = u.GetInt()

	return nil
}

package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
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
	// this is correct and verified
	// the additional size field is not included in the size
	//
	// player info race has an additional size field which other snap items
	// do not have
	// but its size is only 1 (4 bytes) not 2 (8 bytes) if we were to count the size field
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

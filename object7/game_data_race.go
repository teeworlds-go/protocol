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
type GameDataRace struct {
	ItemId int

	BestTime  int
	Precision int
	RaceFlags int
}

func (o *GameDataRace) Id() int {
	return o.ItemId
}

func (o *GameDataRace) TypeId() int {
	return network7.ObjGameDataRace
}

func (o *GameDataRace) Size() int {
	// this is correct and verified
	// the additional size field is not included in the size
	//
	// game data race has an additional size field which other snap items
	// do not have
	// but its size is only 3 (12 bytes) not 4 (16 bytes) if we were to count the size field
	return reflect.TypeOf(GameDataRace{}).NumField() - 1
}

func (o *GameDataRace) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),
		packer.PackInt(o.Size()),

		packer.PackInt(o.BestTime),
		packer.PackInt(o.Precision),
		packer.PackInt(o.RaceFlags),
	)
}

func (o *GameDataRace) Unpack(u *packer.Unpacker) error {
	o.BestTime = u.GetInt()
	o.Precision = u.GetInt()
	o.RaceFlags = u.GetInt()

	return nil
}

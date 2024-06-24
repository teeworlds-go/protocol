package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type GameDataFlag struct {
	ItemId int

	// It is either the client id of the carrier so 0-64 or one of those values
	//
	// -3 - FLAG_MISSING
	// -2 - FLAG_ATSTAND
	// -1 - FLAG_TAKEN
	FlagCarrierRed int

	// It is either the client id of the carrier so 0-64 or one of those values
	//
	// -3 - FLAG_MISSING
	// -2 - FLAG_ATSTAND
	// -1 - FLAG_TAKEN
	FlagCarrierBlue int

	FlagDropTickRed  int
	FlagDropTickBlue int
}

func (o *GameDataFlag) Id() int {
	return o.ItemId
}

func (o *GameDataFlag) Type() int {
	return network7.ObjGameDataFlag
}

func (o *GameDataFlag) Size() int {
	return reflect.TypeOf(GameDataFlag{}).NumField() - 1
}

func (o *GameDataFlag) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.Type()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.FlagCarrierRed),
		packer.PackInt(o.FlagCarrierBlue),
		packer.PackInt(o.FlagDropTickRed),
		packer.PackInt(o.FlagDropTickBlue),
	)
}

func (o *GameDataFlag) Unpack(u *packer.Unpacker) error {
	o.FlagCarrierRed = u.GetInt()
	o.FlagCarrierBlue = u.GetInt()
	o.FlagDropTickRed = u.GetInt()
	o.FlagDropTickBlue = u.GetInt()

	return nil
}

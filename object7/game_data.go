package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type GameData struct {
	ItemId int

	GameStartTick int
	// TODO: add struct to wrap these flags
	//       use network7.GameStateFlag bit operations to turn it into a bunch of booleans

	// GameStateFlags GameStateFlagsStruct
	FlagsRaw int

	GameStateEndTick int
}

func (o *GameData) Id() int {
	return o.ItemId
}

func (o *GameData) TypeId() int {
	return network7.ObjGameData
}

func (o *GameData) Size() int {
	return reflect.TypeOf(GameData{}).NumField() - 1
}

func (o *GameData) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.GameStartTick),
		packer.PackInt(o.FlagsRaw),
		packer.PackInt(o.GameStateEndTick),
	)
}

func (o *GameData) Unpack(u *packer.Unpacker) error {
	o.GameStartTick = u.GetInt()
	o.FlagsRaw = u.GetInt()
	o.GameStateEndTick = u.GetInt()

	return nil
}

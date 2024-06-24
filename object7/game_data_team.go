package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type GameDataTeam struct {
	ItemId int

	TeamscoreRed  int
	TeamscoreBlue int
}

func (o *GameDataTeam) Id() int {
	return o.ItemId
}

func (o *GameDataTeam) Type() int {
	return network7.ObjGameDataTeam
}

func (o *GameDataTeam) Size() int {
	return reflect.TypeOf(GameDataTeam{}).NumField() - 1
}

func (o *GameDataTeam) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.Type()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.TeamscoreRed),
		packer.PackInt(o.TeamscoreBlue),
	)
}

func (o *GameDataTeam) Unpack(u *packer.Unpacker) error {
	o.TeamscoreRed = u.GetInt()
	o.TeamscoreBlue = u.GetInt()

	return nil
}

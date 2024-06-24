package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type PlayerInfo struct {
	ItemId int

	// TODO: parse flags
	PlayerFlags int
	Score       int
	Latency     int
}

func (o *PlayerInfo) Id() int {
	return o.ItemId
}

func (o *PlayerInfo) Type() int {
	return network7.ObjPlayerInfo
}

func (o *PlayerInfo) Size() int {
	return reflect.TypeOf(PlayerInfo{}).NumField() - 1
}

func (o *PlayerInfo) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.Type()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.PlayerFlags),
		packer.PackInt(o.Score),
		packer.PackInt(o.Latency),
	)
}

func (o *PlayerInfo) Unpack(u *packer.Unpacker) error {
	o.PlayerFlags = u.GetInt()
	o.Score = u.GetInt()
	o.Latency = u.GetInt()

	return nil
}

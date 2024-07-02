package object7

import (
	"reflect"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

// only used for demos
// not sent over the network
type DeClientInfo struct {
	ItemId int

	Local           bool
	Team            network7.GameTeam
	Name            [4]int
	Clan            [4]int
	Country         int
	SkinPartNames   [6][6]int
	UseCustomColors [6]int
	SkinPartColors  [6]int
}

func (o *DeClientInfo) Id() int {
	return o.ItemId
}

func (o *DeClientInfo) TypeId() int {
	return network7.ObjDeClientInfo
}

func (o *DeClientInfo) Size() int {
	return reflect.TypeOf(DeClientInfo{}).NumField() - 1
}

func (o *DeClientInfo) Pack() []byte {
	panic("not implemented")
}

func (o *DeClientInfo) Unpack(u *packer.Unpacker) error {
	panic("not implemented")
}

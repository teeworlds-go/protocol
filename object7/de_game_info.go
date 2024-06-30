package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// only used for demos
// never send over the network
type DeGameInfo struct {
	ItemId int

	GameFlags    int
	ScoreLimit   int
	TimeLimit    int
	MatchNum     int
	MatchCurrent int
}

func (o *DeGameInfo) Id() int {
	return o.ItemId
}

func (o *DeGameInfo) TypeId() int {
	return network7.ObjDeGameInfo
}

func (o *DeGameInfo) Size() int {
	return reflect.TypeOf(DeGameInfo{}).NumField() - 1
}

func (o *DeGameInfo) Pack() []byte {
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		packer.PackInt(o.GameFlags),
		packer.PackInt(o.ScoreLimit),
		packer.PackInt(o.TimeLimit),
		packer.PackInt(o.MatchNum),
		packer.PackInt(o.MatchCurrent),
	)
}

func (o *DeGameInfo) Unpack(u *packer.Unpacker) error {
	o.GameFlags = u.GetInt()
	o.ScoreLimit = u.GetInt()
	o.TimeLimit = u.GetInt()
	o.MatchNum = u.GetInt()
	o.MatchCurrent = u.GetInt()

	return nil
}

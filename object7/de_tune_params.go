package object7

import (
	"reflect"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// only used for demos
// never send over the network
type DeTuneParams struct {
	ItemId int

	Params [32]int
}

func (o *DeTuneParams) Id() int {
	return o.ItemId
}

func (o *DeTuneParams) TypeId() int {
	return network7.ObjDeTuneParams
}

func (o *DeTuneParams) Size() int {
	return reflect.TypeOf(DeTuneParams{}).NumField() - 1
}

func (o *DeTuneParams) Pack() []byte {
	payload := []byte{}
	for _, p := range o.Params {
		payload = append(payload, packer.PackInt(p)...)
	}
	return slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),

		payload,
	)
}

func (o *DeTuneParams) Unpack(u *packer.Unpacker) error {
	for i := 0; i < len(o.Params); i++ {
		o.Params[i] = u.GetInt()
	}

	return nil
}

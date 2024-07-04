package object7

import (
	"log/slog"
	"slices"

	"github.com/teeworlds-go/protocol/packer"
)

// used for protocol forward compability
// an imaginary protocol version 0.7.6 could add a new snap item
// it would have an unknown type id followed by a size field
// we just consume size amount of integers and store it in this unknown struct
// then we move onto the next item
type Unknown struct {
	ItemId   int
	ItemType int
	ItemSize int

	Fields []int
}

func (o *Unknown) Id() int {
	return o.ItemId
}

func (o *Unknown) TypeId() int {
	return o.ItemType
}

func (o *Unknown) Size() int {
	return o.ItemSize
}

func (o *Unknown) Pack() []byte {
	data := slices.Concat(
		packer.PackInt(o.TypeId()),
		packer.PackInt(o.Id()),
		packer.PackInt(o.Size()),
	)
	for _, f := range o.Fields {
		data = append(data, packer.PackInt(f)...)
	}
	return data
}

func (o *Unknown) Unpack(u *packer.Unpacker) error {
	o.Fields = make([]int, o.Size())

	for i := 0; i < o.Size(); i++ {
		slog.Debug("unknown unpack", "type", o.TypeId(), "i", i, "size", o.Size())
		o.Fields[i] = u.GetInt()
	}

	return nil
}

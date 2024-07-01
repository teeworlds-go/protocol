package snapshot7

import (
	"fmt"
	"log/slog"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/object7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

const (
	MaxType  = 0x7fff
	MaxId    = 0xffff
	MaxParts = 64
	MaxSize  = MaxParts * 1024
)

type Snapshot struct {
	NumRemovedItems int
	NumItemDeltas   int
	Crc             int

	Items []object7.SnapObject
}

// TODO: this is wasting clock cycles for no reason
//
//	the crc is all snap item payload integers summed together
//	it does not have to be perfectly optimized
//	but repacking every item to get its payload summed is horrible
//
//	i also had another approach with reflect where every snap object would implement
//	Crc() on them selfs
//	but reflect is messy and especially the enum types got annoying to sum
//	because they require specific casting
func CrcItem(o object7.SnapObject) int {
	u := &packer.Unpacker{}
	u.Reset(o.Pack())
	u.GetInt()
	u.GetInt()

	if o.TypeId() == network7.ObjGameDataRace || o.TypeId() == network7.ObjPlayerInfoRace {
		// the backwards compatibility size
		// is not part of the payload
		// and is not used to compute the crc
		u.GetInt()
	}

	crc := 0
	for i := 0; i < o.Size(); i++ {
		crc += u.GetInt()
	}
	return crc
}

func (snap *Snapshot) Unpack(u *packer.Unpacker) error {
	// TODO: add all the error checking the C++ reference implementation has

	snap.NumRemovedItems = u.GetInt()
	snap.NumItemDeltas = u.GetInt()
	u.GetInt() // _zero

	// TODO: copy non deleted items from a delta snapshot

	for i := 0; i < snap.NumRemovedItems; i++ {
		deleted := u.GetInt()
		fmt.Printf("deleted item key = %d\n", deleted)

		// TODO: don't copy those from the delta snapshot
	}

	for i := 0; i < snap.NumItemDeltas; i++ {
		itemType := u.GetInt()
		itemId := u.GetInt()

		slog.Debug("unpack item snap item ", "num", i, "total", snap.NumItemDeltas, "type", itemType, "id", itemId)

		item := object7.NewObject(itemType, itemId, u)
		err := item.Unpack(u)
		if err != nil {
			return err
		}

		snap.Items = append(snap.Items, item)

		// TODO: update old items
	}

	if u.RemainingSize() > 0 {
		// TODO: this should not panic but return an error
		//       once the returned error actually shows up somewhere and can be checked in the tests
		//       https://github.com/teeworlds-go/go-teeworlds-protocol/issues/6
		panic(fmt.Sprintf("unexpected remaining size %d after snapshot unpack\n", u.RemainingSize()))
	}

	crc := 0
	for _, item := range snap.Items {
		crc += CrcItem(item)
	}
	snap.Crc = crc

	return nil
}

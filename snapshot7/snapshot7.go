package snapshot7

import (
	"fmt"

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

	Items []object7.SnapObject
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

		item := object7.NewObject(itemType, itemId)
		err := item.Unpack(u)
		if err != nil {
			return err
		}

		snap.Items = append(snap.Items, item)

		// TODO: update old items
	}

	if u.RemainingSize() > 0 {
		return fmt.Errorf("unexpected remaining size %d after snapshot unpack\n", u.RemainingSize())
	}

	return nil
}

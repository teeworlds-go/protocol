package snapshot7

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/packer"
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

func ItemKey(o object7.SnapObject) int {
	return (o.TypeId() << 16) | (o.Id() & 0xffff)
}

// TODO: this is horrible
func GetItemPayload(o object7.SnapObject) []int {
	data := o.Pack()

	// the + 3 are type, id and size
	ints := make([]int, o.Size()+3)
	u := &packer.Unpacker{}
	u.Reset(data)
	offset := 3
	for i := 0; i < o.Size()+3; i++ {
		if u.RemainingSize() == 0 {
			// this is some hack to identify
			// items with additional size field
			// if we run out of data
			// we just assume there was no size field
			// and then we only cut off
			// type and id to get the payload
			offset = 2
			break
		}
		ints[i] = u.GetInt()
	}
	return ints[offset:]
}

// TODO: don't undiff items the slowest possible way
//
//	there should be way to do it like the C++ implementation
//	which does no unpacking or repacking before applying the diff
func UndiffItemSlow(oldItem object7.SnapObject, diffItem object7.SnapObject) object7.SnapObject {
	if oldItem == nil {
		return diffItem
	}
	if oldItem.TypeId() != diffItem.TypeId() {
		panic("can not diff items of different type")
	}
	if oldItem.Size() != diffItem.Size() {
		panic("can not diff items of different sizes")
	}

	oldPayload := GetItemPayload(oldItem)
	diffPayload := GetItemPayload(diffItem)
	newPayload := []byte{}

	for i := 0; i < oldItem.Size(); i++ {
		diffApplied := oldPayload[i] + diffPayload[i]
		newPayload = append(newPayload, packer.PackInt(diffApplied)...)
	}

	u := &packer.Unpacker{}
	u.Reset(newPayload)

	oldItem.Unpack(u)

	return oldItem
}

// the key is one integer holding both type and id
func (snap *Snapshot) GetItemAtKey(key int) *object7.SnapObject {
	for _, item := range snap.Items {
		if key == ItemKey(item) {
			return &item
		}
	}
	return nil
}

// the key is one integer holding both type and id
func (snap *Snapshot) GetItemIndex(key int) (int, error) {
	for i, item := range snap.Items {
		if key == ItemKey(item) {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

// from has to be the old snapshot we delta against
// and the unpacker has to point to the payload of the new delta snapshot
// the payload starts with NumRemovedItems
//
// it returns the new full snapshot with the delta applied to the from
//
// See also (Snapshot *)Unpack()
func UnpackDelata(from *Snapshot, u *packer.Unpacker) (*Snapshot, error) {
	// TODO: add all the error checking the C++ reference implementation has

	snap := &Snapshot{}

	snap.NumRemovedItems = u.GetInt()
	snap.NumItemDeltas = u.GetInt()
	u.GetInt() // _zero

	slog.Debug("got new snapshot!", "num_deleted", snap.NumRemovedItems, "num_updates", snap.NumItemDeltas)

	deletedKeys := make([]int, snap.NumRemovedItems)
	for d := 0; d < snap.NumRemovedItems; d++ {
		deletedKeys[d] = u.GetInt()
		slog.Debug("delta unpack del key", "key", deletedKeys[d], "d_index", d, "num_deleted", snap.NumRemovedItems, "remaining_data", u.RemainingData())
	}

	for i := 0; i < len(from.Items); i++ {
		fromItem := from.Items[i]
		keep := true

		for _, deletedKey := range deletedKeys {
			if deletedKey == ItemKey(fromItem) {
				slog.Debug("delta del item", "deleted_key", deletedKey, "item_type", fromItem.TypeId(), "item_id", fromItem.Id())
				keep = false
				break
			}
		}

		if keep {
			snap.Items = append(snap.Items, fromItem)
		}
	}

	for i := 0; i < snap.NumItemDeltas; i++ {
		itemType := u.GetInt()
		itemId := u.GetInt()

		slog.Debug("unpack item snap item ", "num", i, "total", snap.NumItemDeltas, "type", itemType, "id", itemId)

		item := object7.NewObject(itemType, itemId, u)
		err := item.Unpack(u)
		if err != nil {
			return nil, err
		}

		key := (itemType << 16) | (itemId & 0xffff)
		oldItem := from.GetItemAtKey(key)

		if oldItem == nil {
			snap.Items = append(snap.Items, item)
		} else {
			item = UndiffItemSlow(*oldItem, item)
			idx, err := snap.GetItemIndex(key)
			if err != nil {
				// TODO: the error message is bogus and it should also not panic
				panic("tried to update item that was not found")
			}
			snap.Items[idx] = item
		}
	}

	if u.RemainingSize() > 0 {
		// TODO: this should not panic but return an error
		//       once the returned error actually shows up somewhere and can be checked in the tests
		//       https://github.com/teeworlds-go/protocol/issues/6
		panic(fmt.Sprintf("unexpected remaining size %d after snapshot unpack\n", u.RemainingSize()))
	}

	crc := 0
	for _, item := range snap.Items {
		crc += CrcItem(item)
	}
	snap.Crc = crc

	return snap, nil
}

// unpacks the snapshot as is
// the unpacker has to point to the payload of the new delta snapshot
// the payload starts with NumRemovedItems
//
// it does not unpack any delta
// just a raw snapshot parser
// useful for inspecting network traffic
// not useful for gameplay relevant things
// because snap items will be missing since it is not merged into the old delta
//
// See also snapshot7.UnpackDelta()
func (snap *Snapshot) Unpack(u *packer.Unpacker) error {
	// TODO: add all the error checking the C++ reference implementation has

	snap.NumRemovedItems = u.GetInt()
	snap.NumItemDeltas = u.GetInt()
	u.GetInt() // _zero

	// TODO: copy non deleted items from a delta snapshot

	for i := 0; i < snap.NumRemovedItems; i++ {
		deleted := u.GetInt()
		slog.Debug("deleted item", "key", deleted)

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
		//       https://github.com/teeworlds-go/protocol/issues/6
		panic(fmt.Sprintf("unexpected remaining size %d after snapshot unpack\n", u.RemainingSize()))
	}

	crc := 0
	for _, item := range snap.Items {
		crc += CrcItem(item)
	}
	snap.Crc = crc

	return nil
}

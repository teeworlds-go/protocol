package snapshot7

import (
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/teeworlds-go/protocol/object7"
)

const (
	// passed to methods
	EmptySnapTick = -1

	// returned by methods
	UninitializedTick = -1
)

var (
	ErrNoAltSnapInSnapStorage = errors.New("there is no alt snap in the storage")
)

// TODO: do we even need this?
//
//	can we just put the snap as is in the map?
type holder struct {
	snap *Snapshot
	tick int
}

// TODO: do we need this at all?
//
//	in teeworlds this makes sense because its a custom
//	data structure
//	but in golang users could just define their own map
//
// TODO: make this an interface with a default implementation which the user can
// change and just replace with their own implementation if they want
type Storage struct {
	mu sync.RWMutex

	// a backlog of a few snapshots
	// kept to unpack new deltas sent by the server
	holder map[int]*holder

	// the alt snap is the snapshot
	// that should be used for everything gameplay related
	// it is the snapshot of the current predicton tick
	// and invalid items were already filtered out
	// TODO: add prediction ticks and item validation
	altSnap holder

	// oldest tick still in the holder
	// not the oldest tick we ever received
	oldestTick int

	// newest tick in the holder
	newestTick int

	// use to store and concatinate data
	// of multi part snapshots
	// use AddIncomingData() and IncomingData() to read and write
	multiPartIncomingData []byte

	// the tick we are currently collecting parts for
	CurrentRecvTick int

	// received parts for the current tick
	// as a bit field
	// to check if we received all previous parts
	// when we get the last part number
	SnapshotParts int
}

func NewStorage() *Storage {
	return &Storage{
		holder:                make(map[int]*holder),
		oldestTick:            UninitializedTick,
		newestTick:            UninitializedTick,
		multiPartIncomingData: make([]byte, 0, MaxSize),
	}
}

func (s *Storage) NewestTick() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.newestTick
}

func (s *Storage) OldestTick() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.oldestTick
}

func (s *Storage) altSnapshot() (*Snapshot, error) {
	if s.altSnap.snap == nil {
		return nil, ErrNoAltSnapInSnapStorage
	}
	return s.altSnap.snap, nil
}

func (s *Storage) AltSnap() (*Snapshot, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.altSnapshot()
}

func (s *Storage) SetAltSnap(tick int, snap *Snapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.altSnap.snap = snap
	s.altSnap.tick = tick
}

func (s *Storage) FindAltSnapItem(typeId, itemId int) (obj object7.SnapObject, found bool, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	altSnap, err := s.altSnapshot()
	if err != nil {
		return nil, false, err
	}
	key := (int(typeId) << 16) | (itemId & 0xffff)
	item, found := altSnap.GetItemAtKey(key) // TODO: does this need to be concurrency safe or is it a read only object?
	if !found {
		return nil, false, nil
	}
	return *item, true, nil
}

func (s *Storage) AddIncomingData(part int, numParts int, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if part == 0 {
		// reset length if we get a new snapshot
		s.multiPartIncomingData = s.multiPartIncomingData[:0]
	}
	if part != numParts-1 {
		// a snapshot payload can be 900 at most
		// so unless it is the last part it should fill
		// all 900 bytes
		if len(data) != MaxPackSize {
			return fmt.Errorf("incomplete part that is not the last expected part: part=%d num_parts=%d expected_size=900 got_size=%d", part, numParts, len(data))
		}
	}
	if len(s.multiPartIncomingData)+len(data) > MaxSize {
		return fmt.Errorf("reached the maximum amount of snapshot data: %d bytes", MaxSize)
	}

	s.multiPartIncomingData = append(s.multiPartIncomingData, data...)

	return nil
}

func (s *Storage) IncomingData() []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return slices.Clone(s.multiPartIncomingData)
}

func (s *Storage) First() (snap *Snapshot, found bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.oldestTick == UninitializedTick {
		return nil, false
	}
	return s.get(s.oldestTick)
}

func (s *Storage) Last() (snap *Snapshot, found bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.newestTick == UninitializedTick {
		return nil, false
	}
	return s.get(s.newestTick)
}

func (s *Storage) PurgeUntil(tick int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for k := range s.holder {
		if k < tick {
			delete(s.holder, k)
		}
	}

	if s.oldestTick != UninitializedTick {
		_, found := s.holder[s.oldestTick]
		if !found {
			s.oldestTick = s.nextTick(s.oldestTick)
		}
	}
	if s.newestTick != UninitializedTick {
		_, found := s.holder[s.oldestTick]
		if !found {
			s.newestTick = s.nextTick(s.newestTick)
		}
	}
}

// you probably never have to use this method
func (s *Storage) Size(tick int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.holder)
}

// you probably never have to use this method
func (s *Storage) TicksSorted() []int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ticksSorted()
}

// you probably never have to use this method
func (s *Storage) ticksSorted() []int {

	ticks := make([]int, 0, len(s.holder))
	for k := range s.holder {
		ticks = append(ticks, k)
	}
	slices.Sort(ticks)
	return ticks
}

func (s *Storage) NextTick(tick int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.nextTick(tick)
}

// you probably never have to use this method
func (s *Storage) nextTick(tick int) int {
	for _, t := range s.ticksSorted() {
		if t > tick {
			return t
		}
	}
	return UninitializedTick
}

// you probably never have to use this method
func (s *Storage) PreviousTick(tick int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.previousTick(tick)
}

func (s *Storage) previousTick(tick int) int {
	ticks := s.ticksSorted()
	slices.Reverse(ticks)
	for _, t := range ticks {
		if t < tick {
			return t
		}
	}
	return UninitializedTick
}

func (s *Storage) Get(tick int) (snap *Snapshot, found bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.get(tick)
}

func (s *Storage) get(tick int) (snap *Snapshot, found bool) {

	if tick == EmptySnapTick {
		// -1 is the magic value for the empty snapshot
		return &Snapshot{}, true
	}
	if tick < 0 {
		panic(fmt.Sprintf("negative ticks not supported! tried to get tick %d", tick))
	}
	holder, found := s.holder[tick]
	if !found {
		return nil, false
	}
	return holder.snap, true
}

func (s *Storage) Add(tick int, snapshot *Snapshot) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if tick < 0 {
		return fmt.Errorf("negative ticks not supported! tried to add tick %d", tick)
	}
	if s.oldestTick == UninitializedTick || tick < s.oldestTick {
		s.oldestTick = tick
	}
	if s.newestTick == UninitializedTick || tick > s.newestTick {
		s.newestTick = tick
	}
	s.holder[tick] = &holder{
		snap: snapshot,
		tick: tick,
	}
	return nil
}

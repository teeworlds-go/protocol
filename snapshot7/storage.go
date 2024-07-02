package snapshot7

import (
	"errors"
	"fmt"
	"slices"
	"sort"
)

// TODO: do we even need this?
//
//	can we just put the snap as is in the map?
type holder struct {
	snap *Snapshot
}

// TODO: do we need this at all?
//
//	in teeworlds this makes sense because its a custom
//	data structure
//	but in golang users could just define their own map
type Storage struct {
	holder     map[int]*holder
	OldestTick int
	NewestTick int
}

func NewStorage() *Storage {
	s := &Storage{}
	s.holder = make(map[int]*holder)
	s.OldestTick = -1
	s.NewestTick = -1
	return s
}

func (s *Storage) First() (*Snapshot, error) {
	if s.OldestTick == -1 {
		return nil, errors.New("no snapshot in storage yet")
	}
	return s.Get(s.OldestTick)
}

func (s *Storage) Last() (*Snapshot, error) {
	if s.NewestTick == -1 {
		return nil, errors.New("no snapshot in storage yet")
	}
	return s.Get(s.NewestTick)
}

func (s *Storage) PurgeUntil(tick int) {
	// TODO: i dont know golang
	//       how to free map values

	deletedTicks := []int{}
	for k := range s.holder {
		if k < tick {
			deletedTicks = append(deletedTicks, k)
		}
	}

	for _, deleted := range deletedTicks {
		// memory management moment
		// fmt.Printf("deleted index=%d (tick %d)\n", deleted, s.holder[deleted].snap.)
		s.holder[deleted] = nil
	}

	if s.OldestTick != -1 {
		if s.holder[s.OldestTick] == nil {
			s.OldestTick = s.NextTick(s.OldestTick)
		}
	}
	if s.NewestTick != -1 {
		if s.holder[s.NewestTick] == nil {
			s.NewestTick = s.NextTick(s.NewestTick)
		}
	}
}

// you probably never have to use this method
func (s *Storage) Size(tick int) int {
	// TODO: this is probably the slowest possible way
	//       to get the size of a map xd
	return len(s.TicksSorted())
}

// you probably never have to use this method
func (s *Storage) TicksSorted() []int {
	ticks := []int{}
	for k, v := range s.holder {
		// TODO: will nil values even be included in the map?
		if v != nil {
			ticks = append(ticks, k)
		}
	}
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i] < ticks[j]
	})
	return ticks
}

// you probably never have to use this method
func (s *Storage) NextTick(tick int) int {
	for _, t := range s.TicksSorted() {
		if t > tick {
			return t
		}
	}
	return -1
}

// you probably never have to use this method
func (s *Storage) PreviousTick(tick int) int {
	ticks := s.TicksSorted()
	slices.Reverse(ticks)
	for _, t := range ticks {
		if t < tick {
			return t
		}
	}
	return -1
}

func (s *Storage) Get(tick int) (*Snapshot, error) {
	if tick == -1 {
		// -1 is the magic value for the empty snapshot
		return &Snapshot{}, nil
	}
	if tick < 0 {
		return nil, fmt.Errorf("negative ticks not supported! tried to get tick %d", tick)
	}
	holder := s.holder[tick]
	if holder == nil {
		return nil, fmt.Errorf("snapshot for tick %d not found", tick)
	}
	return holder.snap, nil
}

func (s *Storage) Add(tick int, snapshot *Snapshot) error {
	if tick < 0 {
		return fmt.Errorf("negative ticks not supported! tried to add tick %d", tick)
	}
	if s.OldestTick == -1 || tick < s.OldestTick {
		s.OldestTick = tick
	}
	if s.NewestTick == -1 || tick > s.NewestTick {
		s.NewestTick = tick
	}
	s.holder[tick] = &holder{snap: snapshot}
	return nil
}

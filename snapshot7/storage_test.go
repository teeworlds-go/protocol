package snapshot7_test

import (
	"slices"
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/snapshot7"
)

func TestMultiPartStorage(t *testing.T) {
	t.Parallel()

	storage := snapshot7.NewStorage()
	err := storage.AddIncomingData(0, 2, []byte{0xaa, 0xbb})
	require.NotNil(t, err)

	err = storage.AddIncomingData(0, 1, []byte{0xff, 0xdd})
	require.NoError(t, err)
	require.Equal(t, []byte{0xff, 0xdd}, storage.IncomingData())

	zeros := []byte{899: 0}
	err = storage.AddIncomingData(0, 2, zeros)
	require.NoError(t, err)
	require.Equal(t, zeros, storage.IncomingData())
	err = storage.AddIncomingData(1, 2, zeros)
	require.NoError(t, err)
	require.Equal(t, slices.Concat(zeros, zeros), storage.IncomingData())
}

func TestStorage(t *testing.T) {
	t.Parallel()

	storage := snapshot7.NewStorage()
	storage.Add(1, &snapshot7.Snapshot{Crc: 100})
	storage.Add(2, &snapshot7.Snapshot{Crc: 200})
	storage.Add(3, &snapshot7.Snapshot{Crc: 300})

	got, found := storage.Get(1)
	require.True(t, found)
	require.Equal(t, 100, got.Crc)

	got, found = storage.Get(2)
	require.True(t, found)
	require.Equal(t, 200, got.Crc)

	got, found = storage.Get(3)
	require.True(t, found)
	require.Equal(t, 300, got.Crc)

	require.Equal(t, 1, storage.NextTick(-99999))
	require.Equal(t, 1, storage.NextTick(0))
	require.Equal(t, 2, storage.NextTick(1))
	require.Equal(t, 3, storage.NextTick(2))
	require.Equal(t, -1, storage.NextTick(3))
	require.Equal(t, -1, storage.NextTick(9999))

	require.Equal(t, -1, storage.PreviousTick(-9999))
	require.Equal(t, -1, storage.PreviousTick(1))
	require.Equal(t, 1, storage.PreviousTick(2))
	require.Equal(t, 2, storage.PreviousTick(3))
	require.Equal(t, 3, storage.PreviousTick(4))
	require.Equal(t, 3, storage.PreviousTick(5))
	require.Equal(t, 3, storage.PreviousTick(99999))

	_, found = storage.Get(4)
	require.False(t, found)

	first, found := storage.First()
	require.True(t, found)
	require.Equal(t, 100, first.Crc)
	last, found := storage.Last()
	require.True(t, found)
	require.Equal(t, 300, last.Crc)
	require.Equal(t, 1, storage.OldestTick())
	require.Equal(t, 3, storage.NewestTick())

	storage.PurgeUntil(2)

	first, found = storage.First()
	require.True(t, found)
	require.Equal(t, 200, first.Crc)
	last, found = storage.Last()
	require.True(t, found)
	require.Equal(t, 300, last.Crc)
	require.Equal(t, 2, storage.OldestTick())
	require.Equal(t, 3, storage.NewestTick())

	_, found = storage.Get(1)
	require.False(t, found)

	got, found = storage.Get(2)
	require.True(t, found)
	require.Equal(t, 200, got.Crc)
}

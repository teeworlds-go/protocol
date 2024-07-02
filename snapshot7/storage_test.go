package snapshot7_test

import (
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/snapshot7"
)

func TestStorage(t *testing.T) {
	t.Parallel()

	storage := snapshot7.NewStorage()
	storage.Add(1, &snapshot7.Snapshot{Crc: 100})
	storage.Add(2, &snapshot7.Snapshot{Crc: 200})
	storage.Add(3, &snapshot7.Snapshot{Crc: 300})

	got, err := storage.Get(1)
	require.NoError(t, err)
	require.Equal(t, 100, got.Crc)

	got, err = storage.Get(2)
	require.NoError(t, err)
	require.Equal(t, 200, got.Crc)

	got, err = storage.Get(3)
	require.NoError(t, err)
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

	got, err = storage.Get(4)
	require.NotNil(t, err)

	first, err := storage.First()
	require.NoError(t, err)
	require.Equal(t, 100, first.Crc)
	last, err := storage.Last()
	require.NoError(t, err)
	require.Equal(t, 300, last.Crc)
	require.Equal(t, 1, storage.OldestTick)
	require.Equal(t, 3, storage.NewestTick)

	storage.PurgeUntil(2)

	first, err = storage.First()
	require.NoError(t, err)
	require.Equal(t, 200, first.Crc)
	last, err = storage.Last()
	require.NoError(t, err)
	require.Equal(t, 300, last.Crc)
	require.Equal(t, 2, storage.OldestTick)
	require.Equal(t, 3, storage.NewestTick)

	got, err = storage.Get(1)
	require.NotNil(t, err)

	got, err = storage.Get(2)
	require.NoError(t, err)
	require.Equal(t, 200, got.Crc)
}

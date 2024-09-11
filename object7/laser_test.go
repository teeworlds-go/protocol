package object7_test

import (
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/packer"
)

func TestLaserStandalone(t *testing.T) {
	t.Parallel()
	// simple pack
	laser := &object7.Laser{
		ItemId:    1,
		X:         200,
		Y:         301,
		FromX:     20,
		FromY:     40,
		StartTick: 7812,
	}

	{
		// this is not verified against anything
		want := []byte{3, 1, 136, 3, 173, 4, 20, 40, 132, 122}
		got := laser.Pack()

		require.Equal(t, want, got)
	}

	// repack
	u := &packer.Unpacker{}
	u.Reset(laser.Pack())
	typeId := u.GetInt()
	require.Equal(t, network7.ObjLaser, int(typeId))
	itemId := u.GetInt()
	require.Equal(t, 1, itemId)
	laser.Unpack(u)

	require.Equal(t, 200, laser.X)
	require.Equal(t, 301, laser.Y)
	require.Equal(t, 20, laser.FromX)
	require.Equal(t, 40, laser.FromY)
	require.Equal(t, 7812, laser.StartTick)
}

func TestLaserStandaloneAllZeros(t *testing.T) {
	t.Parallel()
	// simple pack
	laser := &object7.Laser{
		ItemId:    0,
		X:         0,
		Y:         0,
		FromX:     0,
		FromY:     0,
		StartTick: 0,
	}

	{
		want := []byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		got := laser.Pack()

		require.Equal(t, want, got)
	}

	// repack
	u := &packer.Unpacker{}
	u.Reset(laser.Pack())
	typeId := u.GetInt()
	require.Equal(t, network7.ObjLaser, int(typeId))
	itemId := u.GetInt()
	require.Equal(t, 0, itemId)
	laser.Unpack(u)

	{
		want := 0
		got := laser.X
		require.Equal(t, want, got)
	}
}

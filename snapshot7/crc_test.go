package snapshot7_test

import (
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

func TestLaserCrc(t *testing.T) {
	t.Parallel()
	laser := &object7.Laser{
		ItemId: 9999,

		X:         2,
		Y:         2,
		FromX:     2,
		FromY:     2,
		StartTick: 2,
	}

	require.Equal(t, 10, snapshot7.CrcItem(laser))
}

func TestFlagCrc(t *testing.T) {
	t.Parallel()
	flag := &object7.Flag{
		ItemId: 9999,

		X:    2,
		Y:    2,
		Team: 2,
	}

	require.Equal(t, 6, snapshot7.CrcItem(flag))
}

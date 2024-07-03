package snapshot7_test

import (
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

func TestUndiffMultiByte(t *testing.T) {
	t.Parallel()

	baseChar := &object7.Character{
		X: 100,
	}
	diffChar := &object7.Character{
		X: 1,
	}

	undiffed := snapshot7.UndiffItemSlow(baseChar, diffChar)
	character, ok := undiffed.(*object7.Character)
	require.Equal(t, true, ok)

	require.Equal(t, 101, character.X)
}

func TestUndiffSingleByte(t *testing.T) {
	t.Parallel()

	baseChar := &object7.Character{
		X: 5,
	}
	diffChar := &object7.Character{
		X: 1,
	}

	undiffed := snapshot7.UndiffItemSlow(baseChar, diffChar)
	character, ok := undiffed.(*object7.Character)
	require.Equal(t, true, ok)

	require.Equal(t, 6, character.X)

	baseChar = &object7.Character{
		X: 10,
	}
	diffChar = &object7.Character{
		X: 1,
	}

	undiffed = snapshot7.UndiffItemSlow(baseChar, diffChar)
	character, ok = undiffed.(*object7.Character)
	require.Equal(t, true, ok)

	require.Equal(t, 11, character.X)
}

func TestUndiffMultiByteWithSize(t *testing.T) {
	t.Parallel()

	baseRace := &object7.GameDataRace{
		BestTime: 10,
	}
	diffRace := &object7.GameDataRace{
		BestTime: -1,
	}

	undiffed := snapshot7.UndiffItemSlow(baseRace, diffRace)
	raceData, ok := undiffed.(*object7.GameDataRace)
	require.Equal(t, true, ok)

	require.Equal(t, 9, raceData.BestTime)
}

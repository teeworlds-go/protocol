package teeworlds7_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/snapshot7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

func TestCreateAlt(t *testing.T) {
	client := teeworlds7.NewClient()
	oldSnap := &snapshot7.Snapshot{
		Items: []object7.SnapObject{
			&object7.Character{
				Tick: 10,
				X:    10,
				VelX: 10,
			},
		},
	}
	newSnap := &snapshot7.Snapshot{
		Items: []object7.SnapObject{
			&object7.Character{
				Tick: 15,
				X:    10,
				VelX: 10,
			},
		},
	}
	altSnap := client.CreateAltSnap(oldSnap, newSnap)

	// expect unchanged position in the original snap
	// it should copy the data so we can use the original
	// for network undiffs
	char, ok := newSnap.Items[0].(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 10, char.X)

	// expect evolved position in the old snap
	char, ok = altSnap.Items[0].(*object7.Character)
	require.Equal(t, true, ok)
	// TODO: this should be more than 10 idk how much actually xd
	require.Equal(t, 110, char.X)
}

func TestDeepCopy(t *testing.T) {
	t.Parallel()
	snap := &snapshot7.Snapshot{}
	snap.Items = append(snap.Items, &object7.Character{ItemId: 0, X: 42})
	require.Equal(t, 1, len(snap.Items))

	// I thought this is shallow
	snapCpy := &snapshot7.Snapshot{}
	// snapCpy.Items = append(snapCpy.Items, snap.Items[0])
	snapCpy.Items = make([]object7.SnapObject, 1)
	snapCpy.Items[0] = snap.Items[0]
	copy(snapCpy.Items, snap.Items)

	fmt.Printf("   snap.items: %x\n", snap.Items)
	fmt.Printf("snapCpy.items: %x\n", snapCpy.Items)

	snap.Items = slices.Delete(snap.Items, 0, 1)
	require.Equal(t, 0, len(snap.Items))

	fmt.Printf("   snap.items: %x\n", snap.Items)
	fmt.Printf("snapCpy.items: %x\n", snapCpy.Items)

	require.Equal(t, 1, len(snapCpy.Items))

	char, ok := snapCpy.Items[0].(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 42, char.X)
}

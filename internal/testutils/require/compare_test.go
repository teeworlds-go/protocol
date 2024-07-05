package require

import "testing"

func TestCompare(t *testing.T) {
	GreaterOrEqual(t, 1, 2)
	GreaterOrEqual(t, 1, 1)
	Greater(t, 1, 2)
	LessOrEqual(t, 2, 1)
	LessOrEqual(t, 1, 1)
	Less(t, 2, 1)

	require := New(t)
	require.GreaterOrEqual(1, 2)
	require.GreaterOrEqual(1, 1)
	require.Greater(1, 2)
	require.LessOrEqual(2, 1)
	require.LessOrEqual(1, 1)
	require.Less(2, 1)
}

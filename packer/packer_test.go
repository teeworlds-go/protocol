package packer

import (
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
)

// pack

func TestPackEmptyString(t *testing.T) {
	t.Parallel()
	want := []byte{0x00}
	got := PackStr("")
	require.Equal(t, want, got)
}

func TestPackSimpleString(t *testing.T) {
	t.Parallel()
	want := []byte{'f', 'o', 'o', 0x00}
	got := PackStr("foo")
	require.Equal(t, want, got)
}

func TestPackSmallPositiveInts(t *testing.T) {
	t.Parallel()
	want := []byte{0x01}
	got := PackInt(1)
	require.Equal(t, want, got)
}

func TestPackMultiBytePositiveInts(t *testing.T) {
	t.Parallel()
	want := []byte{0x3F}
	got := PackInt(63)
	require.Equal(t, want, got)

	want = []byte{0x80, 0x01}
	got = PackInt(64)
	require.Equal(t, want, got)

	want = []byte{0x81, 0x01}
	got = PackInt(65)
	require.Equal(t, want, got)
}

func TestPackSmallNegativeInts(t *testing.T) {
	t.Parallel()
	want := []byte{0x40}
	got := PackInt(-1)
	require.Equal(t, want, got)

	want = []byte{0x41}
	got = PackInt(-2)
	require.Equal(t, want, got)
}

func TestPackMultiByteNegativeInts(t *testing.T) {
	t.Parallel()
	want := []byte{0x7E}
	got := PackInt(-63)
	require.Equal(t, want, got)

	got = PackInt(-64)
	want = []byte{0x7F}
	require.Equal(t, want, got)

	want = []byte{0xC0, 0x01}
	got = PackInt(-65)
	require.Equal(t, want, got)
}

// unpack

func TestUnpackSmallPositiveInts(t *testing.T) {
	t.Parallel()
	want := 1
	got := UnpackInt([]byte{0x01})
	require.Equal(t, want, got)

	want = 2
	got = UnpackInt([]byte{0x02})
	require.Equal(t, want, got)

	want = 3
	got = UnpackInt([]byte{0x03})
	require.Equal(t, want, got)
}

func TestUnpackMultiBytePositiveInts(t *testing.T) {
	t.Parallel()
	want := 63
	got := UnpackInt([]byte{0x3f})
	require.Equal(t, want, got)

	want = 64
	got = UnpackInt([]byte{0x80, 0x01})
	require.Equal(t, want, got)

	want = 65
	got = UnpackInt([]byte{0x81, 0x01})
	require.Equal(t, want, got)
}

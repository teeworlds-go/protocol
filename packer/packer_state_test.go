package packer

import (
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
)

// rest and remaining size

func TestUnpackRest(t *testing.T) {
	t.Parallel()
	u := Unpacker{}
	u.Reset([]byte{0x01, 0xff, 0xaa})

	{
		want := 1
		got := u.GetInt()
		require.Equal(t, want, got)

		want = 2
		got = u.RemainingSize()
		require.Equal(t, want, got)
	}

	{
		want := []byte{0xff, 0xaa}
		got := u.Rest()
		require.Equal(t, want, got)
	}

	{
		want := 0
		got := u.RemainingSize()
		require.Equal(t, want, got)
	}
}

// client info

func TestUnpackClientInfo(t *testing.T) {
	t.Parallel()
	u := Unpacker{}
	u.Reset([]byte{
		0x24, 0x00, 0x01, 0x00, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x00,
		0x00, 0x40, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x77, 0x61, 0x72,
		0x64, 0x00, 0x64, 0x75, 0x6f, 0x64, 0x6f, 0x6e, 0x6e, 0x79, 0x00,
		0x00, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73,
		0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73, 0x74, 0x61,
		0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x80, 0xfc, 0xaf, 0x05, 0xeb, 0x83, 0xd0, 0x0a, 0x80, 0xfe,
		0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07, 0x00,
	})

	{
		// message id
		want := 36
		got := u.GetInt()
		require.Equal(t, want, got)

		// client id
		want = 0
		got = u.GetInt()
		require.Equal(t, want, got)

		u.GetInt() // Local bool
		u.GetInt() // Team int
	}

	{
		// name
		want := "gopher"
		got, _ := u.GetString()
		require.Equal(t, want, got)

		// clan
		want = ""
		got, _ = u.GetString()
		require.Equal(t, want, got)
	}

	{
		// country
		want := -1
		got := u.GetInt()
		require.Equal(t, want, got)
	}

	{
		// body
		want := "greensward"
		got, _ := u.GetString()
		require.Equal(t, want, got)
	}
}

// unpack with state

func TestUnpackSimpleInts(t *testing.T) {
	t.Parallel()
	u := Unpacker{}
	u.Reset([]byte{0x01, 0x02, 0x03, 0x0f})

	want := 1
	got := u.GetInt()
	require.Equal(t, want, got)

	want = 2
	got = u.GetInt()
	require.Equal(t, want, got)

	want = 3
	got = u.GetInt()
	require.Equal(t, want, got)

	want = 15
	got = u.GetInt()
	require.Equal(t, want, got)
}

func TestUnpackString(t *testing.T) {
	u := Unpacker{}
	u.Reset([]byte{'f', 'o', 'o', 0x00})

	want := "foo"
	got, _ := u.GetString()
	require.Equal(t, want, got)
}

func TestUnpackTwoStrings(t *testing.T) {
	t.Parallel()
	u := Unpacker{}
	u.Reset([]byte{'f', 'o', 'o', 0x00, 'b', 'a', 'r', 0x00})

	want := "foo"
	got, _ := u.GetString()
	require.Equal(t, want, got)

	want = "bar"
	got, _ = u.GetString()
	require.Equal(t, want, got)

}

func TestUnpackMixed(t *testing.T) {
	t.Parallel()
	u := Unpacker{}
	u.Reset([]byte{0x0F, 0x0F, 'f', 'o', 'o', 0x00, 'b', 'a', 'r', 0x00, 0x01})

	// ints
	{
		want := 15
		got := u.GetInt()
		require.Equal(t, want, got)

		want = 15
		got = u.GetInt()
		require.Equal(t, want, got)
	}

	// strings
	{
		want := "foo"
		got, _ := u.GetString()
		require.Equal(t, want, got)

		want = "bar"
		got, _ = u.GetString()
		require.Equal(t, want, got)
	}

	// ints
	{
		want := 1
		got := u.GetInt()
		require.Equal(t, want, got)
	}
}

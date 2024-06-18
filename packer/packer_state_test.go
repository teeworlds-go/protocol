package packer

import (
	"reflect"
	"testing"
)

func TestUnpackClientInfo(t *testing.T) {
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
		got := u.GetInt()
		want := 36

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		// client id
		got = u.GetInt()
		want = 0

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		u.GetInt() // Local bool
		u.GetInt() // Team int
	}

	{
		// name
		got := u.GetString()
		want := "gopher"

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		// clan
		got = u.GetString()
		want = ""

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	{
		// country
		got := u.GetInt()
		want := -1

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	{
		// body
		got := u.GetString()
		want := "greensward"

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

// unpack with state

func TestUnpackSimpleInts(t *testing.T) {
	u := Unpacker{}
	u.Reset([]byte{0x01, 0x02, 0x03, 0x0f})

	got := u.GetInt()
	want := 1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = u.GetInt()
	want = 2

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = u.GetInt()
	want = 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = u.GetInt()
	want = 15

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackString(t *testing.T) {
	u := Unpacker{}
	u.Reset([]byte{'f', 'o', 'o', 0x00})

	got := u.GetString()
	want := "foo"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackTwoStrings(t *testing.T) {
	u := Unpacker{}
	u.Reset([]byte{'f', 'o', 'o', 0x00, 'b', 'a', 'r', 0x00})

	got := u.GetString()
	want := "foo"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = u.GetString()
	want = "bar"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackMixed(t *testing.T) {
	u := Unpacker{}
	u.Reset([]byte{0x0F, 0x0F, 'f', 'o', 'o', 0x00, 'b', 'a', 'r', 0x00, 0x01})

	// ints
	{
		got := u.GetInt()
		want := 15

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = u.GetInt()
		want = 15

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// strings
	{
		got := u.GetString()
		want := "foo"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = u.GetString()
		want = "bar"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// ints
	{
		got := u.GetInt()
		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

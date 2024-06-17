package packer

import (
	"reflect"
	"testing"
)

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

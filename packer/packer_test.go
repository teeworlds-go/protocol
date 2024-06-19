package packer

import (
	"reflect"
	"testing"
)

// pack

func TestPackEmptyString(t *testing.T) {
	got := PackStr("")
	want := []byte{0x00}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSimpleString(t *testing.T) {
	got := PackStr("foo")
	want := []byte{'f', 'o', 'o', 0x00}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSmallPositiveInts(t *testing.T) {
	got := PackInt(1)
	want := []byte{0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackMultiBytePositiveInts(t *testing.T) {
	got := PackInt(63)
	want := []byte{0x3F}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(64)
	want = []byte{0x80, 0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(65)
	want = []byte{0x81, 0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSmallNegativeInts(t *testing.T) {
	got := PackInt(-1)
	want := []byte{0x40}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(-2)
	want = []byte{0x41}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackMultiByteNegativeInts(t *testing.T) {
	got := PackInt(-63)
	want := []byte{0x7E}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(-64)
	want = []byte{0x7F}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = PackInt(-65)
	want = []byte{0xC0, 0x01}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// unpack

func TestUnpackSmallPositiveInts(t *testing.T) {
	got := UnpackInt([]byte{0x01})
	want := 1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = UnpackInt([]byte{0x02})
	want = 2

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = UnpackInt([]byte{0x03})
	want = 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackMultiBytePositiveInts(t *testing.T) {
	got := UnpackInt([]byte{0x3f})
	want := 63

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = UnpackInt([]byte{0x80, 0x01})
	want = 64

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = UnpackInt([]byte{0x81, 0x01})
	want = 65

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

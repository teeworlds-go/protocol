package packer

import (
	"reflect"
	"testing"
)

// pack

func TestPackSmallPositiveInts(t *testing.T) {
	got, err := PackInt(1)
	want := []byte{0x01}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackMultiBytePositiveInts(t *testing.T) {
	got, err := PackInt(63)
	want := []byte{0x3F}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = PackInt(64)
	want = []byte{0x80, 0x01}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = PackInt(65)
	want = []byte{0x81, 0x01}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackSmallNegativeInts(t *testing.T) {
	got, err := PackInt(-1)
	want := []byte{0x40}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = PackInt(-2)
	want = []byte{0x41}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackMultiByteNegativeInts(t *testing.T) {
	got, err := PackInt(-63)
	want := []byte{0x7E}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = PackInt(-64)
	want = []byte{0x7F}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got, err = PackInt(-65)
	want = []byte{0xC0, 0x01}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

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

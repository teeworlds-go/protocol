package main

import (
	"reflect"
	"slices"
	"testing"
)

// pack

func TestPackFlagsUnset(t *testing.T) {
	flags := PacketFlags{
		connless:    false,
		compression: false,
		resend:      false,
		control:     false,
	}

	got := flags.pack()
	want := []byte{0b0000}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackFlagsCompressionSet(t *testing.T) {
	flags := PacketFlags{
		connless:    false,
		compression: true,
		resend:      false,
		control:     false,
	}

	got := flags.pack()
	want := []byte{0b0100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackFlagsAllSet(t *testing.T) {
	flags := PacketFlags{
		connless:    true,
		compression: true,
		resend:      true,
		control:     true,
	}

	got := flags.pack()
	want := []byte{0b1111}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// unpack

func TestUnpackFlagsAllSet(t *testing.T) {
	got := PacketFlags{}
	want := PacketFlags{
		connless:    true,
		compression: true,
		resend:      true,
		control:     true,
	}

	got.unpack([]byte{0b00111100})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackFlagsControlSet(t *testing.T) {
	got := PacketFlags{}
	want := PacketFlags{
		connless:    false,
		compression: false,
		resend:      false,
		control:     true,
	}

	got.unpack([]byte{0b00000100})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackFlagsUnset(t *testing.T) {
	got := PacketFlags{}
	want := PacketFlags{
		connless:    false,
		compression: false,
		resend:      false,
		control:     false,
	}

	got.unpack([]byte{0b00000000})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// packet header unpack

func TestUnpackCloseWithReason(t *testing.T) {
	got := PacketHeader{}
	want := PacketHeader{
		flags: PacketFlags{
			connless:    false,
			compression: false,
			resend:      false,
			control:     true,
		},
		ack:       10,
		numChunks: 0,
		token:     [4]byte{0xcf, 0x2e, 0xde, 0x1d},
	}

	got.unpack(slices.Concat([]byte{0x04, 0x0a, 0x00, 0xcf, 0x2e, 0xde, 0x1d, 0x04}, []byte("shutdown"), []byte{0x00}))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackHeaderFlagsControlSet(t *testing.T) {
	got := PacketHeader{}
	want := PacketHeader{
		flags: PacketFlags{
			connless:    false,
			compression: false,
			resend:      false,
			control:     true,
		},
		ack:       0,
		numChunks: 0,
		token:     [4]byte{0xff, 0xff, 0xff, 0xff},
	}

	got.unpack([]byte{0b00000100, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackHeaderFlagsAllSet(t *testing.T) {
	got := PacketHeader{}
	want := PacketHeader{
		flags: PacketFlags{
			connless:    true,
			compression: true,
			resend:      true,
			control:     true,
		},
		ack:       0,
		numChunks: 0,
		token:     [4]byte{0xff, 0xff, 0xff, 0xff},
	}

	got.unpack([]byte{0b00111100, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

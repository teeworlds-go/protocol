package protocol7

import (
	"reflect"
	"slices"
	"testing"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/messages7"
)

// update chunk headers

func TestPackUpdateChunkHeaders(t *testing.T) {
	// The chunk header is nil by default
	packet := Packet{}
	packet.Messages = append(packet.Messages, &messages7.SvChat{Message: "foo"})

	{
		got := packet.Messages[0].Header()

		if got != nil {
			t.Errorf("got %v, wanted %v", got, nil)
		}
	}

	// When packing the chunk header will be set automatically
	// Based on the current context
	conn := &Connection{Sequence: 1}
	packet.Pack(conn)

	{
		got := packet.Messages[0].Header()
		want := &chunk7.ChunkHeader{
			Flags: chunk7.ChunkFlags{
				Vital: true,
			},
			Size: 8,
			Seq:  2,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// When the chunk header is already set
	// Packing will only update the size

	var chat *messages7.SvChat
	var ok bool
	if chat, ok = packet.Messages[0].(*messages7.SvChat); ok {
		chat.Message = "hello world"
		packet.Messages[0] = chat
	} else {
		t.Fatal("failed to cast chat message")
	}

	packet.Pack(conn)

	{
		got := packet.Messages[0].Header()
		want := &chunk7.ChunkHeader{
			Flags: chunk7.ChunkFlags{
				Vital: true,
			},
			Size: 16,
			Seq:  2,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

// pack header

func TestPackHeader(t *testing.T) {
	header := PacketHeader{
		Flags: PacketFlags{
			Connless:    false,
			Compression: false,
			Resend:      false,
			Control:     true,
		},
		Ack:       10,
		NumChunks: 0,
		Token:     [4]byte{0xcf, 0x2e, 0xde, 0x1d},
	}
	got := header.Pack()
	want := []byte{0x04, 0x0a, 0x00, 0xcf, 0x2e, 0xde, 0x1d}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// pack flags

func TestPackFlagsUnset(t *testing.T) {
	flags := PacketFlags{
		Connless:    false,
		Compression: false,
		Resend:      false,
		Control:     false,
	}

	got := flags.Pack()
	want := []byte{0b0000}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackFlagsCompressionSet(t *testing.T) {
	flags := PacketFlags{
		Connless:    false,
		Compression: true,
		Resend:      false,
		Control:     false,
	}

	got := flags.Pack()
	want := []byte{0b0100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPackFlagsAllSet(t *testing.T) {
	flags := PacketFlags{
		Connless:    true,
		Compression: true,
		Resend:      true,
		Control:     true,
	}

	got := flags.Pack()
	want := []byte{0b1111}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// unpack

func TestUnpackFlagsAllSet(t *testing.T) {
	got := PacketFlags{}
	want := PacketFlags{
		Connless:    true,
		Compression: true,
		Resend:      true,
		Control:     true,
	}

	got.Unpack([]byte{0b00111100})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackFlagsControlSet(t *testing.T) {
	got := PacketFlags{}
	want := PacketFlags{
		Connless:    false,
		Compression: false,
		Resend:      false,
		Control:     true,
	}

	got.Unpack([]byte{0b00000100})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackFlagsUnset(t *testing.T) {
	got := PacketFlags{}
	want := PacketFlags{
		Connless:    false,
		Compression: false,
		Resend:      false,
		Control:     false,
	}

	got.Unpack([]byte{0b00000000})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// packet header unpack

func TestUnpackCloseWithReason(t *testing.T) {
	got := PacketHeader{}
	want := PacketHeader{
		Flags: PacketFlags{
			Connless:    false,
			Compression: false,
			Resend:      false,
			Control:     true,
		},
		Ack:       10,
		NumChunks: 0,
		Token:     [4]byte{0xcf, 0x2e, 0xde, 0x1d},
	}

	got.Unpack(slices.Concat([]byte{0x04, 0x0a, 0x00, 0xcf, 0x2e, 0xde, 0x1d, 0x04}, []byte("shutdown"), []byte{0x00}))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackHeaderFlagsControlSet(t *testing.T) {
	got := PacketHeader{}
	want := PacketHeader{
		Flags: PacketFlags{
			Connless:    false,
			Compression: false,
			Resend:      false,
			Control:     true,
		},
		Ack:       0,
		NumChunks: 0,
		Token:     [4]byte{0xff, 0xff, 0xff, 0xff},
	}

	got.Unpack([]byte{0b00000100, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnpackHeaderFlagsAllSet(t *testing.T) {
	got := PacketHeader{}
	want := PacketHeader{
		Flags: PacketFlags{
			Connless:    true,
			Compression: true,
			Resend:      true,
			Control:     true,
		},
		Ack:       0,
		NumChunks: 0,
		Token:     [4]byte{0xff, 0xff, 0xff, 0xff},
	}

	got.Unpack([]byte{0b00111100, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

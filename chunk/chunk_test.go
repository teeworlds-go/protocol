package chunk

import (
	"reflect"
	"testing"
)

func TestVitalHeader(t *testing.T) {
	header := ChunkHeader{}
	header.Unpack([]byte{0x40, 0x10, 0x0a})

	want := ChunkHeader {
		Flags: ChunkFlags {
			Vital: true,
			Resend: false,
		},
		Size: 16,
		Seq: 10,
	}

	if !reflect.DeepEqual(header, want) {
		t.Errorf("got %v, wanted %v", header, want)
	}
}


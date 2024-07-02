package chunk7

import (
	"fmt"

	"github.com/teeworlds-go/protocol/packer"
)

const (
	chunkFlagVital  = 1
	chunkFlagResend = 2
)

type ChunkFlags struct {
	Vital  bool
	Resend bool
}

func (flags *ChunkFlags) ToInt() int {
	v := 0
	if flags.Resend {
		v |= chunkFlagResend
	}
	if flags.Vital {
		v |= chunkFlagVital
	}
	return v
}

type ChunkHeader struct {
	Flags ChunkFlags
	Size  int
	// sequence number
	// will be acknowledged in the packet header ack
	Seq int
}

type Chunk struct {
	Header ChunkHeader
	Data   []byte
}

func (header *ChunkHeader) Pack() []byte {
	len := 2
	if header.Flags.Vital {
		len = 3
	}
	data := make([]byte, len)
	data[0] = (byte(header.Flags.ToInt()&0x03) << 6) | ((byte(header.Size) >> 6) & 0x3f)
	data[1] = (byte(header.Size) & 0x3f)
	if header.Flags.Vital {
		data[1] |= (byte(header.Seq) >> 2) & 0xc0
		data[2] = byte(header.Seq) & 0xff
	}
	return data
}

func (header *ChunkHeader) Unpack(u *packer.Unpacker) error {
	data, err := u.GetRaw(2)
	if err != nil {
		return err
	}
	flagBits := (data[0] >> 6) & 0x03
	header.Flags.Vital = (flagBits & chunkFlagVital) != 0
	header.Flags.Resend = (flagBits & chunkFlagResend) != 0
	header.Size = (int(data[0]&0x3F) << 6) | (int(data[1]) & 0x3F)

	if header.Flags.Vital {
		thirdByte, err := u.GetByte()
		if err != nil {
			return err
		}
		header.Seq = int((data[1]&0xC0)<<2) | int(thirdByte)
	}
	return nil
}

// TODO: give this a better name
//
//	but it would be nice to have two unpack methods for all structs
//
//	Unpack(u *packer.Unpacker)
//	UnpackTodoFindAGoodName(data []byte)
//
//	See https://github.com/teeworlds-go/protocol/issues/5
func (header *ChunkHeader) UnpackRaw(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("size=%d not enough data to read chunk header", len(data))
	}

	flagBits := (data[0] >> 6) & 0x03
	header.Flags.Vital = (flagBits & chunkFlagVital) != 0
	header.Flags.Resend = (flagBits & chunkFlagResend) != 0
	header.Size = (int(data[0]&0x3F) << 6) | (int(data[1]) & 0x3F)

	if header.Flags.Vital {
		if len(data) < 3 {
			return fmt.Errorf("size=%d not enough data to read vital chunk header", len(data))
		}
		header.Seq = int((data[1]&0xC0)<<2) | int(data[2])
	}
	return nil
}

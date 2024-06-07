package chunk

const (
	chunkFlagVital = 1
	chunkFlagResend = 2
)

type ChunkFlags struct {
	Vital bool
	Resend bool
}

type ChunkHeader struct {
	Flags ChunkFlags
	Size  int
	// sequence number
	// will be acknowledged in the packet header ack
	Seq   int
}

type Chunk struct {
	Header ChunkHeader
	Data []byte
}

func (header *ChunkHeader) Unpack(data []byte) {
	flagBits := (data[0] >> 6) & 0x03
	header.Flags.Vital = (flagBits & chunkFlagVital) != 0
	header.Flags.Resend = (flagBits & chunkFlagResend) != 0
	header.Size = (int(data[0] & 0x3F) << 6) | (int(data[1]) & 0x3F)

	if header.Flags.Vital {
		header.Seq = int((data[1] & 0xC0) << 2) | int(data[2])
	}
}


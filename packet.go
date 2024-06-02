package main

const (
	packetFlagControl     = 1
	packetFlagResend      = 2
	packetFlagCompression = 4
	packetFlagConnless    = 8
)

type PacketFlags struct {
	connless    bool
	compression bool
	resend      bool
	control     bool
}

type PacketHeader struct {
	flags     PacketFlags
	ack       int
	numChunks int
	token     [4]byte
}

func (header *PacketHeader) unpack(packet []byte) {
	header.flags.unpack(packet)
	header.ack = (int(packet[0]&0x3) << 8) | int(packet[1])
	header.numChunks = int(packet[2])
	copy(header.token[:], packet[3:7])
}

func (flags *PacketFlags) unpack(packetHeaderRaw []byte) {
	flagBits := packetHeaderRaw[0] >> 2
	flags.control = (flagBits & packetFlagControl) != 0
	flags.resend = (flagBits & packetFlagResend) != 0
	flags.compression = (flagBits & packetFlagCompression) != 0
	flags.connless = (flagBits & packetFlagConnless) != 0
}

func (flags *PacketFlags) pack() []byte {
	data := 0

	if flags.control {
		data |= packetFlagControl
	}
	if flags.resend {
		data |= packetFlagResend
	}
	if flags.compression {
		data |= packetFlagCompression
	}
	if flags.connless {
		data |= packetFlagConnless
	}

	return []byte{byte(data)}
}

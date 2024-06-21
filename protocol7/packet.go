package protocol7

import (
	"slices"

	"github.com/teeworlds-go/huffman"
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/messages7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

const (
	packetFlagControl     = 1
	packetFlagResend      = 2
	packetFlagCompression = 4
	packetFlagConnless    = 8
)

type PacketFlags struct {
	Connless    bool
	Compression bool
	Resend      bool
	Control     bool
}

type PacketHeader struct {
	Flags     PacketFlags
	Ack       int
	NumChunks int
	Token     [4]byte

	// connless
	ResponseToken [4]byte
}

type Packet struct {
	Header   PacketHeader
	Messages []messages7.NetMessage
}

func PackChunk(msg messages7.NetMessage, connection *Connection) []byte {
	if msg.MsgType() == network7.TypeControl {
		return msg.Pack()
	}

	msgId := msg.MsgId() << 1
	if msg.System() {
		msgId |= 1
	}

	if msg.Vital() {
		connection.Sequence++
	}

	msgAndSys := packer.PackInt(msgId)
	payload := msg.Pack()

	if msg.Header() == nil {
		header := &chunk7.ChunkHeader{
			Flags: chunk7.ChunkFlags{
				Vital: msg.Vital(),
			},
			Seq: connection.Sequence,
		}
		msg.SetHeader(header)
	}

	msg.Header().Size = len(msgAndSys) + len(payload)

	data := slices.Concat(
		msg.Header().Pack(),
		msgAndSys,
		payload,
	)

	return data
}

func (packet *Packet) Pack(connection *Connection) []byte {
	payload := []byte{}
	control := false

	for _, msg := range packet.Messages {
		payload = append(payload, PackChunk(msg, connection)...)
		if msg.MsgType() == network7.TypeControl {
			control = true
		}
	}

	packet.Header.NumChunks = len(packet.Messages)

	if control {
		packet.Header.Flags.Connless = false
		packet.Header.Flags.Compression = false
		packet.Header.Flags.Resend = false
		packet.Header.Flags.Control = true
	}

	if packet.Header.Flags.Compression {
		// TODO: store huffman object in connection to avoid reallocating memory
		huff := huffman.Huffman{}
		var err error
		payload, err = huff.Compress(payload)
		if err != nil {
			panic(err)
		}
	}

	return slices.Concat(
		packet.Header.Pack(),
		payload,
	)
}

func (header *PacketHeader) Pack() []byte {
	flags := 0
	if header.Flags.Control {
		flags |= packetFlagControl
	}
	if header.Flags.Resend {
		flags |= packetFlagResend
	}
	if header.Flags.Compression {
		flags |= packetFlagCompression
	}
	if header.Flags.Connless {
		flags |= packetFlagConnless
	}

	if header.Flags.Connless {
		version := 1
		return slices.Concat(
			[]byte{byte(((packetFlagConnless << 2) & 0x0fc) | (version & 0x03))},
			header.Token[:],
			header.ResponseToken[:],
		)
	}

	return slices.Concat(
		[]byte{
			byte(((flags << 2) & 0xfc) | ((header.Ack >> 8) & 0x03)),
			byte(header.Ack & 0x0ff),
			byte(header.NumChunks),
		},
		header.Token[:],
	)
}

func (header *PacketHeader) Unpack(packet []byte) {
	header.Flags.Unpack(packet)
	header.Ack = (int(packet[0]&0x3) << 8) | int(packet[1])
	header.NumChunks = int(packet[2])
	copy(header.Token[:], packet[3:7])
}

func (flags *PacketFlags) Unpack(packetHeaderRaw []byte) {
	flagBits := packetHeaderRaw[0] >> 2
	flags.Control = (flagBits & packetFlagControl) != 0
	flags.Resend = (flagBits & packetFlagResend) != 0
	flags.Compression = (flagBits & packetFlagCompression) != 0
	flags.Connless = (flagBits & packetFlagConnless) != 0
}

func (flags *PacketFlags) Pack() []byte {
	data := 0

	if flags.Control {
		data |= packetFlagControl
	}
	if flags.Resend {
		data |= packetFlagResend
	}
	if flags.Compression {
		data |= packetFlagCompression
	}
	if flags.Connless {
		data |= packetFlagConnless
	}

	return []byte{byte(data)}
}

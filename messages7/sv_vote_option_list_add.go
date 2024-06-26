package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// You have to manually set NumOptions to be the amount of Descriptions
// For example:
//
// options := []string{"foo", "bar", "baz"}
// voteAdd := SvVoteOptionListAdd{NumOptions: len(options), Descriptions: options}
type SvVoteOptionListAdd struct {
	ChunkHeader *chunk7.ChunkHeader

	NumOptions   int
	Descriptions []string

	// this is used for compatibility with the
	// ddnet server 0.7 bridge implementation
	// to ensure if a ddnet message is unpacked and packed again
	// it matches it exactly
	//
	// TODO: should this be a boolean instead?
	//       instead of counting on unpack
	//       it could always fill it up to 14 descriptions on pack
	//       if some Compat06 bool is set
	NumUnusedOptions int
}

func (msg *SvVoteOptionListAdd) MsgId() int {
	return network7.MsgGameSvVoteOptionListAdd
}

func (msg *SvVoteOptionListAdd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvVoteOptionListAdd) System() bool {
	return false
}

func (msg *SvVoteOptionListAdd) Vital() bool {
	return true
}

func (msg *SvVoteOptionListAdd) Pack() []byte {
	options := []byte{}
	for _, option := range msg.Descriptions {
		options = append(options, packer.PackStr(option)...)
	}
	if msg.NumUnusedOptions > 0 {
		unused := make([]byte, msg.NumUnusedOptions)
		options = append(options, unused...)
	}

	return slices.Concat(
		packer.PackInt(msg.NumOptions),
		options,
	)
}

func (msg *SvVoteOptionListAdd) Unpack(u *packer.Unpacker) error {
	startRemainingSize := u.RemainingSize()

	msg.NumOptions = u.GetInt()
	msg.Descriptions = make([]string, msg.NumOptions)
	for i := 0; i < msg.NumOptions; i++ {
		msg.Descriptions[i], _ = u.GetString()
	}

	finishRemainingSize := u.RemainingSize()

	if msg.Header() != nil {
		consumedSize := startRemainingSize - finishRemainingSize
		numUnused := (msg.Header().Size - 1) - consumedSize
		if numUnused > 0 {
			msg.NumUnusedOptions = numUnused
		}
	}

	return nil
}

func (msg *SvVoteOptionListAdd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteOptionListAdd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

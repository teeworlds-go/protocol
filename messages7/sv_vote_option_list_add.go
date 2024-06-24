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

	return slices.Concat(
		packer.PackInt(msg.NumOptions),
		options,
	)
}

func (msg *SvVoteOptionListAdd) Unpack(u *packer.Unpacker) error {
	msg.NumOptions = u.GetInt()
	msg.Descriptions = make([]string, msg.NumOptions)
	for i := 0; i < msg.NumOptions; i++ {
		msg.Descriptions[i] = u.GetString()
	}

	return nil
}

func (msg *SvVoteOptionListAdd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteOptionListAdd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

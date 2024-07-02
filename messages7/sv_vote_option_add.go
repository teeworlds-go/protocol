package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvVoteOptionAdd struct {
	ChunkHeader *chunk7.ChunkHeader

	Description string
}

func (msg *SvVoteOptionAdd) MsgId() int {
	return network7.MsgGameSvVoteOptionAdd
}

func (msg *SvVoteOptionAdd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvVoteOptionAdd) System() bool {
	return false
}

func (msg *SvVoteOptionAdd) Vital() bool {
	return true
}

func (msg *SvVoteOptionAdd) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Description),
	)
}

func (msg *SvVoteOptionAdd) Unpack(u *packer.Unpacker) error {
	msg.Description, _ = u.GetString()
	return nil
}

func (msg *SvVoteOptionAdd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteOptionAdd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvVoteOptionRemove struct {
	ChunkHeader *chunk7.ChunkHeader

	Description string
}

func (msg *SvVoteOptionRemove) MsgId() int {
	return network7.MsgGameSvVoteOptionRemove
}

func (msg *SvVoteOptionRemove) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvVoteOptionRemove) System() bool {
	return false
}

func (msg *SvVoteOptionRemove) Vital() bool {
	return true
}

func (msg *SvVoteOptionRemove) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Description),
	)
}

func (msg *SvVoteOptionRemove) Unpack(u *packer.Unpacker) error {
	msg.Description = u.GetString()
	return nil
}

func (msg *SvVoteOptionRemove) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteOptionRemove) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

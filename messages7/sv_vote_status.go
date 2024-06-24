package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvVoteStatus struct {
	ChunkHeader *chunk7.ChunkHeader

	Yes   int
	No    int
	Pass  int
	Total int
}

func (msg *SvVoteStatus) MsgId() int {
	return network7.MsgGameSvVoteStatus
}

func (msg *SvVoteStatus) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvVoteStatus) System() bool {
	return false
}

func (msg *SvVoteStatus) Vital() bool {
	return true
}

func (msg *SvVoteStatus) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.Yes),
		packer.PackInt(msg.No),
		packer.PackInt(msg.Pass),
		packer.PackInt(msg.Total),
	)
}

func (msg *SvVoteStatus) Unpack(u *packer.Unpacker) error {
	msg.Yes = u.GetInt()
	msg.No = u.GetInt()
	msg.Pass = u.GetInt()
	msg.Total = u.GetInt()

	return nil
}

func (msg *SvVoteStatus) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteStatus) StatusHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

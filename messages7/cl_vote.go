package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type ClVote struct {
	ChunkHeader *chunk7.ChunkHeader

	Choice network7.VoteChoice
}

func (msg *ClVote) MsgId() int {
	return network7.MsgGameClVote
}

func (msg *ClVote) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClVote) System() bool {
	return false
}

func (msg *ClVote) Vital() bool {
	return true
}

func (msg *ClVote) Pack() []byte {
	return packer.PackInt(int(msg.Choice))
}

func (msg *ClVote) Unpack(u *packer.Unpacker) error {
	msg.Choice = network7.VoteChoice(u.GetInt())

	return nil
}

func (msg *ClVote) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClVote) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

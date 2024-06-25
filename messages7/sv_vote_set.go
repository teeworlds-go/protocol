package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvVoteSet struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId    int
	Type        network7.Vote
	Timeout     int
	Description string
	Reason      string
}

func (msg *SvVoteSet) MsgId() int {
	return network7.MsgGameSvVoteSet
}

func (msg *SvVoteSet) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvVoteSet) System() bool {
	return false
}

func (msg *SvVoteSet) Vital() bool {
	return true
}

func (msg *SvVoteSet) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.ClientId),
		packer.PackInt(int(msg.Type)),
		packer.PackInt(msg.Timeout),
		packer.PackStr(msg.Description),
		packer.PackStr(msg.Reason),
	)
}

func (msg *SvVoteSet) Unpack(u *packer.Unpacker) error {
	msg.ClientId = u.GetInt()
	msg.Type = network7.Vote(u.GetInt())
	msg.Timeout = u.GetInt()
	msg.Description, _ = u.GetString()
	msg.Reason, _ = u.GetString()
	return nil
}

func (msg *SvVoteSet) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteSet) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

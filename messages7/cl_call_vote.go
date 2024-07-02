package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type ClCallVote struct {
	ChunkHeader *chunk7.ChunkHeader

	Type   string
	Value  string
	Reason string
	Force  bool
}

func (msg *ClCallVote) MsgId() int {
	return network7.MsgGameClCallVote
}

func (msg *ClCallVote) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ClCallVote) System() bool {
	return false
}

func (msg *ClCallVote) Vital() bool {
	return true
}

func (msg *ClCallVote) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Type),
		packer.PackStr(msg.Value),
		packer.PackStr(msg.Reason),
		packer.PackBool(msg.Force),
	)
}

func (msg *ClCallVote) Unpack(u *packer.Unpacker) error {
	var err error
	msg.Type, err = u.GetString()
	if err != nil {
		return err
	}
	msg.Value, err = u.GetString()
	if err != nil {
		return err
	}
	msg.Reason, err = u.GetString()
	if err != nil {
		return err
	}
	msg.Force = u.GetInt() != 0

	return nil
}

func (msg *ClCallVote) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClCallVote) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

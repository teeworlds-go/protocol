package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// This message is used by the server to empty all vote entries in the client menu.
type SvVoteClearOptions struct {
	ChunkHeader *chunk7.ChunkHeader
}

func (msg *SvVoteClearOptions) MsgId() int {
	return network7.MsgGameSvVoteClearOptions
}

func (msg *SvVoteClearOptions) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvVoteClearOptions) System() bool {
	return false
}

func (msg *SvVoteClearOptions) Vital() bool {
	return true
}

func (msg *SvVoteClearOptions) Pack() []byte {
	return []byte{}
}

func (msg *SvVoteClearOptions) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *SvVoteClearOptions) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvVoteClearOptions) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

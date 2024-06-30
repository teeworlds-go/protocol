package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// only used for demos
// never send over the network
type DeClientEnter struct {
	ChunkHeader *chunk7.ChunkHeader

	Name     string
	ClientId int
	Team     network7.GameTeam
}

func (msg *DeClientEnter) MsgId() int {
	return network7.MsgGameDeClientEnter
}

func (msg *DeClientEnter) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *DeClientEnter) System() bool {
	return false
}

func (msg *DeClientEnter) Vital() bool {
	return true
}

func (msg *DeClientEnter) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
		packer.PackInt(msg.ClientId),
		packer.PackInt(int(msg.Team)),
	)
}

func (msg *DeClientEnter) Unpack(u *packer.Unpacker) error {
	var err error
	msg.Name, err = u.GetString()
	if err != nil {
		return err
	}
	msg.ClientId = u.GetInt()
	msg.Team = network7.GameTeam(u.GetInt())

	return nil
}

func (msg *DeClientEnter) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *DeClientEnter) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

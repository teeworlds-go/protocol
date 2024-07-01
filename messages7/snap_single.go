package messages7

// same connection as full_ddnet_server_multi_test.go
// but NETMSG_SNAPSINGLE

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/snapshot7"
)

type SnapSingle struct {
	ChunkHeader *chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
	Crc       int
	PartSize  int

	// TODO: remove data when snapshot packing works
	//       as of right now Data and Snapshot are the same thing
	Data     []byte
	Snapshot snapshot7.Snapshot
}

func (msg *SnapSingle) MsgId() int {
	return network7.MsgSysSnapSingle
}

func (msg *SnapSingle) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SnapSingle) System() bool {
	return true
}

func (msg *SnapSingle) Vital() bool {
	return false
}

func (msg *SnapSingle) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.GameTick),
		packer.PackInt(msg.DeltaTick),
		packer.PackInt(msg.Crc),
		packer.PackInt(msg.PartSize),
		msg.Data,
	)
}

func (msg *SnapSingle) Unpack(u *packer.Unpacker) error {
	msg.GameTick = u.GetInt()
	msg.DeltaTick = u.GetInt()
	msg.Crc = u.GetInt()
	msg.PartSize = u.GetInt()
	msg.Data = u.Rest()

	// TODO: should this be optional?
	//       there is also snapshot7.UnpackDelta
	//       which unpacks the snapshot AND applies it as a delta
	//       to get a full snapshot with ALL items in it
	//       which is what one wants in most cases anyways
	//
	//       but there is still the chance that someone wants to inspect
	//       what is sent over the network as is without any
	//       delta applied
	//
	//       for consistency this should also be added to partial snap messages (NETMSG_SNAP)
	//       but you can not really dump the items of snapshot part 3 out of 5 can you?
	//
	//       currently all the tests depend on it
	//       and it is nice to have unit tests that only look at one packet
	//       without any deltas

	// genius
	u.Reset(msg.Data)
	err := msg.Snapshot.Unpack(u)
	if err != nil {
		return err
	}

	return nil
}

func (msg *SnapSingle) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SnapSingle) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

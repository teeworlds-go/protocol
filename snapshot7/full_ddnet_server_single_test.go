package snapshot7_test

// same connection as full_ddnet_server_multi_test.go
// but NETMSG_SNAPSINGLE

import (
	"fmt"
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/object7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

// --------------------------------
// snap single
// --------------------------------

func TestDDNetFullServerSnapSingle(t *testing.T) {
	// vanilla client connected to almost full ddnet rus server
	// map: Multeasymap
	// dumped with tcpdump
	// libtw2 dissector details
	// this is the n-th snapshot that is a snap single the prior snapshots were partials snaps
	//
	// User Datagram Protocol, Src Port: 8316, Dst Port: 51479
	// Teeworlds 0.7 Protocol packet
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 11271652
	//     Delta tick: 22
	//     Crc: 1032529567
	//     Data (114 bytes)
	dump := []byte{
		0x10, 0x04, 0x01, 0x1f, 0x5a, 0x5e, 0xd8,
		0x28, 0x15, 0xa6, 0x6c, 0xfb, 0x4a, 0xd3, 0xd6, 0xd3, 0xe0, 0x54, 0x85, 0x09, 0x72, 0xc4,
		0x0d, 0x2e, 0xb6, 0xe9, 0x2f, 0xbc, 0xfd, 0xb9, 0x5d, 0xd0, 0x35, 0x5d, 0x5e, 0xcf, 0xb5, 0x35,
		0xeb, 0x92, 0xe8, 0xbe, 0xbb, 0x9f, 0xd1, 0x7f, 0x17, 0x24, 0x2d, 0xb8, 0xfa, 0x7f, 0x5f, 0xf5,
		0x38, 0x7b, 0x7a, 0xc1, 0x47, 0x73, 0xeb, 0x4a, 0xd0, 0x62, 0x02, 0x17, 0x6d, 0xd9, 0xe2, 0x58,
		0x57, 0x68, 0xc9, 0xe6, 0x35, 0xf9, 0x0a, 0xd3, 0xd6, 0x7b, 0xf6, 0x5d, 0x90, 0xb4, 0x10, 0xfa,
		0x9a, 0x7c, 0x85, 0x69, 0xeb, 0xbf, 0xe1, 0x2b, 0x4d, 0x5b, 0xef, 0x1e, 0x4f, 0xea, 0x85, 0xe2,
		0x06,
	}

	conn := protocol7.Session{}

	// unpack
	packet := protocol7.Packet{}
	err := packet.Unpack(dump)
	require.NoError(t, err)

	// repack
	conn.Ack = packet.Header.Ack
	repack := packet.Pack(&conn)
	require.Equal(t, dump, repack)

	// content
	require.Equal(t, 1, len(packet.Messages))

	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[0].MsgId())
	msg, ok := packet.Messages[0].(*messages7.SnapSingle)
	require.Equal(t, true, ok)

	// this is not verified
	require.Equal(t, 6, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)

	// this is not verified
	require.Equal(t, 6, len(msg.Snapshot.Items))

	// this is not verified
	item := msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok := item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 4, playerInfo.PlayerFlags)
	// this is odd the score in ddnet should not be 0
	// if a player has no score shouldnt it be -99999? or did that change?
	require.Equal(t, 0, playerInfo.Score)
	// this is highly suspicious. why is there 0 latency for some player on a ddnet rus server?
	require.Equal(t, 0, playerInfo.Latency)

	// this is not verified
	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 0, playerInfo.PlayerFlags)
	require.Equal(t, 758141, playerInfo.Score)
	// this is highly suspicious. why is there 0 latency for some player on a ddnet rus server?
	require.Equal(t, 0, playerInfo.Latency)

	// this is not verified
	item = msg.Snapshot.Items[2]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok := item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 45, character.Tick)
	// that seems like a odd position to be in on Multeasymap
	// this is probably wrong
	require.Equal(t, 0, character.X)
	require.Equal(t, 32, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, -128, character.VelY)
	require.Equal(t, 995, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, 0, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 0, character.HookX)
	require.Equal(t, 33, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, 0, character.Weapon)
	require.Equal(t, 0, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// this is not verified
	// not fully testing ddnet ex item
	item = msg.Snapshot.Items[3]
	require.Equal(t, 32765, item.TypeId())

	// this is not verified
	item = msg.Snapshot.Items[4]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 23, character.Tick)
	// that seems like a odd position to be in on Multeasymap
	// this is probably wrong
	require.Equal(t, -43, character.X)
	require.Equal(t, 155, character.Y)
	require.Equal(t, 588, character.VelX)
	require.Equal(t, 2944, character.VelY)
	require.Equal(t, -196, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, 0, character.HookedPlayer)
	require.Equal(t, 1, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 91, character.HookX)
	require.Equal(t, 422, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	// this for sure is wrong???
	require.Equal(t, 11271786, character.AmmoCount)
	require.Equal(t, 4, character.Weapon)
	require.Equal(t, 5, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)
}

// --------------------------------
// snap single & sv emoticon
// --------------------------------

func TestSvEmoticonAndSnapSingle(t *testing.T) {
	// Teeworlds 0.7 Protocol packet
	// Teeworlds 0.7 Protocol chunk: game.sv_emoticon
	//     Header (vital: 234)
	//     Message: game.sv_emoticon
	//     Client id: 39
	//     Emoticon: exclamation
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 11271828
	//     Delta tick: 24
	//     Crc: 1021386082
	//     Data (190 bytes)
	dump := []byte{
		0x10, 0x05, 0x02, 0x1f, 0x5a, 0x5e,
		0xd8, 0x4a, 0x16, 0xd4, 0x0d, 0x2f, 0x0b, 0x2d, 0xfc, 0x4d, 0x19, 0xae, 0xd4, 0xb4, 0xf5, 0x10,
		0x9f, 0xa1, 0x48, 0x8b, 0x2b, 0xbb, 0xd1, 0x15, 0x48, 0x5a, 0xce, 0x52, 0xed, 0xef, 0x97, 0x58,
		0xef, 0xa2, 0x29, 0x6d, 0x82, 0xb7, 0x6b, 0xb6, 0x91, 0x1a, 0x1b, 0x4e, 0x72, 0xbb, 0x1d, 0x7d,
		0xcf, 0x71, 0x04, 0x55, 0x09, 0x5f, 0x21, 0xfe, 0x2e, 0x48, 0x5a, 0xb8, 0xa8, 0xff, 0x87, 0x42,
		0x26, 0x55, 0xa2, 0x9e, 0xeb, 0x22, 0x2d, 0x37, 0xc6, 0xb0, 0xa4, 0x33, 0x9e, 0x96, 0x1b, 0x0b,
		0xfc, 0xab, 0xaa, 0x87, 0x37, 0x5c, 0xa9, 0x69, 0xeb, 0x19, 0x92, 0xe8, 0x3b, 0x56, 0x3e, 0x0b,
		0x5b, 0x80, 0x73, 0xea, 0xb4, 0x50, 0xaa, 0x54, 0x6f, 0x49, 0xa2, 0x0b, 0xac, 0x38, 0x15, 0x3c,
		0x66, 0x71, 0xc4, 0x5d, 0x90, 0xb4, 0x80, 0xd7, 0x03, 0x60, 0x53, 0x14, 0xa7, 0x5e, 0xf8, 0x90,
		0xa7, 0x17, 0x3c, 0xa1, 0x17, 0xbe, 0x2e, 0x54, 0xa0, 0xde, 0x0f, 0xe4, 0x2d, 0xe8, 0xcc, 0x44,
		0xc2, 0x2f, 0x85, 0x36, 0x7f, 0x17, 0x24, 0x2d, 0x84, 0xfe, 0x9f, 0x25, 0x52, 0xaf, 0x57, 0xdc,
		0x00,
	}

	conn := protocol7.Session{}

	packet := protocol7.Packet{}
	err := packet.Unpack(dump)
	require.NoError(t, err)

	conn.Ack = packet.Header.Ack
	repack := packet.Pack(&conn)

	fmt.Printf("repack: %x\n", repack)
	fmt.Printf("dump:   %x\n", dump)

	require.Equal(t, dump, repack)

	// content
	require.Equal(t, 2, len(packet.Messages))

	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[1].MsgId())
	msg, ok := packet.Messages[1].(*messages7.SnapSingle)
	require.Equal(t, true, ok)

	require.Equal(t, 11271828, msg.GameTick)
	require.Equal(t, 24, msg.DeltaTick)
	require.Equal(t, 1021386082, msg.Crc)
	// TODO: this should match
	// require.Equal(t, 1021386082, msg.Snapshot.Crc)

	// this is not verified
	require.Equal(t, 8, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 1, msg.Snapshot.NumRemovedItems)

	// this is not verified
	require.Equal(t, 8, len(msg.Snapshot.Items))

	// this is not verified
	item := msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok := item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 2, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
}

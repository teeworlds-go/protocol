package messages7_test

import (
	"fmt"
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

func TestKeepDDNetServerTranslationTrailingNullbytes(t *testing.T) {
	// the sv vote option list add is similiar in 0.6 and 0.7 but not exactly the same
	// https://chillerdragon.github.io/teeworlds-protocol/06/game_messages.html#NETMSGTYPE_SV_VOTEOPTIONLISTADD
	// https://chillerdragon.github.io/teeworlds-protocol/07/game_messages.html#NETMSGTYPE_SV_VOTEOPTIONLISTADD
	//
	// 0.6 always sends a fixed amount of strings with a leading amount field
	// to indicate how many of them should be consumed
	//
	// 0.7 only sends as many strings as there are set in the amount field
	//
	// Example 0.6
	// Int(num)    3
	// Str(desc)  "foo"
	// Str(desc)  "bar"
	// Str(desc)  "baz"
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	// Str(desc)  ""
	//
	// Example 0.7
	// Int(num)    3
	// Str(desc)  "foo"
	// Str(desc)  "bar"
	// Str(desc)  "baz"
	//
	// the ddnet server as of right does not do any kind of translation of this message
	// the message id is also the same
	// so the ddnet server sends the 0.6 message to 0.7 connections
	//
	// https://github.com/ddnet/ddnet/blob/1cf8761b441bf2381a28a476d4ea889450060284/src/game/server/gamecontext.cpp#L1409-L1470
	//

	dumpVoteList := []byte{
		0x00, 0x04, 0x01, 0x1f, 0x5a, 0x5e, 0xd8,
		0x42, 0x05, 0xbf, 0x18, 0x05, 0x33, 0x36, 0x35, 0x39, 0x37, 0x20, 0xe2, 0x9a, 0x91, 0x20, 0x7c,
		0x20, 0x30, 0x35, 0x3a, 0x32, 0x36, 0x20, 0xe2, 0x97, 0xb7, 0x00, 0x47, 0x65, 0x6e, 0x65, 0x71,
		0x75, 0x65, 0x20, 0x62, 0x79, 0x20, 0x54, 0x68, 0x65, 0x6d, 0x69, 0x78, 0x20, 0x7c, 0x20, 0x34,
		0x2f, 0x35, 0x20, 0xe2, 0x98, 0x85, 0x00, 0x35, 0x31, 0x37, 0x38, 0x36, 0x20, 0xe2, 0x9a, 0x91,
		0x20, 0x7c, 0x20, 0x32, 0x36, 0x3a, 0x33, 0x31, 0x20, 0xe2, 0x97, 0xb7, 0x00, 0x47, 0x65, 0x52,
		0x6f, 0x6c, 0x6c, 0x41, 0x20, 0x62, 0x79, 0x20, 0x45, 0x76, 0x6f, 0x6c, 0x69, 0x20, 0x7c, 0x20,
		0x33, 0x2f, 0x35, 0x20, 0xe2, 0x98, 0x85, 0x00, 0x33, 0x30, 0x39, 0x32, 0x31, 0x20, 0xe2, 0x9a,
		0x91, 0x20, 0x7c, 0x20, 0x32, 0x35, 0x3a, 0x30, 0x32, 0x20, 0xe2, 0x97, 0xb7, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	conn := protocol7.Session{}

	packet := protocol7.Packet{}
	err := packet.Unpack(dumpVoteList)
	require.NoError(t, err)

	conn.Ack = packet.Header.Ack
	repack := packet.Pack(&conn)

	fmt.Printf("  want: %x\n", dumpVoteList)
	fmt.Printf("repack: %x\n", repack)

	require.Equal(t, dumpVoteList, repack)
}

func TestVoteListAdd(t *testing.T) {
	// unpack
	fullChunk := []byte{0x40, 0x06, 0x06, 0x18, 0x01, 0x66, 0x6f, 0x6f, 0x00}
	u := &packer.Unpacker{}
	u.Reset(fullChunk)

	header := &chunk7.ChunkHeader{}
	header.Unpack(u)

	msg, sys, err := u.GetMsgAndSys()
	require.NoError(t, err)
	require.Equal(t, network7.MsgGameSvVoteOptionListAdd, msg)
	require.Equal(t, false, sys)

	listAdd := &messages7.SvVoteOptionListAdd{ChunkHeader: header}
	listAdd.Unpack(u)

	require.Equal(t, 1, listAdd.NumOptions)
	require.Equal(t, 1, len(listAdd.Descriptions))
	require.Equal(t, "foo", listAdd.Descriptions[0])
	require.Equal(t, 0, u.RemainingSize())

	// pack
	require.Equal(t, []byte{1, 'f', 'o', 'o', 0x00}, listAdd.Pack())
}

package messages7_test

import (
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

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

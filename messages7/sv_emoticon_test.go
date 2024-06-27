package messages7_test

import (
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

func TestFullPacket(t *testing.T) {
	t.Parallel()
	packet := protocol7.Packet{}
	packet.Messages = append(
		packet.Messages,
		&messages7.SvEmoticon{
			ClientId: 0,
			Emoticon: network7.EmoteGhost,
		},
	)

	{
		// if this test breaks because the session tokens are actually used
		// this is not necessarily a bad thing
		session := &protocol7.Session{
			ServerToken: [4]byte{0x55, 0x55, 0x55, 0x55},
			ClientToken: [4]byte{0xfa, 0xfa, 0xfa, 0xfa},
		}

		want := []byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x40, 0x03, 0x01, 0x14, 0x00, 0x07}
		got := packet.Pack(session)

		require.Equal(t, want, got)
	}
}

func TestSvEmoticonStandalone(t *testing.T) {
	t.Parallel()
	// simple pack
	emoticon := &messages7.SvEmoticon{
		ClientId: 0,
		Emoticon: network7.EmoteExclamation,
	}

	{
		want := []byte{0x00, 0x01}
		got := emoticon.Pack()

		require.Equal(t, want, got)
	}

	// repack
	u := &packer.Unpacker{}
	u.Reset(emoticon.Pack())
	emoticon.Unpack(u)

	{
		want := network7.EmoteExclamation
		got := emoticon.Emoticon
		require.Equal(t, want, got)
	}
}

func TestSvEmoticonStandaloneCrazyGirlEdition(t *testing.T) {
	t.Parallel()
	// simple pack
	emoticon := &messages7.SvEmoticon{
		ClientId: -99999,
		Emoticon: 999,
	}

	{
		want := []byte{222, 154, 12, 167, 15}
		got := emoticon.Pack()

		require.Equal(t, want, got)
	}

	// repack
	u := &packer.Unpacker{}
	u.Reset(emoticon.Pack())
	emoticon.Unpack(u)

	{
		want := network7.Emote(999)
		got := emoticon.Emoticon
		require.Equal(t, want, got)
	}
}

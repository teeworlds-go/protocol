package messages7

import (
	"fmt"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SvGameMsg struct {
	ChunkHeader *chunk7.ChunkHeader

	GameMsgId  network7.GameMsg
	Parameters [3]int
}

func (msg *SvGameMsg) MsgId() int {
	return network7.MsgGameSvGameMsg
}

func (msg *SvGameMsg) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvGameMsg) System() bool {
	return false
}

func (msg *SvGameMsg) Vital() bool {
	return true
}

func (msg *SvGameMsg) Pack() []byte {
	if msg.GameMsgId == network7.GameMsgTeamSwap ||
		msg.GameMsgId == network7.GameMsgSpecInvalidId ||
		msg.GameMsgId == network7.GameMsgTeamShuffle ||
		msg.GameMsgId == network7.GameMsgTeamBalance ||
		msg.GameMsgId == network7.GameMsgCtfDrop ||
		msg.GameMsgId == network7.GameMsgCtfReturn {
		return []byte{}
	}
	if msg.GameMsgId == network7.GameMsgTeamAll ||
		msg.GameMsgId == network7.GameMsgTeamBalanceVictim ||
		msg.GameMsgId == network7.GameMsgCtfGrab ||
		msg.GameMsgId == network7.GameMsgGamePaused {
		return packer.PackInt(msg.Parameters[0])
	}
	if msg.GameMsgId == network7.GameMsgCtfCapture {
		return slices.Concat(
			packer.PackInt(msg.Parameters[0]),
			packer.PackInt(msg.Parameters[1]),
			packer.PackInt(msg.Parameters[2]),
		)
	}
	// TODO: return error if message id is unknown
	return []byte{}
}

func (msg *SvGameMsg) Unpack(u *packer.Unpacker) error {
	msg.GameMsgId = network7.GameMsg(u.GetInt())
	if msg.GameMsgId == network7.GameMsgTeamSwap ||
		msg.GameMsgId == network7.GameMsgSpecInvalidId ||
		msg.GameMsgId == network7.GameMsgTeamShuffle ||
		msg.GameMsgId == network7.GameMsgTeamBalance ||
		msg.GameMsgId == network7.GameMsgCtfDrop ||
		msg.GameMsgId == network7.GameMsgCtfReturn {
		return nil
	}
	if msg.GameMsgId == network7.GameMsgTeamAll ||
		msg.GameMsgId == network7.GameMsgTeamBalanceVictim ||
		msg.GameMsgId == network7.GameMsgCtfGrab ||
		msg.GameMsgId == network7.GameMsgGamePaused {
		msg.Parameters[0] = u.GetInt()
		return nil
	}
	if msg.GameMsgId == network7.GameMsgCtfCapture {
		msg.Parameters[0] = u.GetInt()
		msg.Parameters[1] = u.GetInt()
		msg.Parameters[2] = u.GetInt()
		return nil
	}
	return fmt.Errorf("failed to unpack SvGameMsg GameMsgId=%d is unknown", msg.GameMsgId)
}

func (msg *SvGameMsg) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvGameMsg) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

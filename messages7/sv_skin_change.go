package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvSkinChange struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId              int
	Body                  string
	Marking               string
	Decoration            string
	Hands                 string
	Feet                  string
	Eyes                  string
	CustomColorBody       bool
	CustomColorMarking    bool
	CustomColorDecoration bool
	CustomColorHands      bool
	CustomColorFeet       bool
	CustomColorEyes       bool
	ColorBody             int
	ColorMarking          int
	ColorDecoration       int
	ColorHands            int
	ColorFeet             int
	ColorEyes             int
}

func (msg *SvSkinChange) MsgId() int {
	return network7.MsgGameSvSkinChange
}

func (msg *SvSkinChange) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvSkinChange) System() bool {
	return false
}

func (msg *SvSkinChange) Vital() bool {
	return true
}

func (msg *SvSkinChange) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.ClientId),
		packer.PackStr(msg.Body),
		packer.PackStr(msg.Marking),
		packer.PackStr(msg.Decoration),
		packer.PackStr(msg.Hands),
		packer.PackStr(msg.Feet),
		packer.PackStr(msg.Eyes),
		packer.PackBool(msg.CustomColorBody),
		packer.PackBool(msg.CustomColorMarking),
		packer.PackBool(msg.CustomColorDecoration),
		packer.PackBool(msg.CustomColorHands),
		packer.PackBool(msg.CustomColorFeet),
		packer.PackBool(msg.CustomColorEyes),
		packer.PackInt(msg.ColorBody),
		packer.PackInt(msg.ColorMarking),
		packer.PackInt(msg.ColorDecoration),
		packer.PackInt(msg.ColorHands),
		packer.PackInt(msg.ColorFeet),
		packer.PackInt(msg.ColorEyes),
	)
}

func (msg *SvSkinChange) Unpack(u *packer.Unpacker) error {
	msg.ClientId = u.GetInt()
	msg.Body, _ = u.GetString()
	msg.Marking, _ = u.GetString()
	msg.Decoration, _ = u.GetString()
	msg.Hands, _ = u.GetString()
	msg.Feet, _ = u.GetString()
	msg.Eyes, _ = u.GetString()
	msg.CustomColorBody = u.GetInt() != 0
	msg.CustomColorMarking = u.GetInt() != 0
	msg.CustomColorDecoration = u.GetInt() != 0
	msg.CustomColorHands = u.GetInt() != 0
	msg.CustomColorFeet = u.GetInt() != 0
	msg.CustomColorEyes = u.GetInt() != 0
	msg.ColorBody = u.GetInt()
	msg.ColorMarking = u.GetInt()
	msg.ColorDecoration = u.GetInt()
	msg.ColorHands = u.GetInt()
	msg.ColorFeet = u.GetInt()
	msg.ColorEyes = u.GetInt()
	return nil
}

func (msg *SvSkinChange) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvSkinChange) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

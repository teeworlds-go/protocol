package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type ClStartInfo struct {
	ChunkHeader *chunk7.ChunkHeader

	Name                  string
	Clan                  string
	Country               int
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

func (info ClStartInfo) MsgId() int {
	return network7.MsgGameClStartInfo
}

func (info ClStartInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (info ClStartInfo) System() bool {
	return false
}

func (info ClStartInfo) Vital() bool {
	return true
}

func (info ClStartInfo) Pack() []byte {
	return slices.Concat(
		packer.PackStr(info.Name),
		packer.PackStr(info.Clan),
		packer.PackInt(info.Country),
		packer.PackStr(info.Body),
		packer.PackStr(info.Marking),
		packer.PackStr(info.Decoration),
		packer.PackStr(info.Hands),
		packer.PackStr(info.Feet),
		packer.PackStr(info.Eyes),
		packer.PackBool(info.CustomColorBody),
		packer.PackBool(info.CustomColorMarking),
		packer.PackBool(info.CustomColorDecoration),
		packer.PackBool(info.CustomColorHands),
		packer.PackBool(info.CustomColorFeet),
		packer.PackBool(info.CustomColorEyes),
		packer.PackInt(info.ColorBody),
		packer.PackInt(info.ColorMarking),
		packer.PackInt(info.ColorDecoration),
		packer.PackInt(info.ColorHands),
		packer.PackInt(info.ColorFeet),
		packer.PackInt(info.ColorEyes),
	)
}

func (info *ClStartInfo) Unpack(u *packer.Unpacker) {
	info.Name = u.GetString()
	info.Clan = u.GetString()
	info.Country = u.GetInt()
	info.Body = u.GetString()
	info.Marking = u.GetString()
	info.Decoration = u.GetString()
	info.Hands = u.GetString()
	info.Feet = u.GetString()
	info.Eyes = u.GetString()
	info.CustomColorBody = u.GetInt() != 0
	info.CustomColorMarking = u.GetInt() != 0
	info.CustomColorDecoration = u.GetInt() != 0
	info.CustomColorHands = u.GetInt() != 0
	info.CustomColorFeet = u.GetInt() != 0
	info.CustomColorEyes = u.GetInt() != 0
	info.ColorBody = u.GetInt()
	info.ColorMarking = u.GetInt()
	info.ColorDecoration = u.GetInt()
	info.ColorHands = u.GetInt()
	info.ColorFeet = u.GetInt()
	info.ColorEyes = u.GetInt()
}

func (msg *ClStartInfo) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *ClStartInfo) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

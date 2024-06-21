package messages7

import (
	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type SvClientInfo struct {
	ChunkHeader *chunk7.ChunkHeader

	ClientId              int
	Local                 bool
	Team                  int
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
	ColorHands            int
	ColorFeet             int
	ColorEyes             int
	Silent                bool
}

func (info *SvClientInfo) MsgId() int {
	return network7.MsgGameSvClientInfo
}

func (info *SvClientInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (info *SvClientInfo) System() bool {
	return false
}

func (info *SvClientInfo) Vital() bool {
	return true
}

func (info *SvClientInfo) Unpack(u *packer.Unpacker) {
	info.ClientId = u.GetInt()
	info.Local = u.GetInt() != 0
	info.Team = u.GetInt()
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
	info.ColorHands = u.GetInt()
	info.ColorFeet = u.GetInt()
	info.ColorEyes = u.GetInt()
	info.Silent = u.GetInt() != 0
}

func (msg *SvClientInfo) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvClientInfo) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

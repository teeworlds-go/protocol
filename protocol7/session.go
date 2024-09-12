package protocol7

import "github.com/teeworlds-go/protocol/messages7"

// NewSession creates a new session with default values that can be used to request a new token from the server.
func NewSession() Session {
	return Session{
		ClientToken: [4]byte{0x01, 0x02, 0x03, 0x04},
		ServerToken: [4]byte{0xff, 0xff, 0xff, 0xff}, // special value for requesting a new token
		Ack:         0,
		Sequence:    0,
		PeerAck:     0,
	}
}

// teeworlds low level protocol
// keeping track of connection state
// resends and anti spoof tokens
type Session struct {
	ClientToken [4]byte
	ServerToken [4]byte

	// The amount of vital chunks received
	Ack int

	// The amount of vital chunks sent
	Sequence int

	// The amount of vital chunks acknowledged by the peer
	PeerAck int
}

// TODO: should this be removed? All of this could be set in Packet.Pack()
func (connection *Session) BuildResponse() *Packet {
	return &Packet{
		Header: PacketHeader{
			Flags: PacketFlags{
				Connless:    false,
				Compression: false,
				Resend:      false,
				Control:     false,
			},
			Ack:       0, // will be set in Packet.Pack()
			NumChunks: 0, // will be set in Packet.Pack()
			Token:     connection.ServerToken,
		},
	}
}

func (connection *Session) CtrlToken() *Packet {
	response := connection.BuildResponse()
	response.Header.Flags.Control = true
	response.Messages = append(
		response.Messages,
		&messages7.CtrlToken{
			Token: connection.ClientToken,
		},
	)
	return response
}

func (client *Session) MsgStartInfo() *messages7.ClStartInfo {
	return &messages7.ClStartInfo{
		Name:                  "gopher",
		Clan:                  "",
		Country:               0,
		Body:                  "greensward",
		Marking:               "duodonny",
		Decoration:            "",
		Hands:                 "standard",
		Feet:                  "standard",
		Eyes:                  "standard",
		CustomColorBody:       false,
		CustomColorMarking:    false,
		CustomColorDecoration: false,
		CustomColorHands:      false,
		CustomColorFeet:       false,
		CustomColorEyes:       false,
		ColorBody:             0,
		ColorMarking:          0,
		ColorDecoration:       0,
		ColorHands:            0,
		ColorFeet:             0,
		ColorEyes:             0,
	}
}

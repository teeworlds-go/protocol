package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Info struct {
	ChunkHeader *chunk7.ChunkHeader

	// The official name is "NetVersion" but a more fitting name in my opinion would be "Protocol Version".
	// The variable C++ implementations GAME_NETVERSION always expands to "0.7 802f1be60a05665f"
	// If the server gets another string it actually rejects the connection. This is what prohibits 0.6 clients to join 0.7 servers.
	//
	// Recommended value is network7.NetVersion
	Version string

	// Password to enter password protected servers
	// If the server does not require a password it will ignore this string
	//
	// Recommended value is ""
	Password string

	// Another version field which does not have to match the servers version to establish a connection.
	// The first version field makes sure that client and server use the same major protocol and are compatible.
	// This "Client Version" field then informs the server about the clients minor version.
	// The server can use it to activate some non protocol breaking features that were introduced in minor releases.
	//
	// The official teeworlds 0.7.5 client sends the value 0x0705
	// So the recommended value is network7.ClientVersion
	ClientVersion int
}

func (msg Info) MsgId() int {
	return network7.MsgSysInfo
}

func (info Info) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg Info) System() bool {
	return true
}

func (msg Info) Vital() bool {
	return true
}

func (msg Info) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Version),
		packer.PackStr(msg.Password),
		packer.PackInt(msg.ClientVersion),
	)
}

func (msg *Info) Unpack(u *packer.Unpacker) {
	msg.Version = u.GetString()
	msg.Password = u.GetString()
	msg.ClientVersion = u.GetInt()
}

func (msg *Info) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *Info) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

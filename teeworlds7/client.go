package teeworlds7

import (
	"log"
	"net"

	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

type Player struct {
	Info messages7.SvClientInfo
}

type Game struct {
	Players []Player
}

type Client struct {
	Name    string
	Clan    string
	Country int

	// chunks to be sent on next packet send
	// use client.SendMessage() to put your chunks here
	QueuedMessages []messages7.NetMessage

	// hooks from the user
	Callbacks UserMsgCallbacks

	// udp connection
	Conn net.Conn

	// teeworlds session
	Session protocol7.Session

	// teeworlds game state
	Game Game
}

func (client *Client) throwError(err error) {
	if client.Callbacks.InternalError != nil {
		client.Callbacks.InternalError(err)
		return
	}
	log.Fatal(err)
}

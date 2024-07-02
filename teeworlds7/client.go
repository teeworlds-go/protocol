package teeworlds7

import (
	"log"
	"net"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

type Player struct {
	Info messages7.SvClientInfo
}

type Game struct {
	Players []Player
	Input   *messages7.Input
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

	// old snapshots used to unpack new deltas
	SnapshotStorage *snapshot7.Storage

	// for snapshot syncing
	LastPredTime *time.Time

	// teeworlds game state
	Game Game
}

func NewClient() *Client {
	client := &Client{}
	client.SnapshotStorage = snapshot7.NewStorage()
	client.Game.Input = &messages7.Input{}
	return client
}

func (client *Client) throwError(err error) {
	for _, callback := range client.Callbacks.InternalError {
		if callback(err) == false {
			return
		}
	}

	log.Fatal(err)
}

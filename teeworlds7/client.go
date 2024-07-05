package teeworlds7

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

type Player struct {
	Info messages7.SvClientInfo
}

type Game struct {
	Players       []Player
	Snap          *GameSnap
	Input         *messages7.Input
	LastSentInput messages7.Input
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

	// when the last packet was sent
	// tracked to know when to send keepalives
	LastSend      time.Time
	LastInputSend time.Time

	// teeworlds session
	Session protocol7.Session

	// old snapshots used to unpack new deltas
	SnapshotStorage *snapshot7.Storage

	// teeworlds game state
	Game Game

	// might be -1 if we do not know our own id yet
	LocalClientId int
}

// TODO: add this for all items and move it to a different file
func (client *Client) SnapFindCharacter(ClientId int) (*object7.Character, error) {
	item, err := client.SnapshotStorage.FindAltSnapItem(network7.ObjCharacter, ClientId)
	if err != nil {
		return nil, err
	}
	character, ok := item.(*object7.Character)
	if ok == false {
		return nil, errors.New("failed to cast character")
	}
	return character, nil
}

func NewClient() *Client {
	client := &Client{}
	client.SnapshotStorage = snapshot7.NewStorage()
	client.Game.Snap = &GameSnap{}
	client.Game.Input = &messages7.Input{}
	client.LocalClientId = -1
	client.LastSend = time.Now()
	return client
}

func (client *Client) sendInputIfNeeded() bool {
	diff := time.Now().Sub(client.LastSend)
	send := false
	// at least every 10hz or on change
	if diff.Microseconds() > 1000000 {
		send = true
	}
	if client.Game.LastSentInput != *client.Game.Input {
		send = true
	}

	if send {
		client.SendInput()
	}

	return send
}

func (client *Client) gameTick() {
	defaultAction := func() {
		if client.sendInputIfNeeded() == true {
			return
		}

		diff := time.Now().Sub(client.LastSend)
		if diff.Seconds() > 2 {
			client.SendKeepAlive()
		}
	}

	for _, callback := range client.Callbacks.Tick {
		callback(defaultAction)
	}
}

func (client *Client) throwError(err error) {
	for _, callback := range client.Callbacks.InternalError {
		if callback(err) == false {
			return
		}
	}

	log.Fatal(err)
}

package teeworlds7

import (
	"fmt"
	"net"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

const (
	UnknownClientId = -1
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
//
//	this would be more useful to have on the Snapshot struct directly
//	so it can be used everywhere not only in a client
//	and the client then can just wrap it to acces the alt snap
func (client *Client) SnapFindCharacter(clientId int) (character *object7.Character, found bool, err error) {
	item, found, err := client.SnapshotStorage.FindAltSnapItem(network7.ObjCharacter, clientId)
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	character, ok := item.(*object7.Character)
	if !ok {
		panic(fmt.Sprintf("type assertion failed: found client snap item is not a *object7.Character: %T", item))
	}
	return character, true, nil
}

func NewClient() *Client {
	return &Client{
		SnapshotStorage: snapshot7.NewStorage(),
		Game: Game{
			Snap:  &GameSnap{},
			Input: &messages7.Input{},
		},
		LocalClientId: UnknownClientId,
		LastSend:      time.Now(),
	}
}

func (client *Client) sendInputIfNeeded() (sent bool, err error) {
	send := false
	// at least every 10hz or on change
	if time.Since(client.LastSend) > 100*time.Millisecond {
		send = true
	} else if client.Game.Input != nil && client.Game.LastSentInput != *client.Game.Input {
		send = true
	}

	if !send {
		return false, nil
	}

	err = client.SendInput()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (client *Client) gameTick() error {
	defaultAction := func() error {

		// either input or keepalive
		sent, err := client.sendInputIfNeeded()
		if err != nil {
			// TODO: FIXME: propagate error correctly back to the caller
			return fmt.Errorf("failed to send input: %w", err)
		} else if sent {
			return nil
		}

		// keepalive in case we did not send anything
		// rounded to seconds, which is why at least 3 seconds need to pass before
		// another keepalive is sent
		if time.Since(client.LastSend).Seconds() > 2 {
			err = client.SendKeepAlive()
			if err != nil {
				return fmt.Errorf("failed to send keepalive: %w", err)

			}
		}
		return nil
	}

	var err error
	for _, callback := range client.Callbacks.Tick {
		err = callback(defaultAction)
		if err != nil {
			return err
		}
	}
	return nil
}

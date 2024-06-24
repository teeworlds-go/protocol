package object7

import (
	"log"

	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type SnapObject interface {
	Id() int
	Type() int

	// number of packed integers
	// not the number of bytes
	Size() int

	Pack() []byte
	Unpack(u *packer.Unpacker) error
}

// Comes without payload
// you have to call item.Unpack(u) manually after getting it
func NewObject(itemType int, itemId int) SnapObject {
	if itemType == network7.ObjFlag {
		return &Flag{ItemId: itemId}
	} else if itemType == network7.ObjPickup {
		return &Pickup{ItemId: itemId}
	} else if itemType == network7.ObjGameData {
		return &GameData{ItemId: itemId}
	} else if itemType == network7.ObjGameDataTeam {
		return &GameDataTeam{ItemId: itemId}
	} else if itemType == network7.ObjGameDataFlag {
		return &GameDataFlag{ItemId: itemId}
	} else if itemType == network7.ObjCharacter {
		return &Character{ItemId: itemId}
	} else if itemType == network7.ObjPlayerInfo {
		return &PlayerInfo{ItemId: itemId}
	}

	// TODO: remove this is just for debugging
	log.Fatalf("unknown item type %d\n", itemType)

	unknown := &Unknown{
		ItemId:   itemId,
		ItemType: itemType,
	}
	return unknown
}

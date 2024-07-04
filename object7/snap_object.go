package object7

import (
	"fmt"
	"log"

	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SnapObject interface {
	// id separating this snap item from other items with same type
	// the ids are unique per type
	// for players it matches their ClientId
	Id() int

	// type of the snap item
	TypeId() int

	// number of packed integers
	// not the number of bytes
	//
	// only counting the payload
	// not the type id or item id
	// and also not the optional size field (game data race, player info race)
	Size() int

	Pack() []byte
	Unpack(u *packer.Unpacker) error
}

// Comes without payload
// you have to call item.Unpack(u) manually after getting it
//
// it might consume one integer for the size field of the given unpacker
// if it is an item with size field
func NewObject(typeId int, itemId int, u *packer.Unpacker) SnapObject {
	if typeId == network7.ObjPlayerInput {
		return &PlayerInput{ItemId: itemId}
	} else if typeId == network7.ObjProjectile {
		return &Projectile{ItemId: itemId}
	} else if typeId == network7.ObjLaser {
		return &Laser{ItemId: itemId}
	} else if typeId == network7.ObjPickup {
		return &Pickup{ItemId: itemId}
	} else if typeId == network7.ObjFlag {
		return &Flag{ItemId: itemId}
	} else if typeId == network7.ObjGameData {
		return &GameData{ItemId: itemId}
	} else if typeId == network7.ObjGameDataTeam {
		return &GameDataTeam{ItemId: itemId}
	} else if typeId == network7.ObjGameDataFlag {
		return &GameDataFlag{ItemId: itemId}
	} else if typeId == network7.ObjCharacter {
		return &Character{ItemId: itemId}
	} else if typeId == network7.ObjPlayerInfo {
		return &PlayerInfo{ItemId: itemId}
	} else if typeId == network7.ObjSpectatorInfo {
		return &SpectatorInfo{ItemId: itemId}
	} else if typeId == network7.ObjDeClientInfo {
		return &DeClientInfo{ItemId: itemId}
	} else if typeId == network7.ObjDeGameInfo {
		return &DeGameInfo{ItemId: itemId}
	} else if typeId == network7.ObjDeTuneParams {
		return &DeTuneParams{ItemId: itemId}
	} else if typeId == network7.ObjExplosion {
		return &Explosion{ItemId: itemId}
	} else if typeId == network7.ObjSpawn {
		return &Spawn{ItemId: itemId}
	} else if typeId == network7.ObjHammerHit {
		return &HammerHit{ItemId: itemId}
	} else if typeId == network7.ObjDeath {
		return &Death{ItemId: itemId}
	} else if typeId == network7.ObjSoundWorld {
		return &SoundWorld{ItemId: itemId}
	} else if typeId == network7.ObjDamage {
		return &Damage{ItemId: itemId}
	} else if typeId == network7.ObjPlayerInfoRace {
		race := &PlayerInfoRace{ItemId: itemId}
		size := u.GetInt()
		if size != race.Size() {
			log.Panicf("got race info with size %d but expected size %d\n", size, race.Size())
		}
		return race
	} else if typeId == network7.ObjGameDataRace {
		race := &GameDataRace{ItemId: itemId}
		size := u.GetInt()
		fmt.Printf("got gamedata race red size=%d remaining unpacker data=%x\n", size, u.RemainingData())
		if size != race.Size() {
			log.Panicf("got game data race with size %d but expected size %d\n", size, race.Size())
		}
		return race
	}

	// TODO: add this panic and remove it again once all tests pass
	// log.Panicf("unknown item type %d\n", typeId)

	unknown := &Unknown{
		ItemId:   itemId,
		ItemType: typeId,
		ItemSize: u.GetInt(),
	}
	return unknown
}

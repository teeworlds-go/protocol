package object7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
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
	Size() int

	Pack() []byte
	Unpack(u *packer.Unpacker) error
}

// Comes without payload
// you have to call item.Unpack(u) manually after getting it
func NewObject(typeId int, itemId int) SnapObject {
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
		return &PlayerInfoRace{ItemId: itemId}
	} else if typeId == network7.ObjGameDataRace {
		return &GameDataRace{ItemId: itemId}
	}

	// TODO: add this panic and remove it again once all tests pass
	// log.Panicf("unknown item type %d\n", typeId)

	unknown := &Unknown{
		ItemId:   itemId,
		ItemType: typeId,
	}
	return unknown
}

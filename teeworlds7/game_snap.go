package teeworlds7

import (
	"log/slog"

	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

type GameSnap struct {
	PlayerInputs    []*object7.PlayerInput
	Projectiles     []*object7.Projectile
	Lasers          []*object7.Laser
	Pickups         []*object7.Pickup
	Flags           []*object7.Flag
	GameDatas       []*object7.GameData
	GameDataTeams   []*object7.GameDataTeam
	GameDataFlags   []*object7.GameDataFlag
	Characters      []*object7.Character
	PlayerInfos     []*object7.PlayerInfo
	SpectatorInfos  []*object7.SpectatorInfo
	DeClientInfos   []*object7.DeClientInfo
	DeGameInfos     []*object7.DeGameInfo
	DeTuneParamss   []*object7.DeTuneParams
	Explosions      []*object7.Explosion
	Spawns          []*object7.Spawn
	HammerHits      []*object7.HammerHit
	Deaths          []*object7.Death
	SoundWorlds     []*object7.SoundWorld
	Damages         []*object7.Damage
	PlayerInfoRaces []*object7.PlayerInfoRace
	GameDataRaces   []*object7.GameDataRace
	Unknowns        []*object7.Unknown
}

func (gs *GameSnap) fill(snap *snapshot7.Snapshot) {
	// reset length to refill all new characters
	// but keep capacity to avoid reallocation
	// when the amount of characters did not change
	gs.Characters = gs.Characters[:0]
	gs.PlayerInputs = gs.PlayerInputs[:0]
	gs.Projectiles = gs.Projectiles[:0]
	gs.Lasers = gs.Lasers[:0]
	gs.Pickups = gs.Pickups[:0]
	gs.Flags = gs.Flags[:0]
	gs.GameDatas = gs.GameDatas[:0]
	gs.GameDataTeams = gs.GameDataTeams[:0]
	gs.GameDataFlags = gs.GameDataFlags[:0]
	gs.Characters = gs.Characters[:0]
	gs.PlayerInfos = gs.PlayerInfos[:0]
	gs.SpectatorInfos = gs.SpectatorInfos[:0]
	gs.DeClientInfos = gs.DeClientInfos[:0]
	gs.DeGameInfos = gs.DeGameInfos[:0]
	gs.DeTuneParamss = gs.DeTuneParamss[:0]
	gs.Explosions = gs.Explosions[:0]
	gs.Spawns = gs.Spawns[:0]
	gs.HammerHits = gs.HammerHits[:0]
	gs.Deaths = gs.Deaths[:0]
	gs.SoundWorlds = gs.SoundWorlds[:0]
	gs.Damages = gs.Damages[:0]
	gs.PlayerInfoRaces = gs.PlayerInfoRaces[:0]
	gs.GameDataRaces = gs.GameDataRaces[:0]
	gs.Unknowns = gs.Unknowns[:0]

	for _, snapItem := range snap.Items {
		switch item := snapItem.(type) {
		case *object7.PlayerInput:
			gs.PlayerInputs = append(gs.PlayerInputs, item)
		case *object7.Projectile:
			gs.Projectiles = append(gs.Projectiles, item)
		case *object7.Laser:
			gs.Lasers = append(gs.Lasers, item)
		case *object7.Pickup:
			gs.Pickups = append(gs.Pickups, item)
		case *object7.Flag:
			gs.Flags = append(gs.Flags, item)
		case *object7.GameData:
			gs.GameDatas = append(gs.GameDatas, item)
		case *object7.GameDataTeam:
			gs.GameDataTeams = append(gs.GameDataTeams, item)
		case *object7.GameDataFlag:
			gs.GameDataFlags = append(gs.GameDataFlags, item)
		case *object7.Character:
			gs.Characters = append(gs.Characters, item)
		case *object7.PlayerInfo:
			gs.PlayerInfos = append(gs.PlayerInfos, item)
		case *object7.SpectatorInfo:
			gs.SpectatorInfos = append(gs.SpectatorInfos, item)
		case *object7.DeClientInfo:
			gs.DeClientInfos = append(gs.DeClientInfos, item)
		case *object7.DeGameInfo:
			gs.DeGameInfos = append(gs.DeGameInfos, item)
		case *object7.DeTuneParams:
			gs.DeTuneParamss = append(gs.DeTuneParamss, item)
		case *object7.Explosion:
			gs.Explosions = append(gs.Explosions, item)
		case *object7.Spawn:
			gs.Spawns = append(gs.Spawns, item)
		case *object7.HammerHit:
			gs.HammerHits = append(gs.HammerHits, item)
		case *object7.Death:
			gs.Deaths = append(gs.Deaths, item)
		case *object7.SoundWorld:
			gs.SoundWorlds = append(gs.SoundWorlds, item)
		case *object7.Damage:
			gs.Damages = append(gs.Damages, item)
		case *object7.PlayerInfoRace:
			gs.PlayerInfoRaces = append(gs.PlayerInfoRaces, item)
		case *object7.GameDataRace:
			gs.GameDataRaces = append(gs.GameDataRaces, item)
		case *object7.Unknown:
			gs.Unknowns = append(gs.Unknowns, item)
		default:
			slog.Debug("snap item not added to game snap", "type", snapItem.TypeId())
		}
	}
}

package network7

const (
	MaxClients    = 64
	NetVersion    = "0.7 802f1be60a05665f"
	ClientVersion = 0x0705

	ChatAll     ChatMode = 1
	ChatTeam    ChatMode = 2
	ChatWhisper ChatMode = 3

	TeamSpectators GameTeam = -1
	TeamRed        GameTeam = 0
	TeamBlue       GameTeam = 1

	SpecFreeView Spec = 0
	SpecPlayer   Spec = 1
	SpecFlagRed  Spec = 2
	SpecFlagBlue Spec = 3

	VoteChoiceNo  VoteChoice = -1
	VoteChoiceYes VoteChoice = 1

	VoteUnknown   Vote = 0
	VoteStartOp   Vote = 1
	VoteStartKick Vote = 2
	VoteStartSpec Vote = 3
	VoteEndAbort  Vote = 4
	VoteEndPass   Vote = 5
	VoteEndFail   Vote = 6

	PickupHealth  Pickup = 0
	PickupArmor   Pickup = 1
	PickupGrenade Pickup = 2
	PickupShotgun Pickup = 3
	PickupLaser   Pickup = 4
	PickupNinja   Pickup = 5
	PickupGun     Pickup = 6
	PickupHammer  Pickup = 7

	GamestateflagWarmup         GameStateFlag = 1
	GamestateflagSuddendeath    GameStateFlag = 2
	GamestateflagRoundover      GameStateFlag = 4
	GamestateflagGameover       GameStateFlag = 8
	GamestateflagPaused         GameStateFlag = 16
	GamestateflagStartcountdown GameStateFlag = 32

	GameMsgTeamSwap          GameMsg = 0
	GameMsgSpecInvalidId     GameMsg = 1
	GameMsgTeamShuffle       GameMsg = 2
	GameMsgTeamBalance       GameMsg = 3
	GameMsgCtfDrop           GameMsg = 4
	GameMsgCtfReturn         GameMsg = 5
	GameMsgTeamAll           GameMsg = 6
	GameMsgTeamBalanceVictim GameMsg = 7
	GameMsgCtfGrab           GameMsg = 8
	GameMsgCtfCapture        GameMsg = 9
	GameMsgGamePaused        GameMsg = 10

	EyeEmoteNormal   EyeEmote = 0
	EyeEmotePain     EyeEmote = 1
	EyeEmoteHappy    EyeEmote = 2
	EyeEmoteSurprise EyeEmote = 3
	EyeEmoteAngry    EyeEmote = 4
	EyeEmoteBlink    EyeEmote = 5
	NumEyeEmotes     EyeEmote = 6

	// oop!
	EmoticonOop Emoticon = 0
	// !
	EmoticonExclamation Emoticon = 1
	EmoticonHearts      Emoticon = 2
	// tear
	EmoticonDrop Emoticon = 3
	// ...
	EmoticonDotdot Emoticon = 4
	EmoticonMusic  Emoticon = 5
	EmoticonSorry  Emoticon = 6
	EmoticonGhost  Emoticon = 7
	// annoyed
	EmoticonSushi Emoticon = 8
	// angry
	EmoticonSplattee Emoticon = 9
	EmoticonDeviltee Emoticon = 10
	// swearing
	EmoticonZomg Emoticon = 11
	EmoticonZzz  Emoticon = 12
	EmoticonWtf  Emoticon = 13
	// happy
	EmoticonEyes Emoticon = 14
	// ??
	EmoticonQuestion Emoticon = 15

	MsgCtrlKeepAlive = 0x00
	MsgCtrlConnect   = 0x01
	MsgCtrlAccept    = 0x02
	MsgCtrlToken     = 0x05
	MsgCtrlClose     = 0x04

	ObjInvalid        = 0
	ObjPlayerInput    = 1
	ObjProjectile     = 2
	ObjLaser          = 3
	ObjPickup         = 4
	ObjFlag           = 5
	ObjGameData       = 6
	ObjGameDataTeam   = 7
	ObjGameDataFlag   = 8
	ObjCharacterCore  = 9
	ObjCharacter      = 10
	ObjPlayerInfo     = 11
	ObjSpectatorInfo  = 12
	ObjDeClientInfo   = 13
	ObjDeGameInfo     = 14
	ObjDeTuneParams   = 15
	ObjCommon         = 16
	ObjExplosion      = 17
	ObjSpawn          = 18
	ObjHammerHit      = 19
	ObjDeath          = 20
	ObjSoundWorld     = 21
	ObjDamage         = 22
	ObjPlayerInfoRace = 23
	ObjGameDataRace   = 24
	NumNetobjtypes    = 25

	// TODO: these should preferrably all be devide dinto different type dintegers
	// same as ChatMode, etc. so that the user can easily see which integer to pass
	// to which function as which parameter
	MsgSysInfo            = 1
	MsgSysMapChange       = 2
	MsgSysMapData         = 3
	MsgSysServerInfo      = 4
	MsgSysConReady        = 5
	MsgSysSnap            = 6
	MsgSysSnapEmpty       = 7
	MsgSysSnapSingle      = 8
	MsgSysSnapSmall       = 9
	MsgSysInputTiming     = 10
	MsgSysRconAuthOn      = 11
	MsgSysRconAuthOff     = 12
	MsgSysRconLine        = 13
	MsgSysRconCmdAdd      = 14
	MsgSysRconCmdRem      = 15
	MsgSysAuthChallenge   = 16 // unused
	MsgSysAuthResult      = 17 // unused
	MsgSysReady           = 18
	MsgSysEnterGame       = 19
	MsgSysInput           = 20
	MsgSysRconCmd         = 21
	MsgSysRconAuth        = 22
	MsgSysRequestMapData  = 23
	MsgSysAuthStart       = 24 // unused
	MsgSysAuthResponse    = 25 // unused
	MsgSysPing            = 26
	MsgSysPingReply       = 27
	MsgSysError           = 28 // unused
	MsgSysMaplistEntryAdd = 29
	MsgSysMaplistEntryRem = 30

	MsgGameSvMotd              = 1
	MsgGameSvBroadcast         = 2
	MsgGameSvChat              = 3
	MsgGameSvTeam              = 4
	MsgGameSvKillMsg           = 5
	MsgGameSvTuneParams        = 6
	MsgGameSvExtraProjectile   = 7 // unused
	MsgGameSvReadyToEnter      = 8
	MsgGameSvWeaponPickup      = 9
	MsgGameSvEmoticon          = 10
	MsgGameSvVoteClearOptions  = 11
	MsgGameSvVoteOptionListAdd = 12
	MsgGameSvVoteOptionAdd     = 13
	MsgGameSvVoteOptionRemove  = 14
	MsgGameSvVoteSet           = 15
	MsgGameSvVoteStatus        = 16
	MsgGameSvServerSettings    = 17
	MsgGameSvClientInfo        = 18
	MsgGameSvGameInfo          = 19
	MsgGameSvClientDrop        = 20
	MsgGameSvGameMsg           = 21
	MsgGameDeClientEnter       = 22
	MsgGameDeClientLeave       = 23
	MsgGameClSay               = 24
	MsgGameClSetTeam           = 25
	MsgGameClSetSpectatorMode  = 26
	MsgGameClStartInfo         = 27
	MsgGameClKill              = 28
	MsgGameClReadyChange       = 29
	MsgGameClEmoticon          = 30
	MsgGameClVote              = 31
	MsgGameClCallVote          = 32
	MsgGameSvSkinChange        = 33
	MsgGameClSkinChange        = 34
	MsgGameSvRaceFinish        = 35
	MsgGameSvCheckpoint        = 36
	MsgGameSvCommandInfo       = 37
	MsgGameSvCommandInfoRemove = 38
	MsgGameClCommand           = 39

	TypeControl  MsgType = 1
	TypeNet      MsgType = 2
	TypeConnless MsgType = 3

	// can be sent by the server in kill messages
	WeaponAllGame    WeaponAll = -3
	WeaponAllSelf    WeaponAll = -2
	WeaponAllWorld   WeaponAll = -1
	WeaponAllHammer  WeaponAll = 0
	WeaponAllGun     WeaponAll = 1
	WeaponAllShotgun WeaponAll = 2
	WeaponAllGrenade WeaponAll = 3
	WeaponAllLaser   WeaponAll = 4
	WeaponAllNinja   WeaponAll = 5
	NumAllWeapons    WeaponAll = 6

	// can be sent by the client when requesting weapon switch
	// or by the server in kill messages
	WeaponHammer  Weapon = 0
	WeaponGun     Weapon = 1
	WeaponShotgun Weapon = 2
	WeaponGrenade Weapon = 3
	WeaponLaser   Weapon = 4
	WeaponNinja   Weapon = 5
	NumWeapons    Weapon = 6
)

type GameMsg int
type GameStateFlag int
type Spec int
type VoteChoice int
type Vote int
type Pickup int
type EyeEmote int
type Emoticon int
type ChatMode int
type GameTeam int
type WeaponAll int
type Weapon int

type MsgType int

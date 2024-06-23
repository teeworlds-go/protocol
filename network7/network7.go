package network7

const (
	MaxClients    = 64
	NetVersion    = "0.7 802f1be60a05665f"
	ClientVersion = 0x0705

	MsgCtrlKeepAlive = 0x00
	MsgCtrlConnect   = 0x01
	MsgCtrlAccept    = 0x02
	MsgCtrlToken     = 0x05
	MsgCtrlClose     = 0x04

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

	MsgGameSvMotd       = 1
	MsgGameSvChat       = 3
	MsgGameReadyToEnter = 8
	MsgGameSvClientInfo = 18
	MsgGameClStartInfo  = 27

	TypeControl  MsgType = 1
	TypeNet      MsgType = 2
	TypeConnless MsgType = 3

	WeaponHammer  Weapon = 0
	WeaponGun     Weapon = 1
	WeaponShotgun Weapon = 2
	WeaponGrenade Weapon = 3
	WeaponLaser   Weapon = 4
	WeaponNinja   Weapon = 5
	NumWeapons    Weapon = 6
)

type MsgType int
type Weapon int

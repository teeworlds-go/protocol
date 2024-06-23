package teeworlds7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

// Processes the incoming packet
// It might print to the console
// It might send a response packet
type DefaultAction func()

// TODO: this should be a map but the type checker broke me
//
// // key is the network7.MessageId
// UserMsgCallbacks map[int]UserMsgCallback
type UserMsgCallbacks struct {
	PacketIn      func(*protocol7.Packet)
	PacketOut     func(*protocol7.Packet)
	MsgUnknown    func(*messages7.Unknown, DefaultAction)
	InternalError func(error)

	CtrlKeepAlive func(*messages7.CtrlKeepAlive, DefaultAction)
	CtrlConnect   func(*messages7.CtrlConnect, DefaultAction)
	CtrlAccept    func(*messages7.CtrlAccept, DefaultAction)
	CtrlToken     func(*messages7.CtrlToken, DefaultAction)
	CtrlClose     func(*messages7.CtrlClose, DefaultAction)

	SysInfo            func(*messages7.Info, DefaultAction)
	SysMapChange       func(*messages7.MapChange, DefaultAction)
	SysMapData         func(*messages7.MapData, DefaultAction)
	SysServerInfo      func(*messages7.ServerInfo, DefaultAction)
	SysConReady        func(*messages7.ConReady, DefaultAction)
	SysSnap            func(*messages7.Snap, DefaultAction)
	SysSnapEmpty       func(*messages7.SnapEmpty, DefaultAction)
	SysSnapSingle      func(*messages7.SnapSingle, DefaultAction)
	SysSnapSmall       func(*messages7.SnapSmall, DefaultAction)
	SysInputTiming     func(*messages7.InputTiming, DefaultAction)
	SysRconAuthOn      func(*messages7.RconAuthOn, DefaultAction)
	SysRconAuthOff     func(*messages7.RconAuthOff, DefaultAction)
	SysRconLine        func(*messages7.RconLine, DefaultAction)
	SysRconCmdAdd      func(*messages7.RconCmdAdd, DefaultAction)
	SysRconCmdRem      func(*messages7.RconCmdRem, DefaultAction)
	SysAuthChallenge   func(*messages7.AuthChallenge, DefaultAction)
	SysAuthResult      func(*messages7.AuthResult, DefaultAction)
	SysReady           func(*messages7.Ready, DefaultAction)
	SysEnterGame       func(*messages7.EnterGame, DefaultAction)
	SysInput           func(*messages7.Input, DefaultAction)
	SysRconCmd         func(*messages7.RconCmd, DefaultAction)
	SysRconAuth        func(*messages7.RconAuth, DefaultAction)
	SysRequestMapData  func(*messages7.RequestMapData, DefaultAction)
	SysAuthStart       func(*messages7.AuthStart, DefaultAction)
	SysAuthResponse    func(*messages7.AuthResponse, DefaultAction)
	SysPing            func(*messages7.Ping, DefaultAction)
	SysPingReply       func(*messages7.PingReply, DefaultAction)
	SysError           func(*messages7.Error, DefaultAction)
	SysMaplistEntryAdd func(*messages7.MaplistEntryAdd, DefaultAction)
	SysMaplistEntryRem func(*messages7.MaplistEntryRem, DefaultAction)

	GameSvMotd      func(*messages7.SvMotd, DefaultAction)
	GameSvBroadcast func(*messages7.SvBroadcast, DefaultAction)
	GameSvChat      func(*messages7.SvChat, DefaultAction)
	GameSvTeam      func(*messages7.SvTeam, DefaultAction)
	// GameSvKillMsg           func(*messages7.SvKillMsg, DefaultAction)
	// GameSvTuneParams        func(*messages7.SvTuneParams, DefaultAction)
	// GameSvExtraProjectile   func(*messages7.SvExtraProjectile, DefaultAction)
	GameReadyToEnter func(*messages7.ReadyToEnter, DefaultAction)
	// GameWeaponPickup        func(*messages7.WeaponPickup, DefaultAction)
	// GameEmoticon            func(*messages7.Emoticon, DefaultAction)
	// GameSvVoteClearoptions  func(*messages7.SvVoteClearoptions, DefaultAction)
	// GameSvVoteOptionlistadd func(*messages7.SvVoteOptionlistadd, DefaultAction)
	// GameSvVotePptionadd     func(*messages7.SvVotePptionadd, DefaultAction)
	// GameSvVoteOptionremove  func(*messages7.SvVoteOptionremove, DefaultAction)
	// GameSvVoteSet           func(*messages7.SvVoteSet, DefaultAction)
	// GameSvVoteStatus        func(*messages7.SvVoteStatus, DefaultAction)
	// GameSvServerSettings    func(*messages7.SvServerSettings, DefaultAction)
	GameSvClientInfo func(*messages7.SvClientInfo, DefaultAction)
	// GameSvGameInfo          func(*messages7.SvGameInfo, DefaultAction)
	// GameSvClientDrop        func(*messages7.SvClientDrop, DefaultAction)
	// GameSvGameMsg           func(*messages7.SvGameMsg, DefaultAction)
	// GameDeClientEnter       func(*messages7.DeClientEnter, DefaultAction)
	// GameDeClientLeave       func(*messages7.DeClientLeave, DefaultAction)
	// GameClSay               func(*messages7.ClSay, DefaultAction)
	// GameClSetTeam           func(*messages7.ClSetTeam, DefaultAction)
	// GameClSetSpectatorMode  func(*messages7.ClSetSpectatorMode, DefaultAction)
	GameClStartInfo func(*messages7.ClStartInfo, DefaultAction)
	// GameClKill              func(*messages7.ClKill, DefaultAction)
	// GameClReadyChange       func(*messages7.ClReadyChange, DefaultAction)
	// GameClEmoticon          func(*messages7.ClEmoticon, DefaultAction)
	// GameClVote              func(*messages7.ClVote, DefaultAction)
	// GameClCallVote          func(*messages7.ClCallVote, DefaultAction)
	// GameSvSkinChange        func(*messages7.SvSkinChange, DefaultAction)
	// GameClSkinChange        func(*messages7.ClSkinChange, DefaultAction)
	// GameSvRaceFinish        func(*messages7.SvRaceFinish, DefaultAction)
	// GameSvCheckpoint        func(*messages7.SvCheckpoint, DefaultAction)
	// GameSvCommandInfo       func(*messages7.SvCommandInfo, DefaultAction)
	// GameSvCommandInfoRemove func(*messages7.SvCommandInfoRemove, DefaultAction)
	// GameClCommand           func(*messages7.ClCommand, DefaultAction)
}

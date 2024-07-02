package teeworlds7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
	"github.com/teeworlds-go/go-teeworlds-protocol/snapshot7"
)

// Processes the incoming packet
// It might print to the console
// It might send a response packet
type DefaultAction func()

// Internal method to call user hooks and register default behavior for a given message
// Example:
//
// userMsgCallback(
//
//	client.Callbacks.GameSvMotd,
//	&messages7.SvMotd{},
//	func() { fmt.Println("default action") },
//
// )
func userMsgCallback[T any](userCallbacks []func(T, DefaultAction), msg T, defaultAction DefaultAction) {
	if len(userCallbacks) == 0 {
		defaultAction()
		return
	}

	for _, callback := range userCallbacks {
		callback(msg, defaultAction)
	}
}

// TODO: this should be a map but the type checker broke me
//
// // key is the network7.MessageId
// UserMsgCallbacks map[int]UserMsgCallback
type UserMsgCallbacks struct {
	// return false to drop the packet
	PacketIn []func(*protocol7.Packet) bool

	// return false to drop the packet
	PacketOut []func(*protocol7.Packet) bool

	// return false to drop the message
	MessageOut []func(*messages7.NetMessage) bool

	// ctrl out
	CtrlKeepAliveOut []func(*messages7.CtrlKeepAlive) bool
	CtrlConnectOut   []func(*messages7.CtrlConnect) bool
	CtrlAcceptOut    []func(*messages7.CtrlAccept) bool
	CtrlTokenOut     []func(*messages7.CtrlToken) bool
	CtrlCloseOut     []func(*messages7.CtrlClose) bool
	// sys out
	SysInfoOut            []func(*messages7.Info) bool
	SysMapChangeOut       []func(*messages7.MapChange) bool
	SysMapDataOut         []func(*messages7.MapData) bool
	SysServerInfoOut      []func(*messages7.ServerInfo) bool
	SysConReadyOut        []func(*messages7.ConReady) bool
	SysSnapOut            []func(*messages7.Snap) bool
	SysSnapEmptyOut       []func(*messages7.SnapEmpty) bool
	SysSnapSingleOut      []func(*messages7.SnapSingle) bool
	SysSnapSmallOut       []func(*messages7.SnapSmall) bool
	SysInputTimingOut     []func(*messages7.InputTiming) bool
	SysRconAuthOnOut      []func(*messages7.RconAuthOn) bool
	SysRconAuthOffOut     []func(*messages7.RconAuthOff) bool
	SysRconLineOut        []func(*messages7.RconLine) bool
	SysRconCmdAddOut      []func(*messages7.RconCmdAdd) bool
	SysRconCmdRemOut      []func(*messages7.RconCmdRem) bool
	SysAuthChallengeOut   []func(*messages7.AuthChallenge) bool
	SysAuthResultOut      []func(*messages7.AuthResult) bool
	SysReadyOut           []func(*messages7.Ready) bool
	SysEnterGameOut       []func(*messages7.EnterGame) bool
	SysInputOut           []func(*messages7.Input) bool
	SysRconCmdOut         []func(*messages7.RconCmd) bool
	SysRconAuthOut        []func(*messages7.RconAuth) bool
	SysRequestMapDataOut  []func(*messages7.RequestMapData) bool
	SysAuthStartOut       []func(*messages7.AuthStart) bool
	SysAuthResponseOut    []func(*messages7.AuthResponse) bool
	SysPingOut            []func(*messages7.Ping) bool
	SysPingReplyOut       []func(*messages7.PingReply) bool
	SysErrorOut           []func(*messages7.Error) bool
	SysMaplistEntryAddOut []func(*messages7.MaplistEntryAdd) bool
	SysMaplistEntryRemOut []func(*messages7.MaplistEntryRem) bool
	// game out
	GameSvMotdOut              []func(*messages7.SvMotd) bool
	GameSvBroadcastOut         []func(*messages7.SvBroadcast) bool
	GameSvChatOut              []func(*messages7.SvChat) bool
	GameSvTeamOut              []func(*messages7.SvTeam) bool
	GameSvKillMsgOut           []func(*messages7.SvKillMsg) bool
	GameSvTuneParamsOut        []func(*messages7.SvTuneParams) bool
	GameSvExtraProjectileOut   []func(*messages7.SvExtraProjectile) bool
	GameSvReadyToEnterOut      []func(*messages7.SvReadyToEnter) bool
	GameSvWeaponPickupOut      []func(*messages7.SvWeaponPickup) bool
	GameSvEmoticonOut          []func(*messages7.SvEmoticon) bool
	GameSvVoteClearOptionsOut  []func(*messages7.SvVoteClearOptions) bool
	GameSvVoteOptionListAddOut []func(*messages7.SvVoteOptionListAdd) bool
	GameSvVoteOptionAddOut     []func(*messages7.SvVoteOptionAdd) bool
	GameSvVoteOptionRemoveOut  []func(*messages7.SvVoteOptionRemove) bool
	GameSvVoteSetOut           []func(*messages7.SvVoteSet) bool
	GameSvVoteStatusOut        []func(*messages7.SvVoteStatus) bool
	GameSvServerSettingsOut    []func(*messages7.SvServerSettings) bool
	GameSvClientInfoOut        []func(*messages7.SvClientInfo) bool
	GameSvGameInfoOut          []func(*messages7.SvGameInfo) bool
	GameSvClientDropOut        []func(*messages7.SvClientDrop) bool
	GameSvGameMsgOut           []func(*messages7.SvGameMsg) bool
	GameDeClientEnterOut       []func(*messages7.DeClientEnter) bool
	GameDeClientLeaveOut       []func(*messages7.DeClientLeave) bool
	GameClSayOut               []func(*messages7.ClSay) bool
	GameClSetTeamOut           []func(*messages7.ClSetTeam) bool
	GameClSetSpectatorModeOut  []func(*messages7.ClSetSpectatorMode) bool
	GameClStartInfoOut         []func(*messages7.ClStartInfo) bool
	GameClKillOut              []func(*messages7.ClKill) bool
	GameClReadyChangeOut       []func(*messages7.ClReadyChange) bool
	GameClEmoticonOut          []func(*messages7.ClEmoticon) bool
	GameClVoteOut              []func(*messages7.ClVote) bool
	GameClCallVoteOut          []func(*messages7.ClCallVote) bool
	GameSvSkinChangeOut        []func(*messages7.SvSkinChange) bool
	GameClSkinChangeOut        []func(*messages7.ClSkinChange) bool
	GameSvRaceFinishOut        []func(*messages7.SvRaceFinish) bool
	GameSvCheckpointOut        []func(*messages7.SvCheckpoint) bool
	GameSvCommandInfoOut       []func(*messages7.SvCommandInfo) bool
	GameSvCommandInfoRemoveOut []func(*messages7.SvCommandInfoRemove) bool
	GameClCommandOut           []func(*messages7.ClCommand) bool

	// return false to drop the error (ignore it)
	//
	// return true to pass the error on and finally throw
	InternalError []func(error) bool
	MsgUnknown    []func(*messages7.Unknown, DefaultAction)
	Snapshot      []func(*snapshot7.Snapshot, DefaultAction)

	CtrlKeepAlive []func(*messages7.CtrlKeepAlive, DefaultAction)
	CtrlConnect   []func(*messages7.CtrlConnect, DefaultAction)
	CtrlAccept    []func(*messages7.CtrlAccept, DefaultAction)
	CtrlToken     []func(*messages7.CtrlToken, DefaultAction)
	CtrlClose     []func(*messages7.CtrlClose, DefaultAction)

	SysInfo            []func(*messages7.Info, DefaultAction)
	SysMapChange       []func(*messages7.MapChange, DefaultAction)
	SysMapData         []func(*messages7.MapData, DefaultAction)
	SysServerInfo      []func(*messages7.ServerInfo, DefaultAction)
	SysConReady        []func(*messages7.ConReady, DefaultAction)
	SysSnap            []func(*messages7.Snap, DefaultAction)
	SysSnapEmpty       []func(*messages7.SnapEmpty, DefaultAction)
	SysSnapSingle      []func(*messages7.SnapSingle, DefaultAction)
	SysSnapSmall       []func(*messages7.SnapSmall, DefaultAction)
	SysInputTiming     []func(*messages7.InputTiming, DefaultAction)
	SysRconAuthOn      []func(*messages7.RconAuthOn, DefaultAction)
	SysRconAuthOff     []func(*messages7.RconAuthOff, DefaultAction)
	SysRconLine        []func(*messages7.RconLine, DefaultAction)
	SysRconCmdAdd      []func(*messages7.RconCmdAdd, DefaultAction)
	SysRconCmdRem      []func(*messages7.RconCmdRem, DefaultAction)
	SysAuthChallenge   []func(*messages7.AuthChallenge, DefaultAction)
	SysAuthResult      []func(*messages7.AuthResult, DefaultAction)
	SysReady           []func(*messages7.Ready, DefaultAction)
	SysEnterGame       []func(*messages7.EnterGame, DefaultAction)
	SysInput           []func(*messages7.Input, DefaultAction)
	SysRconCmd         []func(*messages7.RconCmd, DefaultAction)
	SysRconAuth        []func(*messages7.RconAuth, DefaultAction)
	SysRequestMapData  []func(*messages7.RequestMapData, DefaultAction)
	SysAuthStart       []func(*messages7.AuthStart, DefaultAction)
	SysAuthResponse    []func(*messages7.AuthResponse, DefaultAction)
	SysPing            []func(*messages7.Ping, DefaultAction)
	SysPingReply       []func(*messages7.PingReply, DefaultAction)
	SysError           []func(*messages7.Error, DefaultAction)
	SysMaplistEntryAdd []func(*messages7.MaplistEntryAdd, DefaultAction)
	SysMaplistEntryRem []func(*messages7.MaplistEntryRem, DefaultAction)

	GameSvMotd              []func(*messages7.SvMotd, DefaultAction)
	GameSvBroadcast         []func(*messages7.SvBroadcast, DefaultAction)
	GameSvChat              []func(*messages7.SvChat, DefaultAction)
	GameSvTeam              []func(*messages7.SvTeam, DefaultAction)
	GameSvKillMsg           []func(*messages7.SvKillMsg, DefaultAction)
	GameSvTuneParams        []func(*messages7.SvTuneParams, DefaultAction)
	GameSvExtraProjectile   []func(*messages7.SvExtraProjectile, DefaultAction)
	GameSvReadyToEnter      []func(*messages7.SvReadyToEnter, DefaultAction)
	GameSvWeaponPickup      []func(*messages7.SvWeaponPickup, DefaultAction)
	GameSvEmoticon          []func(*messages7.SvEmoticon, DefaultAction)
	GameSvVoteClearOptions  []func(*messages7.SvVoteClearOptions, DefaultAction)
	GameSvVoteOptionListAdd []func(*messages7.SvVoteOptionListAdd, DefaultAction)
	GameSvVoteOptionAdd     []func(*messages7.SvVoteOptionAdd, DefaultAction)
	GameSvVoteOptionRemove  []func(*messages7.SvVoteOptionRemove, DefaultAction)
	GameSvVoteSet           []func(*messages7.SvVoteSet, DefaultAction)
	GameSvVoteStatus        []func(*messages7.SvVoteStatus, DefaultAction)
	GameSvServerSettings    []func(*messages7.SvServerSettings, DefaultAction)
	GameSvClientInfo        []func(*messages7.SvClientInfo, DefaultAction)
	GameSvGameInfo          []func(*messages7.SvGameInfo, DefaultAction)
	GameSvClientDrop        []func(*messages7.SvClientDrop, DefaultAction)
	GameSvGameMsg           []func(*messages7.SvGameMsg, DefaultAction)
	GameDeClientEnter       []func(*messages7.DeClientEnter, DefaultAction)
	GameDeClientLeave       []func(*messages7.DeClientLeave, DefaultAction)
	GameClSay               []func(*messages7.ClSay, DefaultAction)
	GameClSetTeam           []func(*messages7.ClSetTeam, DefaultAction)
	GameClSetSpectatorMode  []func(*messages7.ClSetSpectatorMode, DefaultAction)
	GameClStartInfo         []func(*messages7.ClStartInfo, DefaultAction)
	GameClKill              []func(*messages7.ClKill, DefaultAction)
	GameClReadyChange       []func(*messages7.ClReadyChange, DefaultAction)
	GameClEmoticon          []func(*messages7.ClEmoticon, DefaultAction)
	GameClVote              []func(*messages7.ClVote, DefaultAction)
	GameClCallVote          []func(*messages7.ClCallVote, DefaultAction)
	GameSvSkinChange        []func(*messages7.SvSkinChange, DefaultAction)
	GameClSkinChange        []func(*messages7.ClSkinChange, DefaultAction)
	GameSvRaceFinish        []func(*messages7.SvRaceFinish, DefaultAction)
	GameSvCheckpoint        []func(*messages7.SvCheckpoint, DefaultAction)
	GameSvCommandInfo       []func(*messages7.SvCommandInfo, DefaultAction)
	GameSvCommandInfoRemove []func(*messages7.SvCommandInfoRemove, DefaultAction)
	GameClCommand           []func(*messages7.ClCommand, DefaultAction)
}

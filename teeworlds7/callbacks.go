package teeworlds7

import (
	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

// Processes the incoming packet
// It might print to the console
// It might send a response packet
type DefaultAction func() error

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
func userMsgCallback[T any](userCallbacks []func(T, DefaultAction) error, msg T, defaultAction DefaultAction) error {
	if len(userCallbacks) == 0 {
		if defaultAction != nil {
			return defaultAction()

		}
		return nil
	}

	var err error
	for _, callback := range userCallbacks {
		err = callback(msg, defaultAction)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: this should be a map but the type checker broke me
//
// // key is the network7.MessageId
// UserMsgCallbacks map[int]UserMsgCallback
type UserMsgCallbacks struct {
	Tick []func(DefaultAction) error

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
	InternalError []func(error) (bool, error)
	MsgUnknown    []func(*messages7.Unknown, DefaultAction) error
	Snapshot      []func(*snapshot7.Snapshot, DefaultAction) error

	CtrlKeepAlive []func(*messages7.CtrlKeepAlive, DefaultAction) error
	CtrlConnect   []func(*messages7.CtrlConnect, DefaultAction) error
	CtrlAccept    []func(*messages7.CtrlAccept, DefaultAction) error
	CtrlToken     []func(*messages7.CtrlToken, DefaultAction) error
	CtrlClose     []func(*messages7.CtrlClose, DefaultAction) error

	SysInfo            []func(*messages7.Info, DefaultAction) error
	SysMapChange       []func(*messages7.MapChange, DefaultAction) error
	SysMapData         []func(*messages7.MapData, DefaultAction) error
	SysServerInfo      []func(*messages7.ServerInfo, DefaultAction) error
	SysConReady        []func(*messages7.ConReady, DefaultAction) error
	SysSnap            []func(*messages7.Snap, DefaultAction) error
	SysSnapEmpty       []func(*messages7.SnapEmpty, DefaultAction) error
	SysSnapSingle      []func(*messages7.SnapSingle, DefaultAction) error
	SysSnapSmall       []func(*messages7.SnapSmall, DefaultAction) error
	SysInputTiming     []func(*messages7.InputTiming, DefaultAction) error
	SysRconAuthOn      []func(*messages7.RconAuthOn, DefaultAction) error
	SysRconAuthOff     []func(*messages7.RconAuthOff, DefaultAction) error
	SysRconLine        []func(*messages7.RconLine, DefaultAction) error
	SysRconCmdAdd      []func(*messages7.RconCmdAdd, DefaultAction) error
	SysRconCmdRem      []func(*messages7.RconCmdRem, DefaultAction) error
	SysAuthChallenge   []func(*messages7.AuthChallenge, DefaultAction) error
	SysAuthResult      []func(*messages7.AuthResult, DefaultAction) error
	SysReady           []func(*messages7.Ready, DefaultAction) error
	SysEnterGame       []func(*messages7.EnterGame, DefaultAction) error
	SysInput           []func(*messages7.Input, DefaultAction) error
	SysRconCmd         []func(*messages7.RconCmd, DefaultAction) error
	SysRconAuth        []func(*messages7.RconAuth, DefaultAction) error
	SysRequestMapData  []func(*messages7.RequestMapData, DefaultAction) error
	SysAuthStart       []func(*messages7.AuthStart, DefaultAction) error
	SysAuthResponse    []func(*messages7.AuthResponse, DefaultAction) error
	SysPing            []func(*messages7.Ping, DefaultAction) error
	SysPingReply       []func(*messages7.PingReply, DefaultAction) error
	SysError           []func(*messages7.Error, DefaultAction) error
	SysMaplistEntryAdd []func(*messages7.MaplistEntryAdd, DefaultAction) error
	SysMaplistEntryRem []func(*messages7.MaplistEntryRem, DefaultAction) error

	GameSvMotd              []func(*messages7.SvMotd, DefaultAction) error
	GameSvBroadcast         []func(*messages7.SvBroadcast, DefaultAction) error
	GameSvChat              []func(*messages7.SvChat, DefaultAction) error
	GameSvTeam              []func(*messages7.SvTeam, DefaultAction) error
	GameSvKillMsg           []func(*messages7.SvKillMsg, DefaultAction) error
	GameSvTuneParams        []func(*messages7.SvTuneParams, DefaultAction) error
	GameSvExtraProjectile   []func(*messages7.SvExtraProjectile, DefaultAction) error
	GameSvReadyToEnter      []func(*messages7.SvReadyToEnter, DefaultAction) error
	GameSvWeaponPickup      []func(*messages7.SvWeaponPickup, DefaultAction) error
	GameSvEmoticon          []func(*messages7.SvEmoticon, DefaultAction) error
	GameSvVoteClearOptions  []func(*messages7.SvVoteClearOptions, DefaultAction) error
	GameSvVoteOptionListAdd []func(*messages7.SvVoteOptionListAdd, DefaultAction) error
	GameSvVoteOptionAdd     []func(*messages7.SvVoteOptionAdd, DefaultAction) error
	GameSvVoteOptionRemove  []func(*messages7.SvVoteOptionRemove, DefaultAction) error
	GameSvVoteSet           []func(*messages7.SvVoteSet, DefaultAction) error
	GameSvVoteStatus        []func(*messages7.SvVoteStatus, DefaultAction) error
	GameSvServerSettings    []func(*messages7.SvServerSettings, DefaultAction) error
	GameSvClientInfo        []func(*messages7.SvClientInfo, DefaultAction) error
	GameSvGameInfo          []func(*messages7.SvGameInfo, DefaultAction) error
	GameSvClientDrop        []func(*messages7.SvClientDrop, DefaultAction) error
	GameSvGameMsg           []func(*messages7.SvGameMsg, DefaultAction) error
	GameDeClientEnter       []func(*messages7.DeClientEnter, DefaultAction) error
	GameDeClientLeave       []func(*messages7.DeClientLeave, DefaultAction) error
	GameClSay               []func(*messages7.ClSay, DefaultAction) error
	GameClSetTeam           []func(*messages7.ClSetTeam, DefaultAction) error
	GameClSetSpectatorMode  []func(*messages7.ClSetSpectatorMode, DefaultAction) error
	GameClStartInfo         []func(*messages7.ClStartInfo, DefaultAction) error
	GameClKill              []func(*messages7.ClKill, DefaultAction) error
	GameClReadyChange       []func(*messages7.ClReadyChange, DefaultAction) error
	GameClEmoticon          []func(*messages7.ClEmoticon, DefaultAction) error
	GameClVote              []func(*messages7.ClVote, DefaultAction) error
	GameClCallVote          []func(*messages7.ClCallVote, DefaultAction) error
	GameSvSkinChange        []func(*messages7.SvSkinChange, DefaultAction) error
	GameClSkinChange        []func(*messages7.ClSkinChange, DefaultAction) error
	GameSvRaceFinish        []func(*messages7.SvRaceFinish, DefaultAction) error
	GameSvCheckpoint        []func(*messages7.SvCheckpoint, DefaultAction) error
	GameSvCommandInfo       []func(*messages7.SvCommandInfo, DefaultAction) error
	GameSvCommandInfoRemove []func(*messages7.SvCommandInfoRemove, DefaultAction) error
	GameClCommand           []func(*messages7.ClCommand, DefaultAction) error
}

package teeworlds7

import (
	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

// --------------------------------
// special cases
// --------------------------------

func (client *Client) OnTick(callback func(defaultAction DefaultAction) error) {
	client.Callbacks.Tick = append(client.Callbacks.Tick, callback)
}

// if not implemented by the user the application might throw and exit
//
// return nil to drop the error
// return the error or another error
// in order to shutdown the application.
func (client *Client) OnError(callback func(err error) error) {
	client.Callbacks.InternalError = append(client.Callbacks.InternalError, callback)
}

// inspect outgoing traffic
// and alter it before it gets sent to the server
//
// return false to drop the packet
func (client *Client) OnSendPacket(callback func(packet *protocol7.Packet) bool) {
	client.Callbacks.PacketOut = append(client.Callbacks.PacketOut, callback)
}

// read incoming traffic
// and alter it before it hits the internal state machine
//
// return false to drop the packet
func (client *Client) OnPacket(callback func(packet *protocol7.Packet) bool) {
	client.Callbacks.PacketIn = append(client.Callbacks.PacketIn, callback)
}

func (client *Client) OnUnknown(callback func(msg *messages7.Unknown, defaultAction DefaultAction) error) {
	client.Callbacks.MsgUnknown = append(client.Callbacks.MsgUnknown, callback)
}

// will be called when a snap, snap single or empty snapshot is received
// if you want to know which type of snapshot was received look at OnMsgSnap(), OnMsgSnapEmpty(), OnMsgSnapSingle(), OnMsgSnapSmall()
func (client *Client) OnSnapshot(callback func(snap *snapshot7.Snapshot, defaultAction DefaultAction) error) {
	client.Callbacks.Snapshot = append(client.Callbacks.Snapshot, callback)
}

// --------------------------------
// incoming control messages
// --------------------------------

func (client *Client) OnKeepAlive(callback func(msg *messages7.CtrlKeepAlive, defaultAction DefaultAction) error) {
	client.Callbacks.CtrlKeepAlive = append(client.Callbacks.CtrlKeepAlive, callback)
}

// This is just misleading. It should never be called. This message is only received by the server.
// func (client *Client) OnCtrlConnect(callback func(msg *messages7.CtrlConnect, defaultAction DefaultAction) error) {
// 	client.Callbacks.CtrlConnect = append(// 	client.Callbacks, callback)
// }

func (client *Client) OnAccept(callback func(msg *messages7.CtrlAccept, defaultAction DefaultAction) error) {
	client.Callbacks.CtrlAccept = append(client.Callbacks.CtrlAccept, callback)
}

func (client *Client) OnDisconnect(callback func(msg *messages7.CtrlClose, defaultAction DefaultAction) error) {
	client.Callbacks.CtrlClose = append(client.Callbacks.CtrlClose, callback)
}

func (client *Client) OnToken(callback func(msg *messages7.CtrlToken, defaultAction DefaultAction) error) {
	client.Callbacks.CtrlToken = append(client.Callbacks.CtrlToken, callback)
}

// --------------------------------
// incoming game messages
// --------------------------------

func (client *Client) OnMotd(callback func(msg *messages7.SvMotd, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvMotd = append(client.Callbacks.GameSvMotd, callback)
}

func (client *Client) OnBroadcast(callback func(msg *messages7.SvBroadcast, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvBroadcast = append(client.Callbacks.GameSvBroadcast, callback)
}

func (client *Client) OnChat(callback func(msg *messages7.SvChat, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvChat = append(client.Callbacks.GameSvChat, callback)
}

func (client *Client) OnTeam(callback func(msg *messages7.SvTeam, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvTeam = append(client.Callbacks.GameSvTeam, callback)
}

func (client *Client) OnKillMsg(callback func(msg *messages7.SvKillMsg, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvKillMsg = append(client.Callbacks.GameSvKillMsg, callback)
}

func (client *Client) OnTuneParams(callback func(msg *messages7.SvTuneParams, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvTuneParams = append(client.Callbacks.GameSvTuneParams, callback)
}

func (client *Client) OnExtraProjectile(callback func(msg *messages7.SvExtraProjectile, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvExtraProjectile = append(client.Callbacks.GameSvExtraProjectile, callback)
}

func (client *Client) OnReadyToEnter(callback func(msg *messages7.SvReadyToEnter, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvReadyToEnter = append(client.Callbacks.GameSvReadyToEnter, callback)
}

func (client *Client) OnWeaponPickup(callback func(msg *messages7.SvWeaponPickup, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvWeaponPickup = append(client.Callbacks.GameSvWeaponPickup, callback)
}

func (client *Client) OnEmoticon(callback func(msg *messages7.SvEmoticon, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvEmoticon = append(client.Callbacks.GameSvEmoticon, callback)
}

func (client *Client) OnVoteClearoptions(callback func(msg *messages7.SvVoteClearOptions, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvVoteClearOptions = append(client.Callbacks.GameSvVoteClearOptions, callback)
}

func (client *Client) OnVoteOptionlistadd(callback func(msg *messages7.SvVoteOptionListAdd, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvVoteOptionListAdd = append(client.Callbacks.GameSvVoteOptionListAdd, callback)
}

func (client *Client) OnVotePptionadd(callback func(msg *messages7.SvVoteOptionAdd, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvVoteOptionAdd = append(client.Callbacks.GameSvVoteOptionAdd, callback)
}

func (client *Client) OnVoteOptionremove(callback func(msg *messages7.SvVoteOptionRemove, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvVoteOptionRemove = append(client.Callbacks.GameSvVoteOptionRemove, callback)
}

func (client *Client) OnVoteSet(callback func(msg *messages7.SvVoteSet, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvVoteSet = append(client.Callbacks.GameSvVoteSet, callback)
}

func (client *Client) OnVoteStatus(callback func(msg *messages7.SvVoteStatus, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvVoteStatus = append(client.Callbacks.GameSvVoteStatus, callback)
}

func (client *Client) OnServerSettings(callback func(msg *messages7.SvServerSettings, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvServerSettings = append(client.Callbacks.GameSvServerSettings, callback)
}

func (client *Client) OnClientInfo(callback func(msg *messages7.SvClientInfo, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvClientInfo = append(client.Callbacks.GameSvClientInfo, callback)
}

func (client *Client) OnGameInfo(callback func(msg *messages7.SvGameInfo, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvGameInfo = append(client.Callbacks.GameSvGameInfo, callback)
}

func (client *Client) OnClientDrop(callback func(msg *messages7.SvClientDrop, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvClientDrop = append(client.Callbacks.GameSvClientDrop, callback)
}

func (client *Client) OnGameMsg(callback func(msg *messages7.SvGameMsg, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvGameMsg = append(client.Callbacks.GameSvGameMsg, callback)
}

// demo only
// func (client *Client) OnClientEnter(callback func(msg *messages7.DeClientEnter, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameDeClientEnter = append(client.Callbacks.GameDeClientEnter, callback)
// }

// demo only
// func (client *Client) OnClientLeave(callback func(msg *messages7.DeClientLeave, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameDeClientLeave = append(client.Callbacks.GameDeClientLeave, callback)
// }

// send by client
// func (client *Client) OnSay(callback func(msg *messages7.ClSay, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClSay = append(client.Callbacks.GameClSay, callback)
// }

// send by client
// func (client *Client) OnSetTeam(callback func(msg *messages7.ClSetTeam, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClSetTeam = append(client.Callbacks.GameClSetTeam, callback)
// }

// send by client
// func (client *Client) OnSetSpectatorMode(callback func(msg *messages7.ClSetSpectatorMode, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClSetSpectatorMode = append(client.Callbacks.GameClSetSpectatorMode, callback)
// }

// send by client
// func (client *Client) OnStartInfo(callback func(msg *messages7.ClStartInfo, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClStartInfo = append(client.Callbacks.GameClStartInfo, callback)
// }

// send by client
// func (client *Client) OnKill(callback func(msg *messages7.ClKill, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClKill = append(client.Callbacks.GameClKill, callback)
// }

// send by client
// func (client *Client) OnReadyChange(callback func(msg *messages7.ClReadyChange, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClReadyChange = append(client.Callbacks.GameClReadyChange, callback)
// }

// send by client
// func (client *Client) OnEmoticon(callback func(msg *messages7.ClEmoticon, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClEmoticon = append(client.Callbacks.GameClEmoticon, callback)
// }

// send by client
// func (client *Client) OnVote(callback func(msg *messages7.ClVote, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClVote = append(client.Callbacks.GameClVote, callback)
// }

// send by client
// func (client *Client) OnCallVote(callback func(msg *messages7.ClCallVote, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClCallVote = append(client.Callbacks.GameClCallVote, callback)
// }

func (client *Client) OnSkinChange(callback func(msg *messages7.SvSkinChange, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvSkinChange = append(client.Callbacks.GameSvSkinChange, callback)
}

// send by client
// func (client *Client) OnSkinChange(callback func(msg *messages7.ClSkinChange, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClSkinChange = append(client.Callbacks.GameClSkinChange, callback)
// }

func (client *Client) OnRaceFinish(callback func(msg *messages7.SvRaceFinish, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvRaceFinish = append(client.Callbacks.GameSvRaceFinish, callback)
}

func (client *Client) OnCheckpoint(callback func(msg *messages7.SvCheckpoint, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvCheckpoint = append(client.Callbacks.GameSvCheckpoint, callback)
}

func (client *Client) OnCommandInfo(callback func(msg *messages7.SvCommandInfo, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvCommandInfo = append(client.Callbacks.GameSvCommandInfo, callback)
}

func (client *Client) OnCommandInfoRemove(callback func(msg *messages7.SvCommandInfoRemove, defaultAction DefaultAction) error) {
	client.Callbacks.GameSvCommandInfoRemove = append(client.Callbacks.GameSvCommandInfoRemove, callback)
}

// send by client
// func (client *Client) OnCommand(callback func(msg *messages7.ClCommand, defaultAction DefaultAction) error) {
// 	client.Callbacks.GameClCommand = append(client.Callbacks.GameClCommand, callback)
// }

// --------------------------------
// incoming system messages
// --------------------------------

func (client *Client) OnMapChange(callback func(msg *messages7.MapChange, defaultAction DefaultAction) error) {
	client.Callbacks.SysMapChange = append(client.Callbacks.SysMapChange, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnap(callback func(msg *messages7.Snap, defaultAction DefaultAction) error) {
	client.Callbacks.SysSnap = append(client.Callbacks.SysSnap, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnapEmpty(callback func(msg *messages7.SnapEmpty, defaultAction DefaultAction) error) {
	client.Callbacks.SysSnapEmpty = append(client.Callbacks.SysSnapEmpty, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnapSingle(callback func(msg *messages7.SnapSingle, defaultAction DefaultAction) error) {
	client.Callbacks.SysSnapSingle = append(client.Callbacks.SysSnapSingle, callback)
}

// You probably want to use OnSnapshot() instead
func (client *Client) OnMsgSnapSmall(callback func(msg *messages7.SnapSmall, defaultAction DefaultAction) error) {
	client.Callbacks.SysSnapSmall = append(client.Callbacks.SysSnapSmall, callback)
}

func (client *Client) OnServerInfo(callback func(msg *messages7.ServerInfo, defaultAction DefaultAction) error) {
	client.Callbacks.SysServerInfo = append(client.Callbacks.SysServerInfo, callback)
}

// --------------------------------
// outgoing system messages
// --------------------------------

func (client *Client) OnSendInfo(callback func(msg *messages7.Info) bool) {
	client.Callbacks.SysInfoOut = append(client.Callbacks.SysInfoOut, callback)
}

func (client *Client) OnSendConReady(callback func(msg *messages7.ConReady) bool) {
	client.Callbacks.SysConReadyOut = append(client.Callbacks.SysConReadyOut, callback)
}

func (client *Client) OnSendInputTiming(callback func(msg *messages7.InputTiming) bool) {
	client.Callbacks.SysInputTimingOut = append(client.Callbacks.SysInputTimingOut, callback)
}

func (client *Client) OnSendReady(callback func(msg *messages7.Ready) bool) {
	client.Callbacks.SysReadyOut = append(client.Callbacks.SysReadyOut, callback)
}

func (client *Client) OnSendEnterGame(callback func(msg *messages7.EnterGame) bool) {
	client.Callbacks.SysEnterGameOut = append(client.Callbacks.SysEnterGameOut, callback)
}

func (client *Client) OnSendInput(callback func(msg *messages7.Input) bool) {
	client.Callbacks.SysInputOut = append(client.Callbacks.SysInputOut, callback)
}

func (client *Client) OnSendRconCmd(callback func(msg *messages7.RconCmd) bool) {
	client.Callbacks.SysRconCmdOut = append(client.Callbacks.SysRconCmdOut, callback)
}

func (client *Client) OnSendRconAuth(callback func(msg *messages7.RconAuth) bool) {
	client.Callbacks.SysRconAuthOut = append(client.Callbacks.SysRconAuthOut, callback)
}

func (client *Client) OnSendRequestMapData(callback func(msg *messages7.RequestMapData) bool) {
	client.Callbacks.SysRequestMapDataOut = append(client.Callbacks.SysRequestMapDataOut, callback)
}

func (client *Client) OnSendPing(callback func(msg *messages7.Ping) bool) {
	client.Callbacks.SysPingOut = append(client.Callbacks.SysPingOut, callback)
}

func (client *Client) OnSendPingReply(callback func(msg *messages7.PingReply) bool) {
	client.Callbacks.SysPingReplyOut = append(client.Callbacks.SysPingReplyOut, callback)
}

// --------------------------------
// outgoing game messages
// --------------------------------

// send by server
// func (client *Client) OnSendMotd(callback func(msg *messages7.SvMotd) bool) {
// 	client.Callbacks.GameSvMotdOut = append(client.Callbacks.GameSvMotdOut, callback)
// }

// send by server
// func (client *Client) OnSendBroadcast(callback func(msg *messages7.SvBroadcast) bool) {
// 	client.Callbacks.GameSvBroadcastOut = append(client.Callbacks.GameSvBroadcastOut, callback)
// }

// send by server
// func (client *Client) OnSendChat(callback func(msg *messages7.SvChat) bool) {
// 	client.Callbacks.GameSvChatOut = append(client.Callbacks.GameSvChatOut, callback)
// }

// send by server
// func (client *Client) OnSendTeam(callback func(msg *messages7.SvTeam) bool) {
// 	client.Callbacks.GameSvTeamOut = append(client.Callbacks.GameSvTeamOut, callback)
// }

// send by server
// func (client *Client) OnSendKillMsg(callback func(msg *messages7.SvKillMsg) bool) {
// 	client.Callbacks.GameSvKillMsgOut = append(client.Callbacks.GameSvKillMsgOut, callback)
// }

// send by server
// func (client *Client) OnSendTuneParams(callback func(msg *messages7.SvTuneParams) bool) {
// 	client.Callbacks.GameSvTuneParamsOut = append(client.Callbacks.GameSvTuneParamsOut, callback)
// }

// send by server
// func (client *Client) OnSendExtraProjectile(callback func(msg *messages7.SvExtraProjectile) bool) {
// 	client.Callbacks.GameSvExtraProjectileOut = append(client.Callbacks.GameSvExtraProjectileOut, callback)
// }

// send by server
// func (client *Client) OnSendReadyToEnter(callback func(msg *messages7.SvReadyToEnter) bool) {
// 	client.Callbacks.GameSvReadyToEnterOut = append(client.Callbacks.GameSvReadyToEnterOut, callback)
// }

// send by server
// func (client *Client) OnSendWeaponPickup(callback func(msg *messages7.SvWeaponPickup) bool) {
// 	client.Callbacks.GameSvWeaponPickupOut = append(client.Callbacks.GameSvWeaponPickupOut, callback)
// }

// send by server
// func (client *Client) OnSendEmoticon(callback func(msg *messages7.SvEmoticon) bool) {
// 	client.Callbacks.GameSvEmoticonOut = append(client.Callbacks.GameSvEmoticonOut, callback)
// }

// send by server
// func (client *Client) OnSendVoteClearOptions(callback func(msg *messages7.SvVoteClearOptions) bool) {
// 	client.Callbacks.GameSvVoteClearOptionsOut = append(client.Callbacks.GameSvVoteClearOptionsOut, callback)
// }

// send by server
// func (client *Client) OnSendVoteOptionListAdd(callback func(msg *messages7.SvVoteOptionListAdd) bool) {
// 	client.Callbacks.GameSvVoteOptionListAddOut = append(client.Callbacks.GameSvVoteOptionListAddOut, callback)
// }

// send by server
// func (client *Client) OnSendVoteOptionAdd(callback func(msg *messages7.SvVoteOptionAdd) bool) {
// 	client.Callbacks.GameSvVoteOptionAddOut = append(client.Callbacks.GameSvVoteOptionAddOut, callback)
// }

// send by server
// func (client *Client) OnSendVoteOptionRemove(callback func(msg *messages7.SvVoteOptionRemove) bool) {
// 	client.Callbacks.GameSvVoteOptionRemoveOut = append(client.Callbacks.GameSvVoteOptionRemoveOut, callback)
// }

// send by server
// func (client *Client) OnSendVoteSet(callback func(msg *messages7.SvVoteSet) bool) {
// 	client.Callbacks.GameSvVoteSetOut = append(client.Callbacks.GameSvVoteSetOut, callback)
// }

// send by server
// func (client *Client) OnSendVoteStatus(callback func(msg *messages7.SvVoteStatus) bool) {
// 	client.Callbacks.GameSvVoteStatusOut = append(client.Callbacks.GameSvVoteStatusOut, callback)
// }

// send by server
// func (client *Client) OnSendServerSettings(callback func(msg *messages7.SvServerSettings) bool) {
// 	client.Callbacks.GameSvServerSettingsOut = append(client.Callbacks.GameSvServerSettingsOut, callback)
// }

// send by server
// func (client *Client) OnSendClientInfo(callback func(msg *messages7.SvClientInfo) bool) {
// 	client.Callbacks.GameSvClientInfoOut = append(client.Callbacks.GameSvClientInfoOut, callback)
// }

// send by server
// func (client *Client) OnSendGameInfo(callback func(msg *messages7.SvGameInfo) bool) {
// 	client.Callbacks.GameSvGameInfoOut = append(client.Callbacks.GameSvGameInfoOut, callback)
// }

// send by server
// func (client *Client) OnSendClientDrop(callback func(msg *messages7.SvClientDrop) bool) {
// 	client.Callbacks.GameSvClientDropOut = append(client.Callbacks.GameSvClientDropOut, callback)
// }

// send by server
// func (client *Client) OnSendGameMsg(callback func(msg *messages7.SvGameMsg) bool) {
// 	client.Callbacks.GameSvGameMsgOut = append(client.Callbacks.GameSvGameMsgOut, callback)
// }

// demo only
// func (client *Client) OnSendClientEnter(callback func(msg *messages7.DeClientEnter) bool) {
// 	client.Callbacks.GameDeClientEnterOut = append(client.Callbacks.GameDeClientEnterOut, callback)
// }

// demo only
// func (client *Client) OnSendClientLeave(callback func(msg *messages7.DeClientLeave) bool) {
// 	client.Callbacks.GameDeClientLeaveOut = append(client.Callbacks.GameDeClientLeaveOut, callback)
// }

func (client *Client) OnSendSay(callback func(msg *messages7.ClSay) bool) {
	client.Callbacks.GameClSayOut = append(client.Callbacks.GameClSayOut, callback)
}

func (client *Client) OnSendSetTeam(callback func(msg *messages7.ClSetTeam) bool) {
	client.Callbacks.GameClSetTeamOut = append(client.Callbacks.GameClSetTeamOut, callback)
}

func (client *Client) OnSendSetSpectatorMode(callback func(msg *messages7.ClSetSpectatorMode) bool) {
	client.Callbacks.GameClSetSpectatorModeOut = append(client.Callbacks.GameClSetSpectatorModeOut, callback)
}

func (client *Client) OnSendStartInfo(callback func(msg *messages7.ClStartInfo) bool) {
	client.Callbacks.GameClStartInfoOut = append(client.Callbacks.GameClStartInfoOut, callback)
}

func (client *Client) OnSendKill(callback func(msg *messages7.ClKill) bool) {
	client.Callbacks.GameClKillOut = append(client.Callbacks.GameClKillOut, callback)
}

func (client *Client) OnSendReadyChange(callback func(msg *messages7.ClReadyChange) bool) {
	client.Callbacks.GameClReadyChangeOut = append(client.Callbacks.GameClReadyChangeOut, callback)
}

func (client *Client) OnSendEmoticon(callback func(msg *messages7.ClEmoticon) bool) {
	client.Callbacks.GameClEmoticonOut = append(client.Callbacks.GameClEmoticonOut, callback)
}

func (client *Client) OnSendVote(callback func(msg *messages7.ClVote) bool) {
	client.Callbacks.GameClVoteOut = append(client.Callbacks.GameClVoteOut, callback)
}

func (client *Client) OnSendCallVote(callback func(msg *messages7.ClCallVote) bool) {
	client.Callbacks.GameClCallVoteOut = append(client.Callbacks.GameClCallVoteOut, callback)
}

// send by server
// func (client *Client) OnSendSkinChange(callback func(msg *messages7.SvSkinChange) bool) {
// 	client.Callbacks.GameSvSkinChangeOut = append(client.Callbacks.GameSvSkinChangeOut, callback)
// }

func (client *Client) OnSendSkinChange(callback func(msg *messages7.ClSkinChange) bool) {
	client.Callbacks.GameClSkinChangeOut = append(client.Callbacks.GameClSkinChangeOut, callback)
}

// send by server
// func (client *Client) OnSendRaceFinish(callback func(msg *messages7.SvRaceFinish) bool) {
// 	client.Callbacks.GameSvRaceFinishOut = append(client.Callbacks.GameSvRaceFinishOut, callback)
// }

// send by server
// func (client *Client) OnSendCheckpoint(callback func(msg *messages7.SvCheckpoint) bool) {
// 	client.Callbacks.GameSvCheckpointOut = append(client.Callbacks.GameSvCheckpointOut, callback)
// }

// send by server
// func (client *Client) OnSendCommandInfo(callback func(msg *messages7.SvCommandInfo) bool) {
// 	client.Callbacks.GameSvCommandInfoOut = append(client.Callbacks.GameSvCommandInfoOut, callback)
// }

// send by server
// func (client *Client) OnSendCommandInfoRemove(callback func(msg *messages7.SvCommandInfoRemove) bool) {
// 	client.Callbacks.GameSvCommandInfoRemoveOut = append(client.Callbacks.GameSvCommandInfoRemoveOut, callback)
// }

func (client *Client) OnSendCommand(callback func(msg *messages7.ClCommand) bool) {
	client.Callbacks.GameClCommandOut = append(client.Callbacks.GameClCommandOut, callback)
}

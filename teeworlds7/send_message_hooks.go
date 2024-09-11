package teeworlds7

import (
	"fmt"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
)

func userSendMsgCallback[T any](userCallbacks []func(T) bool, msg T) bool {
	for _, callback := range userCallbacks {
		if !callback(msg) {
			return false
		}
	}
	return true
}

func (client *Client) applyCallbacks(messages []messages7.NetMessage) []messages7.NetMessage {
	filteredMessages := make([]messages7.NetMessage, 0, len(messages))

outer:
	for _, msg := range messages {
		for _, callback := range client.Callbacks.MessageOut {
			if !callback(&msg) {
				continue outer
			}
		}

		switch castMsg := msg.(type) {
		case *messages7.CtrlKeepAlive:
			if !userSendMsgCallback(client.Callbacks.CtrlKeepAliveOut, castMsg) {
				continue
			}
		case *messages7.CtrlConnect:
			if !userSendMsgCallback(client.Callbacks.CtrlConnectOut, castMsg) {
				continue
			}
		case *messages7.CtrlAccept:
			if !userSendMsgCallback(client.Callbacks.CtrlAcceptOut, castMsg) {
				continue
			}
		case *messages7.CtrlToken:
			if !userSendMsgCallback(client.Callbacks.CtrlTokenOut, castMsg) {
				continue
			}
		case *messages7.CtrlClose:
			if !userSendMsgCallback(client.Callbacks.CtrlCloseOut, castMsg) {
				continue
			}
		case *messages7.Info:
			if !userSendMsgCallback(client.Callbacks.SysInfoOut, castMsg) {
				continue
			}
		case *messages7.MapChange:
			if !userSendMsgCallback(client.Callbacks.SysMapChangeOut, castMsg) {
				continue
			}
		case *messages7.ServerInfo:
			if !userSendMsgCallback(client.Callbacks.SysServerInfoOut, castMsg) {
				continue
			}
		case *messages7.ConReady:
			if !userSendMsgCallback(client.Callbacks.SysConReadyOut, castMsg) {
				continue
			}
		case *messages7.Snap:
			if !userSendMsgCallback(client.Callbacks.SysSnapOut, castMsg) {
				continue
			}
		case *messages7.SnapEmpty:
			if !userSendMsgCallback(client.Callbacks.SysSnapEmptyOut, castMsg) {
				continue
			}
		case *messages7.SnapSingle:
			if !userSendMsgCallback(client.Callbacks.SysSnapSingleOut, castMsg) {
				continue
			}
		case *messages7.SnapSmall:
			if !userSendMsgCallback(client.Callbacks.SysSnapSmallOut, castMsg) {
				continue
			}
		case *messages7.InputTiming:
			if !userSendMsgCallback(client.Callbacks.SysInputTimingOut, castMsg) {
				continue
			}
		case *messages7.RconAuthOn:
			if !userSendMsgCallback(client.Callbacks.SysRconAuthOnOut, castMsg) {
				continue
			}
		case *messages7.RconAuthOff:
			if !userSendMsgCallback(client.Callbacks.SysRconAuthOffOut, castMsg) {
				continue
			}
		case *messages7.RconLine:
			if !userSendMsgCallback(client.Callbacks.SysRconLineOut, castMsg) {
				continue
			}
		case *messages7.RconCmdAdd:
			if !userSendMsgCallback(client.Callbacks.SysRconCmdAddOut, castMsg) {
				continue
			}
		case *messages7.RconCmdRem:
			if !userSendMsgCallback(client.Callbacks.SysRconCmdRemOut, castMsg) {
				continue
			}
		case *messages7.AuthChallenge:
			if !userSendMsgCallback(client.Callbacks.SysAuthChallengeOut, castMsg) {
				continue
			}
		case *messages7.AuthResult:
			if !userSendMsgCallback(client.Callbacks.SysAuthResultOut, castMsg) {
				continue
			}
		case *messages7.Ready:
			if !userSendMsgCallback(client.Callbacks.SysReadyOut, castMsg) {
				continue
			}
		case *messages7.EnterGame:
			if !userSendMsgCallback(client.Callbacks.SysEnterGameOut, castMsg) {
				continue
			}
		case *messages7.Input:
			if !userSendMsgCallback(client.Callbacks.SysInputOut, castMsg) {
				continue
			}
			client.Game.LastSentInput = *castMsg
			client.LastInputSend = time.Now()
		case *messages7.RconCmd:
			if !userSendMsgCallback(client.Callbacks.SysRconCmdOut, castMsg) {
				continue
			}
		case *messages7.RconAuth:
			if !userSendMsgCallback(client.Callbacks.SysRconAuthOut, castMsg) {
				continue
			}
		case *messages7.RequestMapData:
			if !userSendMsgCallback(client.Callbacks.SysRequestMapDataOut, castMsg) {
				continue
			}
		case *messages7.AuthStart:
			if !userSendMsgCallback(client.Callbacks.SysAuthStartOut, castMsg) {
				continue
			}
		case *messages7.AuthResponse:
			if !userSendMsgCallback(client.Callbacks.SysAuthResponseOut, castMsg) {
				continue
			}
		case *messages7.Ping:
			if !userSendMsgCallback(client.Callbacks.SysPingOut, castMsg) {
				continue
			}
		case *messages7.PingReply:
			if !userSendMsgCallback(client.Callbacks.SysPingReplyOut, castMsg) {
				continue
			}
		case *messages7.Error:
			if !userSendMsgCallback(client.Callbacks.SysErrorOut, castMsg) {
				continue
			}
		case *messages7.MaplistEntryAdd:
			if !userSendMsgCallback(client.Callbacks.SysMaplistEntryAddOut, castMsg) {
				continue
			}
		case *messages7.MaplistEntryRem:
			if !userSendMsgCallback(client.Callbacks.SysMaplistEntryRemOut, castMsg) {
				continue
			}
		case *messages7.SvMotd:
			if !userSendMsgCallback(client.Callbacks.GameSvMotdOut, castMsg) {
				continue
			}
		case *messages7.SvBroadcast:
			if !userSendMsgCallback(client.Callbacks.GameSvBroadcastOut, castMsg) {
				continue
			}
		case *messages7.SvChat:
			if !userSendMsgCallback(client.Callbacks.GameSvChatOut, castMsg) {
				continue
			}
		case *messages7.SvTeam:
			if !userSendMsgCallback(client.Callbacks.GameSvTeamOut, castMsg) {
				continue
			}
		case *messages7.SvKillMsg:
			if !userSendMsgCallback(client.Callbacks.GameSvKillMsgOut, castMsg) {
				continue
			}
		case *messages7.SvTuneParams:
			if !userSendMsgCallback(client.Callbacks.GameSvTuneParamsOut, castMsg) {
				continue
			}
		case *messages7.SvExtraProjectile:
			if !userSendMsgCallback(client.Callbacks.GameSvExtraProjectileOut, castMsg) {
				continue
			}
		case *messages7.SvReadyToEnter:
			if !userSendMsgCallback(client.Callbacks.GameSvReadyToEnterOut, castMsg) {
				continue
			}
		case *messages7.SvWeaponPickup:
			if !userSendMsgCallback(client.Callbacks.GameSvWeaponPickupOut, castMsg) {
				continue
			}
		case *messages7.SvEmoticon:
			if !userSendMsgCallback(client.Callbacks.GameSvEmoticonOut, castMsg) {
				continue
			}
		case *messages7.SvVoteClearOptions:
			if !userSendMsgCallback(client.Callbacks.GameSvVoteClearOptionsOut, castMsg) {
				continue
			}
		case *messages7.SvVoteOptionListAdd:
			if !userSendMsgCallback(client.Callbacks.GameSvVoteOptionListAddOut, castMsg) {
				continue
			}
		case *messages7.SvVoteOptionAdd:
			if !userSendMsgCallback(client.Callbacks.GameSvVoteOptionAddOut, castMsg) {
				continue
			}
		case *messages7.SvVoteOptionRemove:
			if !userSendMsgCallback(client.Callbacks.GameSvVoteOptionRemoveOut, castMsg) {
				continue
			}
		case *messages7.SvVoteSet:
			if !userSendMsgCallback(client.Callbacks.GameSvVoteSetOut, castMsg) {
				continue
			}
		case *messages7.SvVoteStatus:
			if !userSendMsgCallback(client.Callbacks.GameSvVoteStatusOut, castMsg) {
				continue
			}
		case *messages7.SvServerSettings:
			if !userSendMsgCallback(client.Callbacks.GameSvServerSettingsOut, castMsg) {
				continue
			}
		case *messages7.SvClientInfo:
			if !userSendMsgCallback(client.Callbacks.GameSvClientInfoOut, castMsg) {
				continue
			}
		case *messages7.SvGameInfo:
			if !userSendMsgCallback(client.Callbacks.GameSvGameInfoOut, castMsg) {
				continue
			}
		case *messages7.SvClientDrop:
			if !userSendMsgCallback(client.Callbacks.GameSvClientDropOut, castMsg) {
				continue
			}
		case *messages7.SvGameMsg:
			if !userSendMsgCallback(client.Callbacks.GameSvGameMsgOut, castMsg) {
				continue
			}
		case *messages7.DeClientEnter:
			if !userSendMsgCallback(client.Callbacks.GameDeClientEnterOut, castMsg) {
				continue
			}
		case *messages7.DeClientLeave:
			if !userSendMsgCallback(client.Callbacks.GameDeClientLeaveOut, castMsg) {
				continue
			}
		case *messages7.ClSay:
			if !userSendMsgCallback(client.Callbacks.GameClSayOut, castMsg) {
				continue
			}
		case *messages7.ClSetTeam:
			if !userSendMsgCallback(client.Callbacks.GameClSetTeamOut, castMsg) {
				continue
			}
		case *messages7.ClSetSpectatorMode:
			if !userSendMsgCallback(client.Callbacks.GameClSetSpectatorModeOut, castMsg) {
				continue
			}
		case *messages7.ClStartInfo:
			if !userSendMsgCallback(client.Callbacks.GameClStartInfoOut, castMsg) {
				continue
			}
		case *messages7.ClKill:
			if !userSendMsgCallback(client.Callbacks.GameClKillOut, castMsg) {
				continue
			}
		case *messages7.ClReadyChange:
			if !userSendMsgCallback(client.Callbacks.GameClReadyChangeOut, castMsg) {
				continue
			}
		case *messages7.ClEmoticon:
			if !userSendMsgCallback(client.Callbacks.GameClEmoticonOut, castMsg) {
				continue
			}
		case *messages7.ClVote:
			if !userSendMsgCallback(client.Callbacks.GameClVoteOut, castMsg) {
				continue
			}
		case *messages7.ClCallVote:
			if !userSendMsgCallback(client.Callbacks.GameClCallVoteOut, castMsg) {
				continue
			}
		case *messages7.SvSkinChange:
			if !userSendMsgCallback(client.Callbacks.GameSvSkinChangeOut, castMsg) {
				continue
			}
		case *messages7.ClSkinChange:
			if !userSendMsgCallback(client.Callbacks.GameClSkinChangeOut, castMsg) {
				continue
			}
		case *messages7.SvRaceFinish:
			if !userSendMsgCallback(client.Callbacks.GameSvRaceFinishOut, castMsg) {
				continue
			}
		case *messages7.SvCheckpoint:
			if !userSendMsgCallback(client.Callbacks.GameSvCheckpointOut, castMsg) {
				continue
			}
		case *messages7.SvCommandInfo:
			if !userSendMsgCallback(client.Callbacks.GameSvCommandInfoOut, castMsg) {
				continue
			}
		case *messages7.SvCommandInfoRemove:
			if !userSendMsgCallback(client.Callbacks.GameSvCommandInfoRemoveOut, castMsg) {
				continue
			}
		case *messages7.ClCommand:
			if !userSendMsgCallback(client.Callbacks.GameClCommandOut, castMsg) {
				continue
			}
		case *messages7.Unknown:
		default:
			// TODO: remove this
			panic(fmt.Sprintf("unprocessed msg=%d sys=%v", msg.MsgType(), msg.System()))

			// slog.Error(" unhooked out msg", "type", msg.MsgType(), "system", msg.System())
		}
		// slog.Info("  filter adding msg", "type", msg.MsgType(), "system", msg.System())
		filteredMessages = append(filteredMessages, msg)
	}

	return filteredMessages
}

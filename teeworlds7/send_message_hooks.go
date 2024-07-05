package teeworlds7

import (
	"fmt"
	"time"

	"github.com/teeworlds-go/protocol/messages7"
)

func userSendMsgCallback[T any](userCallbacks []func(T) bool, msg T) bool {
	for _, callback := range userCallbacks {
		if callback(msg) == false {
			return false
		}
	}
	return true
}

func (client *Client) registerMessagesCallbacks(messages []messages7.NetMessage) []messages7.NetMessage {
	filteredMessages := make([]messages7.NetMessage, 0, len(messages))

	for _, msg := range messages {
		for _, callback := range client.Callbacks.MessageOut {
			if callback(&msg) == false {
				continue
			}
		}

		switch castMsg := msg.(type) {
		case *messages7.CtrlKeepAlive:
			if userSendMsgCallback(client.Callbacks.CtrlKeepAliveOut, castMsg) == false {
				continue
			}
		case *messages7.CtrlConnect:
			if userSendMsgCallback(client.Callbacks.CtrlConnectOut, castMsg) == false {
				continue
			}
		case *messages7.CtrlAccept:
			if userSendMsgCallback(client.Callbacks.CtrlAcceptOut, castMsg) == false {
				continue
			}
		case *messages7.CtrlToken:
			if userSendMsgCallback(client.Callbacks.CtrlTokenOut, castMsg) == false {
				continue
			}
		case *messages7.CtrlClose:
			if userSendMsgCallback(client.Callbacks.CtrlCloseOut, castMsg) == false {
				continue
			}
		case *messages7.Info:
			if userSendMsgCallback(client.Callbacks.SysInfoOut, castMsg) == false {
				continue
			}
		case *messages7.MapChange:
			if userSendMsgCallback(client.Callbacks.SysMapChangeOut, castMsg) == false {
				continue
			}
		case *messages7.ServerInfo:
			if userSendMsgCallback(client.Callbacks.SysServerInfoOut, castMsg) == false {
				continue
			}
		case *messages7.ConReady:
			if userSendMsgCallback(client.Callbacks.SysConReadyOut, castMsg) == false {
				continue
			}
		case *messages7.Snap:
			if userSendMsgCallback(client.Callbacks.SysSnapOut, castMsg) == false {
				continue
			}
		case *messages7.SnapEmpty:
			if userSendMsgCallback(client.Callbacks.SysSnapEmptyOut, castMsg) == false {
				continue
			}
		case *messages7.SnapSingle:
			if userSendMsgCallback(client.Callbacks.SysSnapSingleOut, castMsg) == false {
				continue
			}
		case *messages7.SnapSmall:
			if userSendMsgCallback(client.Callbacks.SysSnapSmallOut, castMsg) == false {
				continue
			}
		case *messages7.InputTiming:
			if userSendMsgCallback(client.Callbacks.SysInputTimingOut, castMsg) == false {
				continue
			}
		case *messages7.RconAuthOn:
			if userSendMsgCallback(client.Callbacks.SysRconAuthOnOut, castMsg) == false {
				continue
			}
		case *messages7.RconAuthOff:
			if userSendMsgCallback(client.Callbacks.SysRconAuthOffOut, castMsg) == false {
				continue
			}
		case *messages7.RconLine:
			if userSendMsgCallback(client.Callbacks.SysRconLineOut, castMsg) == false {
				continue
			}
		case *messages7.RconCmdAdd:
			if userSendMsgCallback(client.Callbacks.SysRconCmdAddOut, castMsg) == false {
				continue
			}
		case *messages7.RconCmdRem:
			if userSendMsgCallback(client.Callbacks.SysRconCmdRemOut, castMsg) == false {
				continue
			}
		case *messages7.AuthChallenge:
			if userSendMsgCallback(client.Callbacks.SysAuthChallengeOut, castMsg) == false {
				continue
			}
		case *messages7.AuthResult:
			if userSendMsgCallback(client.Callbacks.SysAuthResultOut, castMsg) == false {
				continue
			}
		case *messages7.Ready:
			if userSendMsgCallback(client.Callbacks.SysReadyOut, castMsg) == false {
				continue
			}
		case *messages7.EnterGame:
			if userSendMsgCallback(client.Callbacks.SysEnterGameOut, castMsg) == false {
				continue
			}
		case *messages7.Input:
			if userSendMsgCallback(client.Callbacks.SysInputOut, castMsg) == false {
				continue
			}
			client.Game.LastSentInput = *castMsg
			client.LastInputSend = time.Now()
		case *messages7.RconCmd:
			if userSendMsgCallback(client.Callbacks.SysRconCmdOut, castMsg) == false {
				continue
			}
		case *messages7.RconAuth:
			if userSendMsgCallback(client.Callbacks.SysRconAuthOut, castMsg) == false {
				continue
			}
		case *messages7.RequestMapData:
			if userSendMsgCallback(client.Callbacks.SysRequestMapDataOut, castMsg) == false {
				continue
			}
		case *messages7.AuthStart:
			if userSendMsgCallback(client.Callbacks.SysAuthStartOut, castMsg) == false {
				continue
			}
		case *messages7.AuthResponse:
			if userSendMsgCallback(client.Callbacks.SysAuthResponseOut, castMsg) == false {
				continue
			}
		case *messages7.Ping:
			if userSendMsgCallback(client.Callbacks.SysPingOut, castMsg) == false {
				continue
			}
		case *messages7.PingReply:
			if userSendMsgCallback(client.Callbacks.SysPingReplyOut, castMsg) == false {
				continue
			}
		case *messages7.Error:
			if userSendMsgCallback(client.Callbacks.SysErrorOut, castMsg) == false {
				continue
			}
		case *messages7.MaplistEntryAdd:
			if userSendMsgCallback(client.Callbacks.SysMaplistEntryAddOut, castMsg) == false {
				continue
			}
		case *messages7.MaplistEntryRem:
			if userSendMsgCallback(client.Callbacks.SysMaplistEntryRemOut, castMsg) == false {
				continue
			}
		case *messages7.SvMotd:
			if userSendMsgCallback(client.Callbacks.GameSvMotdOut, castMsg) == false {
				continue
			}
		case *messages7.SvBroadcast:
			if userSendMsgCallback(client.Callbacks.GameSvBroadcastOut, castMsg) == false {
				continue
			}
		case *messages7.SvChat:
			if userSendMsgCallback(client.Callbacks.GameSvChatOut, castMsg) == false {
				continue
			}
		case *messages7.SvTeam:
			if userSendMsgCallback(client.Callbacks.GameSvTeamOut, castMsg) == false {
				continue
			}
		case *messages7.SvKillMsg:
			if userSendMsgCallback(client.Callbacks.GameSvKillMsgOut, castMsg) == false {
				continue
			}
		case *messages7.SvTuneParams:
			if userSendMsgCallback(client.Callbacks.GameSvTuneParamsOut, castMsg) == false {
				continue
			}
		case *messages7.SvExtraProjectile:
			if userSendMsgCallback(client.Callbacks.GameSvExtraProjectileOut, castMsg) == false {
				continue
			}
		case *messages7.SvReadyToEnter:
			if userSendMsgCallback(client.Callbacks.GameSvReadyToEnterOut, castMsg) == false {
				continue
			}
		case *messages7.SvWeaponPickup:
			if userSendMsgCallback(client.Callbacks.GameSvWeaponPickupOut, castMsg) == false {
				continue
			}
		case *messages7.SvEmoticon:
			if userSendMsgCallback(client.Callbacks.GameSvEmoticonOut, castMsg) == false {
				continue
			}
		case *messages7.SvVoteClearOptions:
			if userSendMsgCallback(client.Callbacks.GameSvVoteClearOptionsOut, castMsg) == false {
				continue
			}
		case *messages7.SvVoteOptionListAdd:
			if userSendMsgCallback(client.Callbacks.GameSvVoteOptionListAddOut, castMsg) == false {
				continue
			}
		case *messages7.SvVoteOptionAdd:
			if userSendMsgCallback(client.Callbacks.GameSvVoteOptionAddOut, castMsg) == false {
				continue
			}
		case *messages7.SvVoteOptionRemove:
			if userSendMsgCallback(client.Callbacks.GameSvVoteOptionRemoveOut, castMsg) == false {
				continue
			}
		case *messages7.SvVoteSet:
			if userSendMsgCallback(client.Callbacks.GameSvVoteSetOut, castMsg) == false {
				continue
			}
		case *messages7.SvVoteStatus:
			if userSendMsgCallback(client.Callbacks.GameSvVoteStatusOut, castMsg) == false {
				continue
			}
		case *messages7.SvServerSettings:
			if userSendMsgCallback(client.Callbacks.GameSvServerSettingsOut, castMsg) == false {
				continue
			}
		case *messages7.SvClientInfo:
			if userSendMsgCallback(client.Callbacks.GameSvClientInfoOut, castMsg) == false {
				continue
			}
		case *messages7.SvGameInfo:
			if userSendMsgCallback(client.Callbacks.GameSvGameInfoOut, castMsg) == false {
				continue
			}
		case *messages7.SvClientDrop:
			if userSendMsgCallback(client.Callbacks.GameSvClientDropOut, castMsg) == false {
				continue
			}
		case *messages7.SvGameMsg:
			if userSendMsgCallback(client.Callbacks.GameSvGameMsgOut, castMsg) == false {
				continue
			}
		case *messages7.DeClientEnter:
			if userSendMsgCallback(client.Callbacks.GameDeClientEnterOut, castMsg) == false {
				continue
			}
		case *messages7.DeClientLeave:
			if userSendMsgCallback(client.Callbacks.GameDeClientLeaveOut, castMsg) == false {
				continue
			}
		case *messages7.ClSay:
			if userSendMsgCallback(client.Callbacks.GameClSayOut, castMsg) == false {
				continue
			}
		case *messages7.ClSetTeam:
			if userSendMsgCallback(client.Callbacks.GameClSetTeamOut, castMsg) == false {
				continue
			}
		case *messages7.ClSetSpectatorMode:
			if userSendMsgCallback(client.Callbacks.GameClSetSpectatorModeOut, castMsg) == false {
				continue
			}
		case *messages7.ClStartInfo:
			if userSendMsgCallback(client.Callbacks.GameClStartInfoOut, castMsg) == false {
				continue
			}
		case *messages7.ClKill:
			if userSendMsgCallback(client.Callbacks.GameClKillOut, castMsg) == false {
				continue
			}
		case *messages7.ClReadyChange:
			if userSendMsgCallback(client.Callbacks.GameClReadyChangeOut, castMsg) == false {
				continue
			}
		case *messages7.ClEmoticon:
			if userSendMsgCallback(client.Callbacks.GameClEmoticonOut, castMsg) == false {
				continue
			}
		case *messages7.ClVote:
			if userSendMsgCallback(client.Callbacks.GameClVoteOut, castMsg) == false {
				continue
			}
		case *messages7.ClCallVote:
			if userSendMsgCallback(client.Callbacks.GameClCallVoteOut, castMsg) == false {
				continue
			}
		case *messages7.SvSkinChange:
			if userSendMsgCallback(client.Callbacks.GameSvSkinChangeOut, castMsg) == false {
				continue
			}
		case *messages7.ClSkinChange:
			if userSendMsgCallback(client.Callbacks.GameClSkinChangeOut, castMsg) == false {
				continue
			}
		case *messages7.SvRaceFinish:
			if userSendMsgCallback(client.Callbacks.GameSvRaceFinishOut, castMsg) == false {
				continue
			}
		case *messages7.SvCheckpoint:
			if userSendMsgCallback(client.Callbacks.GameSvCheckpointOut, castMsg) == false {
				continue
			}
		case *messages7.SvCommandInfo:
			if userSendMsgCallback(client.Callbacks.GameSvCommandInfoOut, castMsg) == false {
				continue
			}
		case *messages7.SvCommandInfoRemove:
			if userSendMsgCallback(client.Callbacks.GameSvCommandInfoRemoveOut, castMsg) == false {
				continue
			}
		case *messages7.ClCommand:
			if userSendMsgCallback(client.Callbacks.GameClCommandOut, castMsg) == false {
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

package teeworlds7

import (
	"fmt"
	"log/slog"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/packer"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

func (client *Client) processSystem(netMsg messages7.NetMessage, response *protocol7.Packet) bool {
	switch msg := netMsg.(type) {
	case *messages7.MapChange:
		userMsgCallback(client.Callbacks.SysMapChange, msg, func() {
			fmt.Println("got map change")
			response.Messages = append(response.Messages, &messages7.Ready{})
		})
	case *messages7.MapData:
		userMsgCallback(client.Callbacks.SysMapData, msg, func() {
			fmt.Printf("got map chunk %x\n", msg.Data)
		})
	case *messages7.ServerInfo:
		userMsgCallback(client.Callbacks.SysServerInfo, msg, func() {
			fmt.Printf("connected to server with name '%s'\n", msg.Name)
		})
	case *messages7.ConReady:
		userMsgCallback(client.Callbacks.SysConReady, msg, func() {
			fmt.Println("connected, sending info")
			info := &messages7.ClStartInfo{
				Name:                  client.Name,
				Clan:                  client.Clan,
				Country:               client.Country,
				Body:                  "greensward",
				Marking:               "duodonny",
				Decoration:            "",
				Hands:                 "standard",
				Feet:                  "standard",
				Eyes:                  "standard",
				CustomColorBody:       true,
				CustomColorMarking:    true,
				CustomColorDecoration: false,
				CustomColorHands:      false,
				CustomColorFeet:       false,
				CustomColorEyes:       false,
				ColorBody:             5635840,
				ColorMarking:          -11141356,
				ColorDecoration:       65408,
				ColorHands:            65408,
				ColorFeet:             65408,
				ColorEyes:             65408,
			}
			response.Messages = append(response.Messages, info)
		})
	case *messages7.Snap:
		userMsgCallback(client.Callbacks.SysSnap, msg, func() {
			deltaTick := msg.GameTick - msg.DeltaTick
			slog.Debug("got snap", "delta_tick", deltaTick, "raw_delta_tick", msg.DeltaTick, "game_tick", msg.GameTick, "part", msg.Part, "num_parts", msg.NumParts)

			err := client.SnapshotStorage.AddIncomingData(msg.Part, msg.NumParts, msg.Data)
			if err != nil {
				// TODO: dont panic
				panic(err)
			}

			// TODO: this is as naive as it gets
			//       we should check if we actually received all the previous parts
			//       teeworlds does some fancy bit stuff here
			//       m_SnapshotParts |= 1<<Part;
			if msg.Part != msg.NumParts-1 {
				// TODO: remove this print
				slog.Info("storing partial snap", "part", msg.Part, "num_parts", msg.NumParts)
				return
			}

			prevSnap, found := client.SnapshotStorage.Get(deltaTick)
			if !found {
				// couldn't find the delta snapshots that the server used
				// to compress this snapshot. force the server to resync
				slog.Error("error, couldn't find the delta snapshot")

				// ack snapshot
				// TODO:
				// m_AckGameTick = -1;
				return
			}

			u := &packer.Unpacker{}
			u.Reset(client.SnapshotStorage.IncomingData())

			newFullSnap, err := snapshot7.UnpackDelta(prevSnap, u)
			if err != nil {
				slog.Error("delta unpack failed!", "error", err)
				return
			}
			err = client.SnapshotStorage.Add(msg.GameTick, newFullSnap)
			if err != nil {
				slog.Error("failed to store snap", "error", err)
			}
			client.SnapshotStorage.PurgeUntil(deltaTick)

			for _, callback := range client.Callbacks.Snapshot {
				callback(newFullSnap, func() {})
			}

			client.Game.Input.AckGameTick = msg.GameTick
			client.Game.Input.PredictionTick = client.SnapshotStorage.NewestTick()
			client.Game.Snap.fill(newFullSnap)
			client.SnapshotStorage.SetAltSnap(msg.GameTick, newFullSnap)

			response.Messages = append(response.Messages, client.Game.Input)
		})
	case *messages7.SnapSingle:
		userMsgCallback(client.Callbacks.SysSnapSingle, msg, func() {
			deltaTick := msg.GameTick - msg.DeltaTick
			slog.Debug("got snap single", "delta_tick", deltaTick, "raw_delta_tick", msg.DeltaTick, "game_tick", msg.GameTick)
			prevSnap, found := client.SnapshotStorage.Get(deltaTick)

			if !found {
				// couldn't find the delta snapshots that the server used
				// to compress this snapshot. force the server to resync
				slog.Error("error, couldn't find the delta snapshot")

				// ack snapshot
				// TODO:
				// m_AckGameTick = -1;
				return
			}

			u := &packer.Unpacker{}
			u.Reset(msg.Data)

			newFullSnap, err := snapshot7.UnpackDelta(prevSnap, u)
			if err != nil {
				slog.Error("delta unpack failed!", "error", err)
				return
			}
			err = client.SnapshotStorage.Add(msg.GameTick, newFullSnap)
			if err != nil {
				slog.Error("failed to store snap", "error", err)
			}
			client.SnapshotStorage.PurgeUntil(deltaTick)

			for _, callback := range client.Callbacks.Snapshot {
				callback(newFullSnap, func() {})
			}

			client.Game.Input.AckGameTick = msg.GameTick
			client.Game.Input.PredictionTick = client.SnapshotStorage.NewestTick()

			// altSnap := client.CreateAltSnap(prevSnap, newFullSnap)
			altSnap := newFullSnap
			client.Game.Snap.fill(altSnap)
			client.SnapshotStorage.SetAltSnap(msg.GameTick, altSnap)

			client.SendInput()
			response.Messages = append(response.Messages, client.Game.Input)
		})
	case *messages7.SnapEmpty:
		userMsgCallback(client.Callbacks.SysSnapEmpty, msg, func() {
			deltaTick := msg.GameTick - msg.DeltaTick
			slog.Debug("got snap empty", "delta_tick", deltaTick, "raw_delta_tick", msg.DeltaTick, "game_tick", msg.GameTick)
			prevSnap, found := client.SnapshotStorage.Get(deltaTick)

			if !found {
				// couldn't find the delta snapshots that the server used
				// to compress this snapshot. force the server to resync
				slog.Error("error, couldn't find the delta snapshot")

				// ack snapshot
				// TODO:
				// m_AckGameTick = -1;
				return
			}

			err := client.SnapshotStorage.Add(msg.GameTick, prevSnap)
			if err != nil {
				slog.Error("failed to store snap", "error", err)
			}
			client.SnapshotStorage.PurgeUntil(deltaTick)

			for _, callback := range client.Callbacks.Snapshot {
				callback(prevSnap, func() {})
			}

			client.Game.Input.AckGameTick = msg.GameTick
			client.Game.Input.PredictionTick = client.SnapshotStorage.NewestTick()
			// blazingly fast empty snaps
			// reuse the old game state
			// there is no need to refill if it is the same snapshot anyways
			// client.Game.Snap.fill(prevSnap)

			response.Messages = append(response.Messages, client.Game.Input)
		})
	case *messages7.InputTiming:
		userMsgCallback(client.Callbacks.SysInputTiming, msg, func() {
			// fmt.Printf("timing time left=%d\n", msg.TimeLeft)
		})
	case *messages7.RconAuthOn:
		userMsgCallback(client.Callbacks.SysRconAuthOn, msg, func() {
			fmt.Println("you are now authenticated in rcon")
		})
	case *messages7.RconAuthOff:
		userMsgCallback(client.Callbacks.SysRconAuthOff, msg, func() {
			fmt.Println("you are no longer authenticated in rcon")
		})
	case *messages7.RconLine:
		userMsgCallback(client.Callbacks.SysRconLine, msg, func() {
			fmt.Printf("[rcon] %s\n", msg.Line)
		})
	case *messages7.RconCmdAdd:
		userMsgCallback(client.Callbacks.SysRconCmdAdd, msg, func() {
			fmt.Printf("got rcon cmd=%s %s %s\n", msg.Name, msg.Params, msg.Help)
		})
	case *messages7.RconCmdRem:
		userMsgCallback(client.Callbacks.SysRconCmdRem, msg, func() {
			fmt.Printf("removed cmd=%s\n", msg.Name)
		})
	case *messages7.Unknown:
		userMsgCallback(client.Callbacks.MsgUnknown, msg, func() {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown system")
		})
	default:
		printUnknownMessage(netMsg, "unprocessed system")
		return false
	}
	return true
}

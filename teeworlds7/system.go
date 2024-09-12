package teeworlds7

import (
	"fmt"
	"log/slog"

	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/packer"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

func (client *Client) processSystem(netMsg messages7.NetMessage, response *protocol7.Packet) (process bool, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to process system message: %w", err)
		}
	}()

	switch msg := netMsg.(type) {
	case *messages7.MapChange:
		err = userMsgCallback(client.Callbacks.SysMapChange, msg, func() error {
			fmt.Println("got map change")
			response.Messages = append(response.Messages, &messages7.Ready{})
			return nil
		})
	case *messages7.MapData:
		err = userMsgCallback(client.Callbacks.SysMapData, msg, func() error {
			fmt.Printf("got map chunk %x\n", msg.Data)
			return nil
		})
	case *messages7.ServerInfo:
		err = userMsgCallback(client.Callbacks.SysServerInfo, msg, func() error {
			fmt.Printf("connected to server with name '%s'\n", msg.Name)
			return nil
		})
	case *messages7.ConReady:
		err = userMsgCallback(client.Callbacks.SysConReady, msg, func() error {
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
			return nil
		})
	case *messages7.Snap:
		err = userMsgCallback(client.Callbacks.SysSnap, msg, func() error {
			deltaTick := msg.GameTick - msg.DeltaTick
			slog.Debug("got snap", "delta_tick", deltaTick, "raw_delta_tick", msg.DeltaTick, "game_tick", msg.GameTick, "part", msg.Part, "num_parts", msg.NumParts)

			err := client.SnapshotStorage.AddIncomingData(msg.Part, msg.NumParts, msg.Data)
			if err != nil {
				return fmt.Errorf("failed to store incoming data snap: %w", err)
			}

			// TODO: this is as naive as it gets
			//       we should check if we actually received all the previous parts
			//       teeworlds does some fancy bit stuff here
			//       m_SnapshotParts |= 1<<Part;
			if msg.Part != msg.NumParts-1 {
				// TODO: remove this print
				slog.Debug("storing partial snap", "part", msg.Part, "num_parts", msg.NumParts)
				return nil
			}

			prevSnap, found := client.SnapshotStorage.Get(deltaTick)
			if !found {
				// couldn't find the delta snapshots that the server used
				// to compress this snapshot. force the server to resync
				slog.Error("error, couldn't find the delta snapshot")

				// ack snapshot
				// TODO:
				// m_AckGameTick = -1;
				return nil
			}

			u := &packer.Unpacker{}
			u.Reset(client.SnapshotStorage.IncomingData())

			newFullSnap, err := snapshot7.UnpackDelta(prevSnap, u)
			if err != nil {
				return fmt.Errorf("delta unpack failed: %w", err)
			}
			err = client.SnapshotStorage.Add(msg.GameTick, newFullSnap)
			if err != nil {
				return fmt.Errorf("failed to store snap: %w", err)
			}
			client.SnapshotStorage.PurgeUntil(deltaTick)

			for _, callback := range client.Callbacks.Snapshot {
				callback(newFullSnap, nil)
			}

			client.Game.Input.AckGameTick = msg.GameTick
			client.Game.Input.PredictionTick = client.SnapshotStorage.NewestTick()
			client.Game.Snap.fill(newFullSnap)
			client.SnapshotStorage.SetAltSnap(msg.GameTick, newFullSnap)

			response.Messages = append(response.Messages, client.Game.Input)
			return nil
		})
	case *messages7.SnapSingle:
		err = userMsgCallback(client.Callbacks.SysSnapSingle, msg, func() error {
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
				return nil
			}

			u := &packer.Unpacker{}
			u.Reset(msg.Data)

			newFullSnap, err := snapshot7.UnpackDelta(prevSnap, u)
			if err != nil {
				return fmt.Errorf("delta unpack failed: %w", err)
			}
			err = client.SnapshotStorage.Add(msg.GameTick, newFullSnap)
			if err != nil {
				return fmt.Errorf("failed to store snap: %w", err)
			}
			client.SnapshotStorage.PurgeUntil(deltaTick)

			for _, callback := range client.Callbacks.Snapshot {
				callback(newFullSnap, nil)
			}

			client.Game.Input.AckGameTick = msg.GameTick
			client.Game.Input.PredictionTick = client.SnapshotStorage.NewestTick()

			// altSnap := client.CreateAltSnap(prevSnap, newFullSnap)
			altSnap := newFullSnap
			client.Game.Snap.fill(altSnap)
			client.SnapshotStorage.SetAltSnap(msg.GameTick, altSnap)

			err = client.SendInput()
			if err != nil {
				return fmt.Errorf("failed to send input: %w", err)
			}
			response.Messages = append(response.Messages, client.Game.Input)
			return nil
		})
	case *messages7.SnapEmpty:
		err = userMsgCallback(client.Callbacks.SysSnapEmpty, msg, func() error {
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
				return nil
			}

			err := client.SnapshotStorage.Add(msg.GameTick, prevSnap)
			if err != nil {
				slog.Error("failed to store snap", "error", err)
			}
			client.SnapshotStorage.PurgeUntil(deltaTick)

			for _, callback := range client.Callbacks.Snapshot {
				err = callback(prevSnap, nil)
				if err != nil {
					return fmt.Errorf("failed to execute snapshot callback: %w", err)
				}
			}

			client.Game.Input.AckGameTick = msg.GameTick
			client.Game.Input.PredictionTick = client.SnapshotStorage.NewestTick()
			// blazingly fast empty snaps
			// reuse the old game state
			// there is no need to refill if it is the same snapshot anyways
			// client.Game.Snap.fill(prevSnap)

			response.Messages = append(response.Messages, client.Game.Input)
			return nil
		})
	case *messages7.InputTiming:
		err = userMsgCallback(client.Callbacks.SysInputTiming, msg, func() error {
			// fmt.Printf("timing time left=%d\n", msg.TimeLeft)
			return nil
		})
	case *messages7.RconAuthOn:
		err = userMsgCallback(client.Callbacks.SysRconAuthOn, msg, func() error {
			fmt.Println("you are now authenticated in rcon")
			return nil
		})
	case *messages7.RconAuthOff:
		err = userMsgCallback(client.Callbacks.SysRconAuthOff, msg, func() error {
			fmt.Println("you are no longer authenticated in rcon")
			return nil
		})
	case *messages7.RconLine:
		err = userMsgCallback(client.Callbacks.SysRconLine, msg, func() error {
			fmt.Printf("[rcon] %s\n", msg.Line)
			return nil
		})
	case *messages7.RconCmdAdd:
		err = userMsgCallback(client.Callbacks.SysRconCmdAdd, msg, func() error {
			fmt.Printf("got rcon cmd=%s %s %s\n", msg.Name, msg.Params, msg.Help)
			return nil
		})
	case *messages7.RconCmdRem:
		err = userMsgCallback(client.Callbacks.SysRconCmdRem, msg, func() error {
			fmt.Printf("removed cmd=%s\n", msg.Name)
			return nil
		})
	case *messages7.Unknown:
		err = userMsgCallback(client.Callbacks.MsgUnknown, msg, func() error {
			// TODO: msg id of unknown messages should not be -1
			fmt.Println("TODO: why is the msg id -1???")
			printUnknownMessage(msg, "unknown system")
			return nil
		})
	default:
		printUnknownMessage(netMsg, "unprocessed system")
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

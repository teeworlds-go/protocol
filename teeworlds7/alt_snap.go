package teeworlds7

import (
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/snapshot7"
)

func evolveCharacter(char object7.Character, tick int) object7.Character {
	for ; char.Tick < tick; char.Tick++ {
		char.X += char.VelX
	}
	return char
}

// TODO: why is this on the client struct? wouldn't that be more useful as a standalone function?

// creates a copy of the new snapshot
// and evolves the characters to their new predicted position based
// on their old poisiton and velocity
//
// it should also drop invalid snap items that do not pass validation
func (client *Client) CreateAltSnap(oldSnap *snapshot7.Snapshot, newSnap *snapshot7.Snapshot) *snapshot7.Snapshot {
	altSnap := &snapshot7.Snapshot{}

	altSnap.Items = make([]object7.SnapObject, len(newSnap.Items))

	for i, newItem := range newSnap.Items {
		altSnap.Items[i] = newItem

		if newItem.TypeId() == network7.ObjCharacter {
			char, ok := altSnap.Items[i].(*object7.Character)
			if ok == false {
				panic("failed to cast character")
			}
			// TODO: this is wrong
			evolveTick := char.Tick + 10

			newChar := evolveCharacter(*char, evolveTick)
			altSnap.Items[i] = &newChar
		}
	}

	return altSnap
}

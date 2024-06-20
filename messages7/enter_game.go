package messages7

import (
	"github.com/teeworlds-go/teeworlds/network7"
)

type EnterGame struct {
}

func (msg EnterGame) MsgId() int {
	return network7.MsgSysEnterGame
}

func (msg EnterGame) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg EnterGame) System() bool {
	return true
}

func (msg EnterGame) Vital() bool {
	return true
}

func (msg EnterGame) Pack() []byte {
	return []byte{}
}

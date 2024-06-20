package messages7

import (
	"github.com/teeworlds-go/teeworlds/network7"
)

type Ready struct {
}

func (msg Ready) MsgId() int {
	return network7.MsgSysReady
}

func (msg Ready) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg Ready) System() bool {
	return true
}

func (msg Ready) Vital() bool {
	return true
}

func (msg Ready) Pack() []byte {
	return []byte{}
}

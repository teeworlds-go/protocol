package messages7

import (
	"github.com/teeworlds-go/teeworlds/network7"
)

type CtrlKeepAlive struct {
}

func (msg CtrlKeepAlive) MsgId() int {
	return network7.MsgCtrlKeepAlive
}

func (msg CtrlKeepAlive) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg CtrlKeepAlive) System() bool {
	return false
}

func (msg CtrlKeepAlive) Vital() bool {
	return false
}

func (msg CtrlKeepAlive) Pack() []byte {
	return []byte{network7.MsgCtrlKeepAlive}
}

package network7

const (
	MaxClients = 64

	MsgCtrlKeepAlive ControlMsg = 0x00
	MsgCtrlConnect   ControlMsg = 0x01
	MsgCtrlAccept    ControlMsg = 0x02
	MsgCtrlToken     ControlMsg = 0x05
	MsgCtrlClose     ControlMsg = 0x04

	MsgSysMapChange  NetMsg = 2
	MsgSysConReady   NetMsg = 5
	MsgSysSnapSingle NetMsg = 8

	MsgGameSvMotd       NetMsg = 1
	MsgGameSvChat       NetMsg = 3
	MsgGameReadyToEnter NetMsg = 8
	MsgGameSvClientInfo NetMsg = 18
	MsgGameClStartInfo  NetMsg = 27
)

type ControlMsg int
type NetMsg int


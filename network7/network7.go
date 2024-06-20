package network7

const (
	MaxClients = 64

	MsgCtrlKeepAlive = 0x00
	MsgCtrlConnect   = 0x01
	MsgCtrlAccept    = 0x02
	MsgCtrlToken     = 0x05
	MsgCtrlClose     = 0x04

	MsgSysInfo       = 1
	MsgSysMapChange  = 2
	MsgSysConReady   = 5
	MsgSysSnapSingle = 8
	MsgSysReady      = 18
	MsgSysEnterGame  = 19

	MsgGameSvMotd       = 1
	MsgGameSvChat       = 3
	MsgGameReadyToEnter = 8
	MsgGameSvClientInfo = 18
	MsgGameClStartInfo  = 27

	TypeControl  MsgType = 1
	TypeNet      MsgType = 2
	TypeConnless MsgType = 3
)

type MsgType int

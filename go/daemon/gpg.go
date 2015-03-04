package main

import (
	keybase_1 "github.com/keybase/client/protocol/go"
	"github.com/maxtaco/go-framed-msgpack-rpc/rpc2"
)

type RemoteGPGUI struct {
	sessionId int
	uicli     keybase_1.GpgUiClient
}

func NewRemoteGPGUI(sessionId int, c *rpc2.Client) *RemoteGPGUI {
	return &RemoteGPGUI{
		sessionId: sessionId,
		uicli:     keybase_1.GpgUiClient{Cli: c},
	}
}

func (r *RemoteGPGUI) SelectKey(arg keybase_1.SelectKeyArg) (string, error) {
	arg.SessionID = r.sessionId
	return r.uicli.SelectKey(arg)
}

func (r *RemoteGPGUI) SelectKeyAndPushOption(arg keybase_1.SelectKeyAndPushOptionArg) (keybase_1.SelectKeyRes, error) {
	arg.SessionID = r.sessionId
	return r.uicli.SelectKeyAndPushOption(arg)
}

func (r *RemoteGPGUI) WantToAddGPGKey(int) (bool, error) {
	return r.uicli.WantToAddGPGKey(r.sessionId)
}

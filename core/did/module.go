package did

import (
	didtypes "github.com/Moonyongjung/xpla-private-chain/x/did/types"
	"github.com/Moonyongjung/xpriv.go/core"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	cmclient "github.com/cosmos/cosmos-sdk/client"
)

type coreModule struct{}

func NewCoreModule() core.CoreModule {
	return &coreModule{}
}

func (c *coreModule) Name() string {
	return DidModule
}

func (c *coreModule) NewTxRouter(builder cmclient.TxBuilder, msgType string, msg interface{}) (cmclient.TxBuilder, error) {
	switch {
	case msgType == DidCreateDidMsgType:
		convertMsg := msg.(didtypes.MsgCreateDID)
		builder.SetMsgs(&convertMsg)

	case msgType == DidUpdateDidMsgType:
		convertMsg := msg.(didtypes.MsgUpdateDID)
		builder.SetMsgs(&convertMsg)

	case msgType == DidDeactivateDidMsgType:
		convertMsg := msg.(didtypes.MsgDeactivateDID)
		builder.SetMsgs(&convertMsg)

	case msgType == DidReplaceDidMonikerMsgType:
		convertMsg := msg.(didtypes.MsgReplaceDIDMoniker)
		builder.SetMsgs(&convertMsg)

	default:
		return nil, util.LogErr(errors.ErrInvalidMsgType, msgType)
	}

	return builder, nil
}

func (c *coreModule) NewQueryRouter(q core.QueryClient) (string, error) {
	return QueryDID(q)
}

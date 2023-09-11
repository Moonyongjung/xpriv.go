package anchor

import (
	anchortypes "github.com/Moonyongjung/xpla-private-chain/x/anchor/types"
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
	return AnchorModule
}

func (c *coreModule) NewTxRouter(builder cmclient.TxBuilder, msgType string, msg interface{}) (cmclient.TxBuilder, error) {
	switch {
	case msgType == AnchorRegisterAnchorAccMsgType:
		convertMsg := msg.(anchortypes.MsgRegisterAnchorAcc)
		builder.SetMsgs(&convertMsg)

	case msgType == AnchorChangeAnchorAccMsgType:
		convertMsg := msg.(anchortypes.MsgChangeAnchorAcc)
		builder.SetMsgs(&convertMsg)

	default:
		return nil, util.LogErr(errors.ErrInvalidMsgType, msgType)
	}

	return builder, nil
}

func (c *coreModule) NewQueryRouter(q core.QueryClient) (string, error) {
	return QueryAnchor(q)
}

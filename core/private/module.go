package private

import (
	privatetypes "github.com/Moonyongjung/xpla-private-chain/x/private/types"
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
	return PrivateModule
}

func (c *coreModule) NewTxRouter(builder cmclient.TxBuilder, msgType string, msg interface{}) (cmclient.TxBuilder, error) {
	switch {
	case msgType == PrivateInitialAdminMsgType:
		convertMsg := msg.(privatetypes.MsgInitialAdmin)
		builder.SetMsgs(&convertMsg)

	case msgType == PrivateAddAdminMsgType:
		convertMsg := msg.(privatetypes.MsgAddAdmin)
		builder.SetMsgs(&convertMsg)

	case msgType == PrivateParticipateMsgType:
		convertMsg := msg.(privatetypes.MsgParticipate)
		builder.SetMsgs(&convertMsg)

	case msgType == PrivateAcceptMsgType:
		convertMsg := msg.(privatetypes.MsgAccept)
		builder.SetMsgs(&convertMsg)

	case msgType == PrivateDenyMsgType:
		convertMsg := msg.(privatetypes.MsgDeny)
		builder.SetMsgs(&convertMsg)

	case msgType == PrivateExileMsgType:
		convertMsg := msg.(privatetypes.MsgExile)
		builder.SetMsgs(&convertMsg)

	case msgType == PrivateQuitMsgType:
		convertMsg := msg.(privatetypes.MsgQuit)
		builder.SetMsgs(&convertMsg)

	default:
		return nil, util.LogErr(errors.ErrInvalidMsgType, msgType)
	}

	return builder, nil
}

func (c *coreModule) NewQueryRouter(q core.QueryClient) (string, error) {
	return QueryPrivate(q)
}

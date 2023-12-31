package mint

import (
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
	return MintModule
}

func (c *coreModule) NewTxRouter(_ cmclient.TxBuilder, _ string, _ interface{}) (cmclient.TxBuilder, error) {
	return nil, util.LogErr(errors.ErrInvalidRequest, c.Name(), "module has not tx")
}

func (c *coreModule) NewQueryRouter(q core.QueryClient) (string, error) {
	return QueryMint(q)
}

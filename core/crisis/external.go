package crisis

import (
	"github.com/Moonyongjung/xpriv.go/provider"
	"github.com/Moonyongjung/xpriv.go/types"
)

type CrisisExternal struct {
	Xplac provider.XplaClient
}

func NewCrisisExternal(xplac provider.XplaClient) (e CrisisExternal) {
	e.Xplac = xplac
	return e
}

// Tx

// Submit proof that an invariant broken to halt the chain.
func (e CrisisExternal) InvariantBroken(invariantBrokenMsg types.InvariantBrokenMsg) provider.XplaClient {
	msg, err := MakeInvariantRouteMsg(invariantBrokenMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(CrisisModule).
		WithMsgType(CrisisInvariantBrokenMsgType).
		WithMsg(msg)
	return e.Xplac
}

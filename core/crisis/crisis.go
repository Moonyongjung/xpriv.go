package crisis

import (
	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types"

	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
)

// (Tx) make msg - invariant broken
func MakeInvariantRouteMsg(invariantBrokenMsg types.InvariantBrokenMsg, privKey key.PrivateKey) (crisistypes.MsgVerifyInvariant, error) {
	return parseInvariantBrokenArgs(invariantBrokenMsg, privKey)
}

package crisis

import (
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
)

// Parsing - invariant broken
func parseInvariantBrokenArgs(invariantBrokenMsg types.InvariantBrokenMsg, privKey key.PrivateKey) (crisistypes.MsgVerifyInvariant, error) {
	if invariantBrokenMsg.ModuleName == "" || invariantBrokenMsg.InvariantRoute == "" {
		return crisistypes.MsgVerifyInvariant{}, util.LogErr(errors.ErrInsufficientParams, "invalid module name or invariant route")
	}

	senderAddr, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return crisistypes.MsgVerifyInvariant{}, util.LogErr(errors.ErrParse, err)
	}
	msg := crisistypes.NewMsgVerifyInvariant(senderAddr, invariantBrokenMsg.ModuleName, invariantBrokenMsg.InvariantRoute)

	return *msg, nil
}

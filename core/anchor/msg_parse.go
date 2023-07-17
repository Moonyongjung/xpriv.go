package anchor

import (
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	anchortypes "github.com/Moonyongjung/xpla-private-chain/x/anchor/types"
)

// Parsing - register anchor account
func ParseRegisterAnchorAccArgs(registerAnchorAccMsg types.RegisterAnchorAccMsg, privKey key.PrivateKey) (anchortypes.MsgRegisterAnchorAcc, error) {
	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return anchortypes.MsgRegisterAnchorAcc{}, util.LogErr(errors.ErrParse, err)
	}

	return anchortypes.NewMsgRegisterAnchorAcc(
		registerAnchorAccMsg.AnchorAccountAddr,
		registerAnchorAccMsg.ValidatorAddr,
		fromAddress.String(),
	), nil
}

// Parsing - change anchor account
func ParseChangeAnchorAccArgs(changeAnchorAccMsg types.ChangeAnchorAccMsg, privKey key.PrivateKey) (anchortypes.MsgChangeAnchorAcc, error) {
	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return anchortypes.MsgChangeAnchorAcc{}, util.LogErr(errors.ErrParse, err)
	}

	return anchortypes.NewMsgChangeAnchorAcc(
		changeAnchorAccMsg.AnchorAccountAddr,
		changeAnchorAccMsg.ValidatorAddr,
		fromAddress.String(),
	), nil
}

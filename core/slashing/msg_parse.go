package slashing

import (
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	"github.com/Moonyongjung/xpla-private-chain/app/params"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

// Parsing - unjail
func parseUnjailArgs(privKey key.PrivateKey) (slashingtypes.MsgUnjail, error) {
	addr, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return slashingtypes.MsgUnjail{}, util.LogErr(errors.ErrParse, err)
	}

	msg := slashingtypes.NewMsgUnjail(sdk.ValAddress(addr))

	return *msg, nil
}

// Parsing - signing info
func parseQuerySigingInfoArgs(signingInfoMsg types.SigningInfoMsg, xplacEncodingConfig params.EncodingConfig) (slashingtypes.QuerySigningInfoRequest, error) {
	if signingInfoMsg.ConsPubKey != "" {
		var pk cryptotypes.PubKey
		err := xplacEncodingConfig.Marshaler.UnmarshalInterfaceJSON([]byte(signingInfoMsg.ConsPubKey), &pk)
		if err != nil {
			return slashingtypes.QuerySigningInfoRequest{}, util.LogErr(errors.ErrFailedToUnmarshal, err)
		}

		return slashingtypes.QuerySigningInfoRequest{
			ConsAddress: sdk.ConsAddress(pk.Address()).String(),
		}, nil
	} else if signingInfoMsg.ConsAddr != "" {
		return slashingtypes.QuerySigningInfoRequest{
			ConsAddress: signingInfoMsg.ConsAddr,
		}, nil
	} else {
		return slashingtypes.QuerySigningInfoRequest{}, util.LogErr(errors.ErrInsufficientParams, "need at least one input")
	}
}

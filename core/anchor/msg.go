package anchor

import (
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/util"

	anchortypes "github.com/Moonyongjung/xpla-private-chain/x/anchor/types"
)

// (Tx) make msg - register anchor account
func MakeRegisterAnchorAccMsg(registerAnchorAccMsg types.RegisterAnchorAccMsg, privKey key.PrivateKey) (anchortypes.MsgRegisterAnchorAcc, error) {
	return ParseRegisterAnchorAccArgs(registerAnchorAccMsg, privKey)
}

// (Tx) make msg - change anchor account
func MakeChangeAnchorAccMsg(changeAnchorAccMsg types.ChangeAnchorAccMsg, privKey key.PrivateKey) (anchortypes.MsgChangeAnchorAcc, error) {
	return ParseChangeAnchorAccArgs(changeAnchorAccMsg, privKey)
}

// (Query) make msg - query anchor account
func MakeAnchorAccMsg(anchorAccMsg types.AnchorAccMsg) (anchortypes.QueryAnchorAccountRequest, error) {
	return anchortypes.QueryAnchorAccountRequest{
		ValidatorAddress: anchorAccMsg.ValidatorAddr,
	}, nil
}

// (Query) make msg - all aggregated blocks
func MakeAllAggregatedBlocksMsg() (anchortypes.QueryAllAggregatedBlocksRequest, error) {
	return anchortypes.QueryAllAggregatedBlocksRequest{}, nil
}

// (Query) make msg - anchor info
func MakeAnchorInfoMsg(anchorInfoMsg types.AnchorInfoMsg) (anchortypes.QueryAnchorInfoRequest, error) {
	heightUint64, err := util.FromStringToUint64(anchorInfoMsg.PrivChainHeight)
	if err != nil {
		return anchortypes.QueryAnchorInfoRequest{}, err
	}
	return anchortypes.QueryAnchorInfoRequest{
		Height: heightUint64,
	}, nil
}

// (Query) make msg - anchor block
func MakeAnchorBlockMsg(anchorBlockMsg types.AnchorBlockMsg) (anchortypes.QueryAnchorBlockRequest, error) {
	heightUint64, err := util.FromStringToUint64(anchorBlockMsg.PrivChainHeight)
	if err != nil {
		return anchortypes.QueryAnchorBlockRequest{}, err
	}
	return anchortypes.QueryAnchorBlockRequest{
		Height:          heightUint64,
		ContractAddress: anchorBlockMsg.AnchorContractAddr,
	}, nil
}

// (Query) make msg - anchor tx body
func MakeAnchorTxBodyMsg(anchorTxBodyMsg types.AnchorTxBodyMsg) (anchortypes.QueryAnchorTxBodyRequest, error) {
	heightUint64, err := util.FromStringToUint64(anchorTxBodyMsg.PrivChainHeight)
	if err != nil {
		return anchortypes.QueryAnchorTxBodyRequest{}, err
	}
	return anchortypes.QueryAnchorTxBodyRequest{
		Height: heightUint64,
	}, nil
}

// (Query) make msg - verify
func MakeAnchorVerifyMsg(anchorVerifyMsg types.AnchorVerifyMsg) (anchortypes.QueryVerifyRequest, error) {
	heightUint64, err := util.FromStringToUint64(anchorVerifyMsg.PrivChainHeight)
	if err != nil {
		return anchortypes.QueryVerifyRequest{}, err
	}
	return anchortypes.QueryVerifyRequest{
		Height:          heightUint64,
		ContractAddress: anchorVerifyMsg.AnchorContractAddr,
	}, nil
}

// (Query) make msg - anchor balances
func MakeAnchorBalancesMsg(anchorBalancesMsg types.AnchorBalancesMsg) (anchortypes.QueryAnchorBalancesRequest, error) {
	return anchortypes.QueryAnchorBalancesRequest{
		ValidatorAddress: anchorBalancesMsg.ValidatorAddr,
	}, nil
}

// (Query) make msg - params of anchor module
func MakeAnchorParamsMsg() (anchortypes.QueryParamsRequest, error) {
	return anchortypes.QueryParamsRequest{}, nil
}

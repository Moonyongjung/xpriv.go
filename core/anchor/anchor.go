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
	return anchortypes.QueryAnchorInfoRequest{
		Height: util.FromStringToUint64(anchorInfoMsg.PrivChainHeight),
	}, nil
}

// (Query) make msg - anchor block
func MakeAnchorBlockMsg(anchorBlockMsg types.AnchorBlockMsg) (anchortypes.QueryAnchorBlockRequest, error) {
	return anchortypes.QueryAnchorBlockRequest{
		Height:          util.FromStringToUint64(anchorBlockMsg.PrivChainHeight),
		ContractAddress: anchorBlockMsg.AnchorContractAddr,
	}, nil
}

// (Query) make msg - anchor tx body
func MakeAnchorTxBodyMsg(anchorTxBodyMsg types.AnchorTxBodyMsg) (anchortypes.QueryAnchorTxBodyRequest, error) {
	return anchortypes.QueryAnchorTxBodyRequest{
		Height: util.FromStringToUint64(anchorTxBodyMsg.PrivChainHeight),
	}, nil
}

// (Query) make msg - verify
func MakeAnchorVerifyMsg(anchorVerifyMsg types.AnchorVerifyMsg) (anchortypes.QueryVerifyRequest, error) {
	return anchortypes.QueryVerifyRequest{
		Height:          util.FromStringToUint64(anchorVerifyMsg.PrivChainHeight),
		ContractAddress: anchorVerifyMsg.AnchorContractAddr,
	}, nil
}

// (Query) make msg - anchor balances
func MakeAnchorBalancesMsg(anchorBalancesMsg types.AnchorBalancesMsg) (anchortypes.QueryAnchorBalancesRequest, error) {
	return anchortypes.QueryAnchorBalancesRequest{
		ValidatorAddress: anchorBalancesMsg.ValidatorAddr,
	}, nil
}

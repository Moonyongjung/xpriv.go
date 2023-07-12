package params

import (
	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types"

	"github.com/Moonyongjung/xpla-private-chain/app/params"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

// (Tx) make msg - param change
func MakeProposalParamChangeMsg(paramChangeMsg types.ParamChangeMsg, privKey key.PrivateKey, encodingConfig params.EncodingConfig) (govtypes.MsgSubmitProposal, error) {
	return parseProposalParamChangeArgs(paramChangeMsg, privKey, encodingConfig)
}

// (Query) make msg - subspace
func MakeQueryParamsSubspaceMsg(subspaceMsg types.SubspaceMsg) (paramsproposal.QueryParamsRequest, error) {
	return paramsproposal.QueryParamsRequest{
		Subspace: subspaceMsg.Subspace,
		Key:      subspaceMsg.Key,
	}, nil
}

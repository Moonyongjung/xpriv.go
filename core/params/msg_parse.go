package params

import (
	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"

	"github.com/Moonyongjung/xpla-private-chain/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramscutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	paramsproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

// Parsing - param change
func parseProposalParamChangeArgs(paramChangeMsg types.ParamChangeMsg, privKey key.PrivateKey, encodingConfig params.EncodingConfig) (govtypes.MsgSubmitProposal, error) {
	var proposal paramscutils.ParamChangeProposalJSON
	var err error

	if paramChangeMsg.JsonFilePath != "" {
		proposal, err = paramscutils.ParseParamChangeProposalJSON(encodingConfig.Amino, paramChangeMsg.JsonFilePath)
		if err != nil {
			return govtypes.MsgSubmitProposal{}, util.LogErr(errors.ErrParse, err)
		}
	} else {
		proposal.Title = paramChangeMsg.Title
		proposal.Description = paramChangeMsg.Description
		proposal.Deposit = paramChangeMsg.Deposit

		var paramChangeJsons paramscutils.ParamChangesJSON
		for _, change := range paramChangeMsg.Changes {
			var targetJson paramscutils.ParamChangeJSON
			if err := encodingConfig.Amino.UnmarshalJSON([]byte(change), &targetJson); err != nil {
				return govtypes.MsgSubmitProposal{}, util.LogErr(errors.ErrFailedToUnmarshal, err)
			}
			paramChangeJsons = append(paramChangeJsons, targetJson)
		}

		proposal.Changes = paramChangeJsons
	}

	deposit, err := sdk.ParseCoinsNormalized(util.DenomAdd(proposal.Deposit))
	if err != nil {
		return govtypes.MsgSubmitProposal{}, util.LogErr(errors.ErrParse, err)
	}

	from, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return govtypes.MsgSubmitProposal{}, util.LogErr(errors.ErrParse, err)
	}
	content := paramsproposal.NewParameterChangeProposal(
		proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
	)

	msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
	if err != nil {
		return govtypes.MsgSubmitProposal{}, util.LogErr(errors.ErrParse, err)
	}

	return *msg, nil
}

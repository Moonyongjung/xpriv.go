package queries

import (
	manchor "github.com/Moonyongjung/xpriv.go/core/anchor"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	anchortypes "github.com/Moonyongjung/xpla-private-chain/x/anchor/types"
)

// Query client for Anchor module.
func (i IXplaClient) QueryAnchor() (string, error) {
	if i.QueryType == types.QueryGrpc {
		return queryByGrpcAnchor(i)
	} else {
		return queryByLcdAnchor(i)
	}

}

func queryByGrpcAnchor(i IXplaClient) (string, error) {
	queryClient := anchortypes.NewQueryClient(i.Ixplac.GetGrpcClient())

	switch {
	// Get anchor account
	case i.Ixplac.GetMsgType() == manchor.AnchorQueryAnchorAccMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorAccountRequest)
		res, err = queryClient.AnchorAccount(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// All aggregated blocks
	case i.Ixplac.GetMsgType() == manchor.AnchorAllAggregatedBlocksMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAllAggregatedBlocksRequest)
		res, err = queryClient.AllAggregatedBlocks(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// anchor info
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorInfoMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorInfoRequest)
		res, err = queryClient.AnchorInfo(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// anchor block
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorBlockMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorBlockRequest)
		res, err = queryClient.AnchorBlock(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// anchor tx body
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorTxBodyMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorTxBodyRequest)
		res, err = queryClient.AnchorTxBody(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// anchor verify
	case i.Ixplac.GetMsgType() == manchor.AnchorVerifyMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryVerifyRequest)
		res, err = queryClient.Verify(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// anchor balances
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorBalancesMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorBalancesRequest)
		res, err = queryClient.AnchorBalances(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err = printProto(i, res)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

const (
	anchorAnchorAccLabel           = "anchor_account"
	anchorAllAggregatedBlocksLabel = "all_aggregated_blocks"
	anchorAnchorTxLabel            = "anchor_tx"
	anchorAnchorBlockLabel         = "anchor_block"
	anchorAnchorTxBodyLabel        = "anchor_tx_body"
	anchorVerifyLabel              = "verify"
	anchorAnchorBalancesLabel      = "anchor_balances"
)

func queryByLcdAnchor(i IXplaClient) (string, error) {
	url := "/xpla/anchor/v1beta1/"

	switch {
	// Get DID
	case i.Ixplac.GetMsgType() == manchor.AnchorQueryAnchorAccMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorAccountRequest)
		url = url + util.MakeQueryLabels(anchorAnchorAccLabel, convertMsg.ValidatorAddress)

		// all aggregated blocks
	case i.Ixplac.GetMsgType() == manchor.AnchorQueryAnchorAccMsgType:
		url = url + util.MakeQueryLabels(anchorAllAggregatedBlocksLabel)

		// anchor info
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorInfoMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorInfoRequest)
		url = url + util.MakeQueryLabels(anchorAnchorTxLabel, util.FromUint64ToString(convertMsg.Height))

		// anchor block
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorBlockMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorBlockRequest)
		url = url + util.MakeQueryLabels(anchorAnchorBlockLabel, util.FromUint64ToString(convertMsg.Height), convertMsg.ContractAddress)

		// anchor tx body
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorTxBodyMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorTxBodyRequest)
		url = url + util.MakeQueryLabels(anchorAnchorBlockLabel, util.FromUint64ToString(convertMsg.Height))

		// verify
	case i.Ixplac.GetMsgType() == manchor.AnchorVerifyMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryVerifyRequest)
		url = url + util.MakeQueryLabels(anchorVerifyLabel, util.FromUint64ToString(convertMsg.Height), convertMsg.ContractAddress)

		// anchor balances
	case i.Ixplac.GetMsgType() == manchor.AnchorAnchorBalancesMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(anchortypes.QueryAnchorBalancesRequest)
		url = url + util.MakeQueryLabels(anchorAnchorBalancesLabel, convertMsg.ValidatorAddress)

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err := util.CtxHttpClient("POST", i.Ixplac.GetLcdURL()+url, i.Ixplac.GetVPByte(), i.Ixplac.GetContext())
	if err != nil {
		return "", err
	}

	return string(out), nil
}

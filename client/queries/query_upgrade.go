package queries

import (
	"context"

	mupgrade "github.com/Moonyongjung/xpla-private-chain.go/core/upgrade"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"

	upgradev1beta1 "cosmossdk.io/api/cosmos/upgrade/v1beta1"
	cmclient "github.com/cosmos/cosmos-sdk/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Query client for upgrade module.
func (i IXplaClient) QueryUpgrade() (string, error) {
	if i.QueryType == types.QueryGrpc {
		return queryByGrpcUpgrade(i)
	} else {
		return queryByLcdUpgrade(i)
	}
}

func queryByGrpcUpgrade(i IXplaClient) (string, error) {
	queryClient := upgradetypes.NewQueryClient(i.Ixplac.GetGrpcClient())

	switch {
	// Upgrade applied
	case i.Ixplac.GetMsgType() == mupgrade.UpgradeAppliedMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(upgradetypes.QueryAppliedPlanRequest)
		appliedPlanRes, err := queryClient.AppliedPlan(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		if appliedPlanRes.Height == 0 {
			return "", util.LogErr(errors.ErrParse, "applied plan height is 0")
		}
		headerData, err := appliedReturnBlockheader(appliedPlanRes, i.Ixplac.GetRpc(), i.Ixplac.GetContext())
		if err != nil {
			return "", err
		}
		return string(headerData), nil

	// Upgrade all module versions
	case i.Ixplac.GetMsgType() == mupgrade.UpgradeQueryAllModuleVersionsMsgType ||
		i.Ixplac.GetMsgType() == mupgrade.UpgradeQueryModuleVersionsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(upgradetypes.QueryModuleVersionsRequest)
		res, err = queryClient.ModuleVersions(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Upgrade plan
	case i.Ixplac.GetMsgType() == mupgrade.UpgradePlanMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(upgradetypes.QueryCurrentPlanRequest)
		res, err = queryClient.CurrentPlan(
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
	upgradeAppliedPlanLabel    = "applied_plan"
	upgradeModuleVersionsLabel = "module_versions"
	upgradeCurrentPlanLabel    = "current_plan"
)

func queryByLcdUpgrade(i IXplaClient) (string, error) {
	url := util.MakeQueryLcdUrl(upgradev1beta1.Query_ServiceDesc.Metadata.(string))

	switch {
	// Upgrade applied
	case i.Ixplac.GetMsgType() == mupgrade.UpgradeAppliedMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(upgradetypes.QueryAppliedPlanRequest)

		url = url + util.MakeQueryLabels(upgradeAppliedPlanLabel, convertMsg.Name)

	// Upgrade all module versions
	case i.Ixplac.GetMsgType() == mupgrade.UpgradeQueryAllModuleVersionsMsgType ||
		i.Ixplac.GetMsgType() == mupgrade.UpgradeQueryModuleVersionsMsgType:

		url = url + upgradeModuleVersionsLabel

	// Upgrade plan
	case i.Ixplac.GetMsgType() == mupgrade.UpgradePlanMsgType:
		url = url + upgradeCurrentPlanLabel

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())

	}

	out, err := util.CtxHttpClient("POST", i.Ixplac.GetLcdURL()+url, i.Ixplac.GetVPByte(), i.Ixplac.GetContext())
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func appliedReturnBlockheader(res *upgradetypes.QueryAppliedPlanResponse, rpcUrl string, ctx context.Context) ([]byte, error) {
	if rpcUrl == "" {
		return nil, util.LogErr(errors.ErrNotSatisfiedOptions, "need RPC URL")
	}
	clientCtx, err := util.NewClient()
	if err != nil {
		return nil, err
	}

	client, err := cmclient.NewClientFromNode(rpcUrl)
	if err != nil {
		return nil, util.LogErr(errors.ErrSdkClient, err)
	}
	clientCtx = clientCtx.WithClient(client)

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, util.LogErr(errors.ErrSdkClient, err)
	}

	headers, err := node.BlockchainInfo(ctx, res.Height, res.Height)
	if err != nil {
		return nil, util.LogErr(errors.ErrSdkClient, err)
	}

	if len(headers.BlockMetas) == 0 {
		return nil, util.LogErr(errors.ErrNotFound, "no headers returns for height", res.Height)
	}

	bytes, err := clientCtx.LegacyAmino.MarshalJSONIndent(headers.BlockMetas[0], "", "  ")
	if err != nil {
		return nil, util.LogErr(errors.ErrFailedToMarshal, err)
	}

	return bytes, nil
}

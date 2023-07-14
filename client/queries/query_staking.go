package queries

import (
	mstaking "github.com/Moonyongjung/xpla-private-chain.go/core/staking"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Query client for staking module.
func (i IXplaClient) QueryStaking() (string, error) {
	if i.QueryType == types.QueryGrpc {
		return queryByGrpcStaking(i)
	} else {
		return queryByLcdStaking(i)
	}
}

func queryByGrpcStaking(i IXplaClient) (string, error) {
	queryClient := stakingtypes.NewQueryClient(i.Ixplac.GetGrpcClient())

	switch {
	// Skating validator
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryValidatorMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorRequest)
		res, err = queryClient.Validator(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking validators
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryValidatorsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorsRequest)
		res, err = queryClient.Validators(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking delegation
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryDelegationMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryDelegationRequest)
		res, err = queryClient.Delegation(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking delegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryDelegationsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryDelegatorDelegationsRequest)
		res, err = queryClient.DelegatorDelegations(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking delegations to
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryDelegationsToMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorDelegationsRequest)
		res, err = queryClient.ValidatorDelegations(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking unbonding delegation
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryUnbondingDelegationMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryUnbondingDelegationRequest)
		res, err = queryClient.UnbondingDelegation(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking unbonding delegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryUnbondingDelegationsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryDelegatorUnbondingDelegationsRequest)
		res, err = queryClient.DelegatorUnbondingDelegations(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking unbonding delegations from
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryUnbondingDelegationsFromMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorUnbondingDelegationsRequest)
		res, err = queryClient.ValidatorUnbondingDelegations(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking redelegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryRedelegationMsgType ||
		i.Ixplac.GetMsgType() == mstaking.StakingQueryRedelegationsMsgType ||
		i.Ixplac.GetMsgType() == mstaking.StakingQueryRedelegationsFromMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryRedelegationsRequest)
		res, err = queryClient.Redelegations(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking historical information
	case i.Ixplac.GetMsgType() == mstaking.StakingHistoricalInfoMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryHistoricalInfoRequest)
		res, err = queryClient.HistoricalInfo(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking pool
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryStakingPoolMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryPoolRequest)
		res, err = queryClient.Pool(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

	// Staking params
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryStakingParamsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryParamsRequest)
		res, err = queryClient.Params(
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
	stakingValidatorsLabel           = "validators"
	stakingDelegatorsLabel           = "delegators"
	stakingDelegationsLabel          = "delegations"
	stakingUnbondingDelegationLabel  = "unbonding_delegation"
	stakingUnbondingDelegationsLabel = "unbonding_delegations"
	stakingRedelegationsLabel        = "redelegations"
	stakingHistoricalInfoLabel       = "historical_info"
	stakingPoolLabel                 = "pool"
	stakingParamsLabel               = "params"
)

func queryByLcdStaking(i IXplaClient) (string, error) {
	url := util.MakeQueryLcdUrl(stakingv1beta1.Query_ServiceDesc.Metadata.(string))

	switch {
	// Skating validator
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryValidatorMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorRequest)

		url = url + util.MakeQueryLabels(stakingValidatorsLabel, convertMsg.ValidatorAddr)

	// Staking validators
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryValidatorsMsgType:
		url = url + stakingValidatorsLabel

	// Staking delegation
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryDelegationMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryDelegationRequest)

		url = url + util.MakeQueryLabels(stakingDelegatorsLabel, convertMsg.DelegatorAddr, stakingValidatorsLabel, convertMsg.ValidatorAddr)

	// Staking delegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryDelegationsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryDelegatorDelegationsRequest)

		url = url + util.MakeQueryLabels(stakingDelegatorsLabel, convertMsg.DelegatorAddr, stakingValidatorsLabel)

	// Staking delegations to
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryDelegationsToMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorDelegationsRequest)

		url = url + util.MakeQueryLabels(stakingValidatorsLabel, convertMsg.ValidatorAddr, stakingDelegationsLabel)

	// Staking unbonding delegation
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryUnbondingDelegationMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryUnbondingDelegationRequest)

		url = url + util.MakeQueryLabels(stakingValidatorsLabel, convertMsg.ValidatorAddr, stakingDelegationsLabel, convertMsg.DelegatorAddr, stakingUnbondingDelegationLabel)

	// Staking unbonding delegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryUnbondingDelegationsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryDelegatorUnbondingDelegationsRequest)

		url = url + util.MakeQueryLabels(stakingDelegatorsLabel, convertMsg.DelegatorAddr, stakingUnbondingDelegationsLabel)

	// Staking unbonding delegations from
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryUnbondingDelegationsFromMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryValidatorUnbondingDelegationsRequest)

		url = url + util.MakeQueryLabels(stakingValidatorsLabel, convertMsg.ValidatorAddr, stakingUnbondingDelegationsLabel)

	// Staking redelegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryRedelegationMsgType ||
		i.Ixplac.GetMsgType() == mstaking.StakingQueryRedelegationsFromMsgType:

		return "", util.LogErr(errors.ErrNotSupport, "unsupported querying delegations by using LCD. query delegations of a delegator")

	// Staking redelegations
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryRedelegationsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryRedelegationsRequest)

		url = url + util.MakeQueryLabels(stakingDelegatorsLabel, convertMsg.DelegatorAddr, stakingRedelegationsLabel)

	// Staking historical information
	case i.Ixplac.GetMsgType() == mstaking.StakingHistoricalInfoMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(stakingtypes.QueryHistoricalInfoRequest)

		url = url + util.MakeQueryLabels(stakingHistoricalInfoLabel, util.FromInt64ToString(convertMsg.Height))

	// Staking pool
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryStakingPoolMsgType:
		url = url + stakingPoolLabel

	// Staking params
	case i.Ixplac.GetMsgType() == mstaking.StakingQueryStakingParamsMsgType:
		url = url + stakingParamsLabel

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err := util.CtxHttpClient("POST", i.Ixplac.GetLcdURL()+url, i.Ixplac.GetVPByte(), i.Ixplac.GetContext())
	if err != nil {
		return "", err
	}

	return string(out), nil
}

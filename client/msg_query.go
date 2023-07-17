package client

import (
	manchor "github.com/Moonyongjung/xpriv.go/core/anchor"
	mauth "github.com/Moonyongjung/xpriv.go/core/auth"
	mbank "github.com/Moonyongjung/xpriv.go/core/bank"
	mbase "github.com/Moonyongjung/xpriv.go/core/base"
	mdid "github.com/Moonyongjung/xpriv.go/core/did"
	mdist "github.com/Moonyongjung/xpriv.go/core/distribution"
	mevidence "github.com/Moonyongjung/xpriv.go/core/evidence"
	mevm "github.com/Moonyongjung/xpriv.go/core/evm"
	mfeegrant "github.com/Moonyongjung/xpriv.go/core/feegrant"
	mgov "github.com/Moonyongjung/xpriv.go/core/gov"
	mmint "github.com/Moonyongjung/xpriv.go/core/mint"
	mparams "github.com/Moonyongjung/xpriv.go/core/params"
	mpriv "github.com/Moonyongjung/xpriv.go/core/private"
	mslashing "github.com/Moonyongjung/xpriv.go/core/slashing"
	mstaking "github.com/Moonyongjung/xpriv.go/core/staking"
	mupgrade "github.com/Moonyongjung/xpriv.go/core/upgrade"
	mwasm "github.com/Moonyongjung/xpriv.go/core/wasm"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"
)

// Anchor module

// Query mapping info of the anchor account.
func (xplac *XplaClient) AnchorAcc(anchorAccMsg types.AnchorAccMsg) *XplaClient {
	msg, err := manchor.MakeAnchorAccMsg(anchorAccMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorQueryAnchorAccMsgType
	xplac.Msg = msg
	return xplac
}

// Query aggregated blocks are saved in the state DB.
func (xplac *XplaClient) AllAggregatedBlocks() *XplaClient {
	msg, err := manchor.MakeAllAggregatedBlocksMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorAllAggregatedBlocksMsgType
	xplac.Msg = msg
	return xplac
}

// Query anchoring info includes from/end height and tx hash are saved in the state DB.
func (xplac *XplaClient) AnchorInfo(anchorInfoMsg types.AnchorInfoMsg) *XplaClient {
	msg, err := manchor.MakeAnchorInfoMsg(anchorInfoMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorAnchorInfoMsgType
	xplac.Msg = msg
	return xplac
}

// Query Anchored block is recorded in the anchor contract of the public chain.
func (xplac *XplaClient) AnchorBlock(anchorBlockMsg types.AnchorBlockMsg) *XplaClient {
	msg, err := manchor.MakeAnchorBlockMsg(anchorBlockMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorAnchorBlockMsgType
	xplac.Msg = msg
	return xplac
}

// Query anchoring transaction body in the public chain.
func (xplac *XplaClient) AnchorTxBody(anchorTxBodyMsg types.AnchorTxBodyMsg) *XplaClient {
	msg, err := manchor.MakeAnchorTxBodyMsg(anchorTxBodyMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorAnchorTxBodyMsgType
	xplac.Msg = msg
	return xplac
}

// Check the consistency of the block in the private chain.
func (xplac *XplaClient) AnchorVerify(anchorVerifyMsg types.AnchorVerifyMsg) *XplaClient {
	msg, err := manchor.MakeAnchorVerifyMsg(anchorVerifyMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorVerifyMsgType
	xplac.Msg = msg
	return xplac
}

// Query balances of the anchor account in the public chain.
func (xplac *XplaClient) AnchorBalances(anchorBalancesMsg types.AnchorBalancesMsg) *XplaClient {
	msg, err := manchor.MakeAnchorBalancesMsg(anchorBalancesMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = manchor.AnchorModule
	xplac.MsgType = manchor.AnchorAnchorBalancesMsgType
	xplac.Msg = msg
	return xplac
}

// Auth module

// Query the current auth parameters.
func (xplac *XplaClient) AuthParams() *XplaClient {
	msg, err := mauth.MakeAuthParamMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mauth.AuthModule
	xplac.MsgType = mauth.AuthQueryParamsMsgType
	xplac.Msg = msg
	return xplac
}

// Query for account by address.
func (xplac *XplaClient) AccAddress(queryAccAddresMsg types.QueryAccAddressMsg) *XplaClient {
	msg, err := mauth.MakeQueryAccAddressMsg(queryAccAddresMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mauth.AuthModule
	xplac.MsgType = mauth.AuthQueryAccAddressMsgType
	xplac.Msg = msg
	return xplac
}

// Query all accounts.
func (xplac *XplaClient) Accounts() *XplaClient {
	msg, err := mauth.MakeQueryAccountsMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mauth.AuthModule
	xplac.MsgType = mauth.AuthQueryAccountsMsgType
	xplac.Msg = msg
	return xplac
}

// Query for paginated transactions that match a set of events.
func (xplac *XplaClient) TxsByEvents(txsByEventsMsg types.QueryTxsByEventsMsg) *XplaClient {
	msg, err := mauth.MakeTxsByEventsMsg(txsByEventsMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mauth.AuthModule
	xplac.MsgType = mauth.AuthQueryTxsByEventsMsgType
	xplac.Msg = msg
	return xplac
}

// Query for a transaction by hash <addr>/<seq> combination or comma-separated signatures in a committed block.
func (xplac *XplaClient) Tx(queryTxMsg types.QueryTxMsg) *XplaClient {
	msg, err := mauth.MakeQueryTxMsg(queryTxMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mauth.AuthModule
	xplac.MsgType = mauth.AuthQueryTxMsgType
	xplac.Msg = msg
	return xplac
}

// Bank module

// Query for account balances by address
func (xplac *XplaClient) BankBalances(bankBalancesMsg types.BankBalancesMsg) *XplaClient {
	if bankBalancesMsg.Denom == "" {
		msg, err := mbank.MakeBankAllBalancesMsg(bankBalancesMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbank.BankModule
		xplac.MsgType = mbank.BankAllBalancesMsgType
		xplac.Msg = msg
	} else {
		msg, err := mbank.MakeBankBalanceMsg(bankBalancesMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbank.BankModule
		xplac.MsgType = mbank.BankBalanceMsgType
		xplac.Msg = msg
	}
	return xplac

}

// Query the client metadata for coin denominations.
func (xplac *XplaClient) DenomMetadata(denomMetadataMsg ...types.DenomMetadataMsg) *XplaClient {
	if len(denomMetadataMsg) == 0 {
		msg, err := mbank.MakeDenomsMetaDataMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbank.BankModule
		xplac.MsgType = mbank.BankDenomsMetadataMsgType
		xplac.Msg = msg
	} else if len(denomMetadataMsg) == 1 {
		msg, err := mbank.MakeDenomMetaDataMsg(denomMetadataMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbank.BankModule
		xplac.MsgType = mbank.BankDenomMetadataMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Query the total supply of coins of the chain.
func (xplac *XplaClient) Total(totalMsg ...types.TotalMsg) *XplaClient {
	if len(totalMsg) == 0 {
		msg, err := mbank.MakeTotalSupplyMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbank.BankModule
		xplac.MsgType = mbank.BankTotalMsgType
		xplac.Msg = msg
	} else if len(totalMsg) == 1 {
		msg, err := mbank.MakeSupplyOfMsg(totalMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbank.BankModule
		xplac.MsgType = mbank.BankTotalSupplyOfMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Base

// Query node info
func (xplac *XplaClient) NodeInfo() *XplaClient {
	msg, err := mbase.MakeBaseNodeInfoMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mbase.Base
	xplac.MsgType = mbase.BaseNodeInfoMsgType
	xplac.Msg = msg
	return xplac
}

// Query syncing
func (xplac *XplaClient) Syncing() *XplaClient {
	msg, err := mbase.MakeBaseSyncingMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mbase.Base
	xplac.MsgType = mbase.BaseSyncingMsgType
	xplac.Msg = msg
	return xplac
}

// Query block
func (xplac *XplaClient) Block(blockMsg ...types.BlockMsg) *XplaClient {
	if len(blockMsg) == 0 {
		msg, err := mbase.MakeBaseLatestBlockMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbase.Base
		xplac.MsgType = mbase.BaseLatestBlockMsgtype
		xplac.Msg = msg
	} else if len(blockMsg) == 1 {
		msg, err := mbase.MakeBaseBlockByHeightMsg(blockMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbase.Base
		xplac.MsgType = mbase.BaseBlockByHeightMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Query validator set
func (xplac *XplaClient) ValidatorSet(validatorSetMsg ...types.ValidatorSetMsg) *XplaClient {
	if len(validatorSetMsg) == 0 {
		msg, err := mbase.MakeLatestValidatorSetMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbase.Base
		xplac.MsgType = mbase.BaseLatestValidatorSetMsgType
		xplac.Msg = msg
	} else if len(validatorSetMsg) == 1 {
		msg, err := mbase.MakeValidatorSetByHeightMsg(validatorSetMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mbase.Base
		xplac.MsgType = mbase.BaseValidatorSetByHeightMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// DID module

// Query DID info.
func (xplac *XplaClient) GetDID(getDIDMsg types.GetDIDMsg) *XplaClient {
	msg, err := mdid.MakeGetDIDMsg(getDIDMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mdid.DidModule
	xplac.MsgType = mdid.DidGetDidMsgType
	xplac.Msg = msg
	return xplac
}

// Distribution module

// Query distribution parameters.
func (xplac *XplaClient) DistributionParams() *XplaClient {
	msg, err := mdist.MakeQueryDistributionParamsMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mdist.DistributionModule
	xplac.MsgType = mdist.DistributionQueryDistributionParamsMsgType
	xplac.Msg = msg
	return xplac
}

// Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.
func (xplac *XplaClient) ValidatorOutstandingRewards(validatorOutstandingRewardsMsg types.ValidatorOutstandingRewardsMsg) *XplaClient {
	msg, err := mdist.MakeValidatorOutstandingRewardsMsg(validatorOutstandingRewardsMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mdist.DistributionModule
	xplac.MsgType = mdist.DistributionValidatorOutstandingRewardsMSgType
	xplac.Msg = msg
	return xplac
}

// Query distribution validator commission.
func (xplac *XplaClient) DistCommission(queryDistCommissionMsg types.QueryDistCommissionMsg) *XplaClient {
	msg, err := mdist.MakeQueryDistCommissionMsg(queryDistCommissionMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mdist.DistributionModule
	xplac.MsgType = mdist.DistributionQueryDistCommissionMsgType
	xplac.Msg = msg
	return xplac
}

// Query distribution validator slashes.
func (xplac *XplaClient) DistSlashes(queryDistSlashesMsg types.QueryDistSlashesMsg) *XplaClient {
	msg, err := mdist.MakeQueryDistSlashesMsg(queryDistSlashesMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mdist.DistributionModule
	xplac.MsgType = mdist.DistributionQuerySlashesMsgType
	xplac.Msg = msg
	return xplac
}

// Query all ditribution delegator rewards or rewards from a particular validator.
func (xplac *XplaClient) DistRewards(queryDistRewardsMsg types.QueryDistRewardsMsg) *XplaClient {
	if queryDistRewardsMsg.DelegatorAddr == "" {
		xplac.Err = util.LogErr(errors.ErrInsufficientParams, "must set a delegator address")
	}

	if queryDistRewardsMsg.ValidatorAddr != "" {
		msg, err := mdist.MakeyQueryDistRewardsMsg(queryDistRewardsMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mdist.DistributionModule
		xplac.MsgType = mdist.DistributionQueryRewardsMsgType
		xplac.Msg = msg
	} else {
		msg, err := mdist.MakeyQueryDistTotalRewardsMsg(queryDistRewardsMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mdist.DistributionModule
		xplac.MsgType = mdist.DistributionQueryTotalRewardsMsgType
		xplac.Msg = msg
	}
	return xplac
}

// Query the amount of coins in the community pool.
func (xplac *XplaClient) CommunityPool() *XplaClient {
	msg, err := mdist.MakeQueryCommunityPoolMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mdist.DistributionModule
	xplac.MsgType = mdist.DistributionQueryCommunityPoolMsgType
	xplac.Msg = msg
	return xplac
}

// Evidence module

// Query for evidence by hash or for all (paginated) submitted evidence.
func (xplac *XplaClient) QueryEvidence(queryEvidenceMsg ...types.QueryEvidenceMsg) *XplaClient {
	if len(queryEvidenceMsg) == 0 {
		msg, err := mevidence.MakeQueryAllEvidenceMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mevidence.EvidenceModule
		xplac.MsgType = mevidence.EvidenceQueryAllMsgType
		xplac.Msg = msg
	} else if len(queryEvidenceMsg) == 1 {
		msg, err := mevidence.MakeQueryEvidenceMsg(queryEvidenceMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mevidence.EvidenceModule
		xplac.MsgType = mevidence.EvidenceQueryMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// EVM module

// Call(as query) solidity contract.
func (xplac *XplaClient) CallSolidityContract(callSolContractMsg types.CallSolContractMsg) *XplaClient {
	msg, err := mevm.MakeCallSolContractMsg(callSolContractMsg, xplac.GetPrivateKey().PubKey().Address().String())
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmCallSolContractMsgType
	xplac.Msg = msg
	return xplac
}

// Query a transaction which is ethereum type information by retrieving hash.
func (xplac *XplaClient) GetTransactionByHash(getTransactionByHashMsg types.GetTransactionByHashMsg) *XplaClient {
	msg, err := mevm.MakeGetTransactionByHashMsg(getTransactionByHashMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmGetTransactionByHashMsgType
	xplac.Msg = msg
	return xplac
}

// Query a block which is ethereum type information by retrieving hash or block height(as number).
func (xplac *XplaClient) GetBlockByHashOrHeight(getBlockByHashHeightMsg types.GetBlockByHashHeightMsg) *XplaClient {
	msg, err := mevm.MakeGetBlockByHashHeightMsg(getBlockByHashHeightMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmGetBlockByHashHeightMsgType
	xplac.Msg = msg
	return xplac
}

// Query a account information which includes account address(hex and bech32), balance and etc.
func (xplac *XplaClient) AccountInfo(accountInfoMsg types.AccountInfoMsg) *XplaClient {
	msg, err := mevm.MakeQueryAccountInfoMsg(accountInfoMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmQueryAccountInfoMsgType
	xplac.Msg = msg
	return xplac
}

// Query suggested gas price.
func (xplac *XplaClient) SuggestGasPrice() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmSuggestGasPriceMsgType
	xplac.Msg = nil
	return xplac
}

// Query chain ID of ethereum type.
func (xplac *XplaClient) EthChainID() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmQueryChainIdMsgType
	xplac.Msg = nil
	return xplac
}

// Query latest block height(as number)
func (xplac *XplaClient) EthBlockNumber() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmQueryCurrentBlockNumberMsgType
	xplac.Msg = nil
	return xplac
}

// Query web3 client version.
func (xplac *XplaClient) Web3ClientVersion() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmWeb3ClientVersionMsgType
	xplac.Msg = nil
	return xplac
}

// Query web3 sha3.
func (xplac *XplaClient) Web3Sha3(web3Sha3Msg types.Web3Sha3Msg) *XplaClient {
	msg, err := mevm.MakeWeb3Sha3Msg(web3Sha3Msg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmWeb3Sha3MsgType
	xplac.Msg = msg
	return xplac
}

// Query current network ID.
func (xplac *XplaClient) NetVersion() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmNetVersionMsgType
	xplac.Msg = nil
	return xplac
}

// Query the number of peers currently connected to the client.
func (xplac *XplaClient) NetPeerCount() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmNetPeerCountMsgType
	xplac.Msg = nil
	return xplac
}

// actively listening for network connections.
func (xplac *XplaClient) NetListening() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmNetListeningMsgType
	xplac.Msg = nil
	return xplac
}

// ethereum protocol version.
func (xplac *XplaClient) EthProtocolVersion() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthProtocolVersionMsgType
	xplac.Msg = nil
	return xplac
}

// Query the sync status object depending on the details of tendermint's sync protocol.
func (xplac *XplaClient) EthSyncing() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthSyncingMsgType
	xplac.Msg = nil
	return xplac
}

// Query all eth accounts.
func (xplac *XplaClient) EthAccounts() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthAccountsMsgType
	xplac.Msg = nil
	return xplac
}

// Query the number of transaction a given block.
func (xplac *XplaClient) EthGetBlockTransactionCount(ethGetBlockTransactionCountMsg types.EthGetBlockTransactionCountMsg) *XplaClient {
	if ethGetBlockTransactionCountMsg.BlockHash == "" && ethGetBlockTransactionCountMsg.BlockHeight == "" {
		xplac.Err = util.LogErr(errors.ErrInsufficientParams, "cannot query, without block hash or height parameter")
	}

	if ethGetBlockTransactionCountMsg.BlockHash != "" && ethGetBlockTransactionCountMsg.BlockHeight != "" {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "select only one parameter, block hash OR height")
	}

	msg, err := mevm.MakeEthGetBlockTransactionCountMsg(ethGetBlockTransactionCountMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthGetBlockTransactionCountMsgType
	xplac.Msg = msg

	return xplac
}

// Query estimate gas.
func (xplac *XplaClient) EstimateGas(invokeSolContractMsg types.InvokeSolContractMsg) *XplaClient {
	msg, err := mevm.MakeEstimateGasSolMsg(invokeSolContractMsg, xplac.GetPrivateKey().PubKey().Address().String())
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthEstimateGasMsgType
	xplac.Msg = msg
	return xplac
}

// Query transaction by block hash and index.
func (xplac *XplaClient) EthGetTransactionByBlockHashAndIndex(getTransactionByBlockHashAndIndexMsg types.GetTransactionByBlockHashAndIndexMsg) *XplaClient {
	msg, err := mevm.MakeGetTransactionByBlockHashAndIndexMsg(getTransactionByBlockHashAndIndexMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmGetTransactionByBlockHashAndIndexMsgType
	xplac.Msg = msg
	return xplac
}

// Query transaction receipt.
func (xplac *XplaClient) EthGetTransactionReceipt(getTransactionReceiptMsg types.GetTransactionReceiptMsg) *XplaClient {
	msg, err := mevm.MakeGetTransactionReceiptMsg(getTransactionReceiptMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmGetTransactionReceiptMsgType
	xplac.Msg = msg
	return xplac
}

// Query filter ID by eth new filter.
func (xplac *XplaClient) EthNewFilter(ethNewFilterMsg types.EthNewFilterMsg) *XplaClient {
	msg, err := mevm.MakeEthNewFilterMsg(ethNewFilterMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthNewFilterMsgType
	xplac.Msg = msg
	return xplac
}

// Query filter ID by eth new block filter.
func (xplac *XplaClient) EthNewBlockFilter() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthNewBlockFilterMsgType
	xplac.Msg = nil
	return xplac
}

// Query filter ID by eth new pending transaction filter.
func (xplac *XplaClient) EthNewPendingTransactionFilter() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthNewPendingTransactionFilterMsgType
	xplac.Msg = nil
	return xplac
}

// Uninstall filter.
func (xplac *XplaClient) EthUninstallFilter(ethUninstallFilterMsg types.EthUninstallFilterMsg) *XplaClient {
	msg, err := mevm.MakeEthUninstallFilterMsg(ethUninstallFilterMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthUninstallFilterMsgType
	xplac.Msg = msg
	return xplac
}

// Query filter changes.
func (xplac *XplaClient) EthGetFilterChanges(ethGetFilterChangesMsg types.EthGetFilterChangesMsg) *XplaClient {
	msg, err := mevm.MakeEthGetFilterChangesMsg(ethGetFilterChangesMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthGetFilterChangesMsgType
	xplac.Msg = msg
	return xplac
}

// Query filter logs.
func (xplac *XplaClient) EthGetFilterLogs(ethGetFilterLogsMsg types.EthGetFilterLogsMsg) *XplaClient {
	msg, err := mevm.MakeEthGetFilterLogsMsg(ethGetFilterLogsMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthGetFilterLogsMsgType
	xplac.Msg = msg
	return xplac
}

// Get logs.
func (xplac *XplaClient) EthGetLogs(ethGetLogsMsg types.EthGetLogsMsg) *XplaClient {
	msg, err := mevm.MakeEthGetLogsMsg(ethGetLogsMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthGetLogsMsgType
	xplac.Msg = msg
	return xplac
}

// Query coinbase.
func (xplac *XplaClient) EthCoinbase() *XplaClient {
	xplac.Module = mevm.EvmModule
	xplac.MsgType = mevm.EvmEthCoinbaseMsgType
	xplac.Msg = nil
	return xplac
}

// Feegrant module

// Query details of fee grants.
func (xplac *XplaClient) QueryFeeGrants(queryFeeGrantMsg types.QueryFeeGrantMsg) *XplaClient {
	if queryFeeGrantMsg.Grantee != "" && queryFeeGrantMsg.Granter != "" {
		msg, err := mfeegrant.MakeQueryFeeGrantMsg(queryFeeGrantMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mfeegrant.FeegrantModule
		xplac.MsgType = mfeegrant.FeegrantQueryGrantMsgType
		xplac.Msg = msg
	} else if queryFeeGrantMsg.Grantee != "" && queryFeeGrantMsg.Granter == "" {
		msg, err := mfeegrant.MakeQueryFeeGrantsByGranteeMsg(queryFeeGrantMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mfeegrant.FeegrantModule
		xplac.MsgType = mfeegrant.FeegrantQueryGrantsByGranteeMsgType
		xplac.Msg = msg
	} else if queryFeeGrantMsg.Grantee == "" && queryFeeGrantMsg.Granter != "" {
		msg, err := mfeegrant.MakeQueryFeeGrantsByGranterMsg(queryFeeGrantMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mfeegrant.FeegrantModule
		xplac.MsgType = mfeegrant.FeegrantQueryGrantsByGranterMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInsufficientParams, "no query grants parameters")
	}

	return xplac
}

// Gov module

// Query details of a singla proposal.
func (xplac *XplaClient) QueryProposal(queryProposal types.QueryProposalMsg) *XplaClient {
	msg, err := mgov.MakeQueryProposalMsg(queryProposal)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mgov.GovModule
	xplac.MsgType = mgov.GovQueryProposalMsgType
	xplac.Msg = msg
	return xplac
}

// Query proposals with optional filters.
func (xplac *XplaClient) QueryProposals(queryProposals types.QueryProposalsMsg) *XplaClient {
	msg, err := mgov.MakeQueryProposalsMsg(queryProposals)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mgov.GovModule
	xplac.MsgType = mgov.GovQueryProposalsMsgType
	xplac.Msg = msg
	return xplac
}

// Query details of a deposit or deposits on a proposal.
func (xplac *XplaClient) QueryDeposit(queryDepositMsg types.QueryDepositMsg) *XplaClient {
	var queryType int
	if xplac.Opts.GrpcURL != "" {
		queryType = types.QueryGrpc
	} else {
		queryType = types.QueryLcd
	}

	if queryDepositMsg.Depositor != "" {
		msg, argsType, err := mgov.MakeQueryDepositMsg(queryDepositMsg, xplac.Grpc, xplac.Context, xplac.Opts.LcdURL, queryType)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		if argsType == "params" {
			xplac.Module = mgov.GovModule
			xplac.MsgType = mgov.GovQueryDepositParamsMsgType
			xplac.Msg = msg
		} else {
			xplac.Module = mgov.GovModule
			xplac.MsgType = mgov.GovQueryDepositRequestMsgType
			xplac.Msg = msg
		}
	} else {
		msg, argsType, err := mgov.MakeQueryDepositsMsg(queryDepositMsg, xplac.Grpc, xplac.Context, xplac.Opts.LcdURL, queryType)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		if argsType == "params" {
			xplac.Module = mgov.GovModule
			xplac.MsgType = mgov.GovQueryDepositsParamsMsgType
			xplac.Msg = msg
		} else {
			xplac.Module = mgov.GovModule
			xplac.MsgType = mgov.GovQueryDepositsRequestMsgType
			xplac.Msg = msg
		}
	}
	return xplac
}

// Query details of a single vote or votes on a proposal.
func (xplac *XplaClient) QueryVote(queryVoteMsg types.QueryVoteMsg) *XplaClient {
	var queryType int
	if xplac.Opts.GrpcURL != "" {
		queryType = types.QueryGrpc
	} else {
		queryType = types.QueryLcd
	}

	if queryVoteMsg.VoterAddr != "" {
		msg, err := mgov.MakeQueryVoteMsg(queryVoteMsg, xplac.Grpc, xplac.Context, xplac.Opts.LcdURL, queryType)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mgov.GovModule
		xplac.MsgType = mgov.GovQueryVoteMsgType
		xplac.Msg = msg

	} else {
		msg, status, err := mgov.MakeQueryVotesMsg(queryVoteMsg, xplac.Grpc, xplac.Context, xplac.Opts.LcdURL, queryType)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		if status == "notPassed" {
			xplac.Module = mgov.GovModule
			xplac.MsgType = mgov.GovQueryVotesNotPassedMsgType
			xplac.Msg = msg
		} else {
			xplac.Module = mgov.GovModule
			xplac.MsgType = mgov.GovQueryVotesPassedMsgType
			xplac.Msg = msg
		}
	}
	return xplac
}

// Query the tally of a proposal vote.
func (xplac *XplaClient) Tally(tallyMsg types.TallyMsg) *XplaClient {
	var queryType int
	if xplac.Opts.GrpcURL != "" {
		queryType = types.QueryGrpc
	} else {
		queryType = types.QueryLcd
	}

	msg, err := mgov.MakeGovTallyMsg(tallyMsg, xplac.Grpc, xplac.Context, xplac.Opts.LcdURL, queryType)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mgov.GovModule
	xplac.MsgType = mgov.GovTallyMsgType
	xplac.Msg = msg
	return xplac
}

// Query parameters of the governance process or the parameters (voting|tallying|deposit) of the governance process.
func (xplac *XplaClient) GovParams(govParamsMsg ...types.GovParamsMsg) *XplaClient {
	if len(govParamsMsg) == 0 {
		xplac.Module = mgov.GovModule
		xplac.MsgType = mgov.GovQueryGovParamsMsgType
		xplac.Msg = nil
	} else if len(govParamsMsg) == 1 {
		msg, err := mgov.MakeGovParamsMsg(govParamsMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mgov.GovModule
		switch govParamsMsg[0].ParamType {
		case "voting":
			xplac.MsgType = mgov.GovQueryGovParamVotingMsgType
		case "tallying":
			xplac.MsgType = mgov.GovQueryGovParamTallyingMsgType
		case "deposit":
			xplac.MsgType = mgov.GovQueryGovParamDepositMsgType
		}
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Query the proposer of a governance proposal.
func (xplac *XplaClient) Proposer(proposerMsg types.ProposerMsg) *XplaClient {
	xplac.Module = mgov.GovModule
	xplac.MsgType = mgov.GovQueryProposerMsgType
	xplac.Msg = proposerMsg.ProposalID
	return xplac
}

// Mint module

// Query the current minting parameters.
func (xplac *XplaClient) MintParams() *XplaClient {
	msg, err := mmint.MakeQueryMintParamsMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mmint.MintModule
	xplac.MsgType = mmint.MintQueryMintParamsMsgType
	xplac.Msg = msg
	return xplac
}

// Query the current minting inflation value.
func (xplac *XplaClient) Inflation() *XplaClient {
	msg, err := mmint.MakeQueryInflationMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mmint.MintModule
	xplac.MsgType = mmint.MintQueryInflationMsgType
	xplac.Msg = msg
	return xplac
}

// Query the current minting annual provisions value.
func (xplac *XplaClient) AnnualProvisions() *XplaClient {
	msg, err := mmint.MakeQueryAnnualProvisionsMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mmint.MintModule
	xplac.MsgType = mmint.MintQueryAnnualProvisionsMsgType
	xplac.Msg = msg
	return xplac
}

// Params module

// Query for raw parameters by subspace and key.
func (xplac *XplaClient) QuerySubspace(subspaceMsg types.SubspaceMsg) *XplaClient {
	msg, err := mparams.MakeQueryParamsSubspaceMsg(subspaceMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mparams.ParamsModule
	xplac.MsgType = mparams.ParamsQuerySubpsaceMsgType
	xplac.Msg = msg

	return xplac
}

// Private module

// Query all administrators in the private chain.
func (xplac *XplaClient) Admin() *XplaClient {
	msg, err := mpriv.MakeQueryAdminMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mpriv.PrivateModule
	xplac.MsgType = mpriv.PrivateQueryAdminMsgType
	xplac.Msg = msg

	return xplac
}

// Query participate state of the DID.
func (xplac *XplaClient) ParticipateState(participateStateMsg types.ParticipateStateMsg) *XplaClient {
	msg, err := mpriv.MakeParticipateStateMsg(participateStateMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mpriv.PrivateModule
	xplac.MsgType = mpriv.PrivateParticipateStateMsgType
	xplac.Msg = msg

	return xplac
}

// Gen the DID signature from DID key.
func (xplac *XplaClient) GenDIDSign(genDIDSignMsg types.GenDIDSignMsg) *XplaClient {
	msg, err := mpriv.MakeGenDIDSignMsg(genDIDSignMsg, xplac.GetLcdURL(), xplac.GetGrpcUrl(), xplac.GetGrpcClient(), xplac.GetContext())
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mpriv.PrivateModule
	xplac.MsgType = mpriv.PrivateGenDIDSignMsgType
	xplac.Msg = msg

	return xplac
}

// Get the verifiable credential which is able to use on the private chain for access control.
// Only the VC owner can query by using DID signature
func (xplac *XplaClient) IssueVC(issueVCMsg types.IssueVCMsg) *XplaClient {
	msg, err := mpriv.MakeIssueVCMsg(issueVCMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mpriv.PrivateModule
	xplac.MsgType = mpriv.PrivateIssueVCMsgType
	xplac.Msg = msg

	return xplac
}

// Get the verifiable presentation which is able to use on the private chain for access control.
// Participants can acess the private chain when this presentation submit to chain as verifier.
// Only the VP owner can query by using DID signature
func (xplac *XplaClient) GetVP(getVPMsg types.GetVPMsg) *XplaClient {
	msg, err := mpriv.MakeGetVPMsg(getVPMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mpriv.PrivateModule
	xplac.MsgType = mpriv.PrivateGetVPMsgType
	xplac.Msg = msg

	return xplac
}

// Slashing module

// Query the current slashing parameters.
func (xplac *XplaClient) SlashingParams() *XplaClient {
	msg, err := mslashing.MakeQuerySlashingParamsMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mslashing.SlashingModule
	xplac.MsgType = mslashing.SlahsingQuerySlashingParamsMsgType
	xplac.Msg = msg
	return xplac
}

// Query a validator's signing information or signing information of all validators.
func (xplac *XplaClient) SigningInfos(signingInfoMsg ...types.SigningInfoMsg) *XplaClient {
	if len(signingInfoMsg) == 0 {
		msg, err := mslashing.MakeQuerySigningInfosMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mslashing.SlashingModule
		xplac.MsgType = mslashing.SlashingQuerySigningInfosMsgType
		xplac.Msg = msg
	} else if len(signingInfoMsg) == 1 {
		msg, err := mslashing.MakeQuerySigningInfoMsg(signingInfoMsg[0], xplac.EncodingConfig)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mslashing.SlashingModule
		xplac.MsgType = mslashing.SlashingQuerySigningInfoMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Staking module

// Query a validator or for all validators.
func (xplac *XplaClient) QueryValidators(queryValidatorMsg ...types.QueryValidatorMsg) *XplaClient {
	if len(queryValidatorMsg) == 0 {
		msg, err := mstaking.MakeQueryValidatorsMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryValidatorsMsgType
		xplac.Msg = msg
	} else if len(queryValidatorMsg) == 1 {
		msg, err := mstaking.MakeQueryValidatorMsg(queryValidatorMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryValidatorMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Query a delegation based on address and validator address, all out going redelegations from a validator or all delegations made by on delegator.
func (xplac *XplaClient) QueryDelegation(queryDelegationMsg types.QueryDelegationMsg) *XplaClient {
	if queryDelegationMsg.DelegatorAddr != "" && queryDelegationMsg.ValidatorAddr != "" {
		msg, err := mstaking.MakeQueryDelegationMsg(queryDelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryDelegationMsgType
		xplac.Msg = msg
	} else if queryDelegationMsg.DelegatorAddr != "" {
		msg, err := mstaking.MakeQueryDelegationsMsg(queryDelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryDelegationsMsgType
		xplac.Msg = msg
	} else if queryDelegationMsg.ValidatorAddr != "" {
		msg, err := mstaking.MakeQueryDelegationsToMsg(queryDelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryDelegationsToMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "wrong delegation message")
	}
	return xplac
}

// Query all unbonding delegatations from a validator, an unbonding-delegation record based on delegator and validator address or all unbonding-delegations records for one delegator.
func (xplac *XplaClient) QueryUnbondingDelegation(queryUnbondingDelegationMsg types.QueryUnbondingDelegationMsg) *XplaClient {
	if queryUnbondingDelegationMsg.DelegatorAddr != "" && queryUnbondingDelegationMsg.ValidatorAddr != "" {
		msg, err := mstaking.MakeQueryUnbondingDelegationMsg(queryUnbondingDelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryUnbondingDelegationMsgType
		xplac.Msg = msg
	} else if queryUnbondingDelegationMsg.DelegatorAddr != "" {
		msg, err := mstaking.MakeQueryUnbondingDelegationsMsg(queryUnbondingDelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryUnbondingDelegationsMsgType
		xplac.Msg = msg
	} else if queryUnbondingDelegationMsg.ValidatorAddr != "" {
		msg, err := mstaking.MakeQueryUnbondingDelegationsFromMsg(queryUnbondingDelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryUnbondingDelegationsFromMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "wrong unbonding delegation message")
	}
	return xplac
}

// Query a redelegation record based on delegator and a source and destination validator.
// Also, query all outgoing redelegatations from a validator or all redelegations records for one delegator.
func (xplac *XplaClient) QueryRedelegation(queryRedelegationMsg types.QueryRedelegationMsg) *XplaClient {
	if queryRedelegationMsg.DelegatorAddr != "" &&
		queryRedelegationMsg.SrcValidatorAddr != "" &&
		queryRedelegationMsg.DstValidatorAddr != "" {
		msg, err := mstaking.MakeQueryRedelegationMsg(queryRedelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryRedelegationMsgType
		xplac.Msg = msg
	} else if queryRedelegationMsg.DelegatorAddr != "" {
		msg, err := mstaking.MakeQueryRedelegationsMsg(queryRedelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryRedelegationsMsgType
		xplac.Msg = msg
	} else if queryRedelegationMsg.SrcValidatorAddr != "" {
		msg, err := mstaking.MakeQueryRedelegationsFromMsg(queryRedelegationMsg)
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mstaking.StakingModule
		xplac.MsgType = mstaking.StakingQueryRedelegationsFromMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "wrong redelegation message")
	}
	return xplac
}

// Query historical info at given height.
func (xplac *XplaClient) HistoricalInfo(historicalInfoMsg types.HistoricalInfoMsg) *XplaClient {
	msg, err := mstaking.MakeHistoricalInfoMsg(historicalInfoMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mstaking.StakingModule
	xplac.MsgType = mstaking.StakingHistoricalInfoMsgType
	xplac.Msg = msg
	return xplac
}

// Query the current staking pool values.
func (xplac *XplaClient) StakingPool() *XplaClient {
	msg, err := mstaking.MakeQueryStakingPoolMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mstaking.StakingModule
	xplac.MsgType = mstaking.StakingQueryStakingPoolMsgType
	xplac.Msg = msg
	return xplac
}

// Query the current staking parameters information.
func (xplac *XplaClient) StakingParams() *XplaClient {
	msg, err := mstaking.MakeQueryStakingParamsMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mstaking.StakingModule
	xplac.MsgType = mstaking.StakingQueryStakingParamsMsgType
	xplac.Msg = msg
	return xplac
}

// Upgrade module

// Block header for height at which a completed upgrade was applied.
func (xplac *XplaClient) UpgradeApplied(appliedMsg types.AppliedMsg) *XplaClient {
	msg, err := mupgrade.MakeAppliedMsg(appliedMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mupgrade.UpgradeModule
	xplac.MsgType = mupgrade.UpgradeAppliedMsgType
	xplac.Msg = msg
	return xplac
}

// Query the list of module versions.
func (xplac *XplaClient) ModulesVersion(queryModulesVersionMsg ...types.QueryModulesVersionMsg) *XplaClient {
	if len(queryModulesVersionMsg) == 0 {
		msg, err := mupgrade.MakeQueryAllModuleVersionMsg()
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mupgrade.UpgradeModule
		xplac.MsgType = mupgrade.UpgradeQueryAllModuleVersionsMsgType
		xplac.Msg = msg
	} else if len(queryModulesVersionMsg) == 1 {
		msg, err := mupgrade.MakeQueryModuleVersionMsg(queryModulesVersionMsg[0])
		if err != nil {
			xplac.Err = err
			return xplac
		}
		xplac.Module = mupgrade.UpgradeModule
		xplac.MsgType = mupgrade.UpgradeQueryModuleVersionsMsgType
		xplac.Msg = msg
	} else {
		xplac.Err = util.LogErr(errors.ErrInvalidRequest, "need only one parameter")
	}
	return xplac
}

// Query upgrade plan(if one exists).
func (xplac *XplaClient) Plan() *XplaClient {
	msg, err := mupgrade.MakePlanMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mupgrade.UpgradeModule
	xplac.MsgType = mupgrade.UpgradePlanMsgType
	xplac.Msg = msg
	return xplac
}

// Wasm module

// Calls contract with given address with query data and prints the returned result.
func (xplac *XplaClient) QueryContract(queryMsg types.QueryMsg) *XplaClient {
	msg, err := mwasm.MakeQueryMsg(queryMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmQueryContractMsgType
	xplac.Msg = msg
	return xplac
}

// Query list all wasm bytecode on the chain.
func (xplac *XplaClient) ListCode() *XplaClient {
	msg, err := mwasm.MakeListcodeMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmListCodeMsgType
	xplac.Msg = msg
	return xplac
}

// Query list wasm all bytecode on the chain for given code ID.
func (xplac *XplaClient) ListContractByCode(listContractByCodeMsg types.ListContractByCodeMsg) *XplaClient {
	msg, err := mwasm.MakeListContractByCodeMsg(listContractByCodeMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmListContractByCodeMsgType
	xplac.Msg = msg
	return xplac
}

// Downloads wasm bytecode for given code ID.
func (xplac *XplaClient) Download(downloadMsg types.DownloadMsg) *XplaClient {
	msg, err := mwasm.MakeDownloadMsg(downloadMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmDownloadMsgType
	xplac.Msg = msg
	return xplac
}

// Prints out metadata of a code ID.
func (xplac *XplaClient) CodeInfo(codeInfoMsg types.CodeInfoMsg) *XplaClient {
	msg, err := mwasm.MakeCodeInfoMsg(codeInfoMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmCodeInfoMsgType
	xplac.Msg = msg
	return xplac
}

// Prints out metadata of a contract given its address.
func (xplac *XplaClient) ContractInfo(contractInfoMsg types.ContractInfoMsg) *XplaClient {
	msg, err := mwasm.MakeContractInfoMsg(contractInfoMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmContractInfoMsgType
	xplac.Msg = msg
	return xplac
}

// Prints out all internal state of a contract given its address.
func (xplac *XplaClient) ContractStateAll(contractStateAllMsg types.ContractStateAllMsg) *XplaClient {
	msg, err := mwasm.MakeContractStateAllMsg(contractStateAllMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmContractStateAllMsgType
	xplac.Msg = msg
	return xplac
}

// Prints out the code history for a contract given its address.
func (xplac *XplaClient) ContractHistory(contractHistoryMsg types.ContractHistoryMsg) *XplaClient {
	msg, err := mwasm.MakeContractHistoryMsg(contractHistoryMsg)
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmContractHistoryMsgType
	xplac.Msg = msg
	return xplac
}

// Query list all pinned code IDs.
func (xplac *XplaClient) Pinned() *XplaClient {
	msg, err := mwasm.MakePinnedMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmPinnedMsgType
	xplac.Msg = msg
	return xplac
}

// Get libwasmvm version.
func (xplac *XplaClient) LibwasmvmVersion() *XplaClient {
	msg, err := mwasm.MakeLibwasmvmVersionMsg()
	if err != nil {
		xplac.Err = err
		return xplac
	}
	xplac.Module = mwasm.WasmModule
	xplac.MsgType = mwasm.WasmLibwasmvmVersionMsgType
	xplac.Msg = msg
	return xplac
}

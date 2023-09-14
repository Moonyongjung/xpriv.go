package provider

import (
	"context"

	"github.com/Moonyongjung/xpla-private-chain/app/params"
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"

	cmclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/grpc"
)

// The standard form of XPLA client is interface type.
// XplaClient is endpoint in order to access xpla.go from external packages.
// If new modules are implemeted, external functions that are used to send tx or query state should be
// enrolled in XplaClient interface.
//
// e.g. - enroll bank module
//
//	  type TxMsgProvider interface {
//		...
//		BankSend(types.BankSendMsg) XplaClient
//		...
//	  }
//
//	  type QueryMsgProvider interface {
//		...
//		BankBalances(types.BankBalancesMsg) XplaClient
//		DenomMetadata(...types.DenomMetadataMsg) XplaClient
//		Total(...types.TotalMsg) XplaClient
//		...
//	  }
//
// The return type of these methods must be always the XplaClient because the client uses mehod chaining.
//
// e.g. - create and sign transaction
//
//	txbytes, err := xplac.BankSend(bankSendMsg).CreateAndSignTx()
//
// e.g. - query
//
//	res, err := xplac.BankBalances(bankBalancesMsg).Query()
type XplaClient interface {
	WithProvider
	GetProvider
	TxProvider
	QueryProvider
	BroadcastProvider
	InfoRequestProvider
	TxMsgProvider
	QueryMsgProvider
	HelperProvider
}

// Optional parameters of client.xplaClient.
type Options struct {
	PrivateKey     key.PrivateKey
	AccountNumber  string
	Sequence       string
	BroadcastMode  string
	GasLimit       string
	GasPrice       string
	GasAdjustment  string
	FeeAmount      string
	SignMode       signing.SignMode
	FeeGranter     sdk.AccAddress
	TimeoutHeight  string
	LcdURL         string
	GrpcURL        string
	RpcURL         string
	EvmRpcURL      string
	Pagination     types.Pagination
	OutputDocument string
	UseVP          bool
	VPPath         string
	VPString       string
}

// Methods set params of client.xplaClient.
type WithProvider interface {
	UpdateXplacInCoreModule() XplaClient
	WithOptions(Options) XplaClient
	WithChainId(string) XplaClient
	WithEncoding(params.EncodingConfig) XplaClient
	WithContext(context.Context) XplaClient
	WithPrivateKey(key.PrivateKey) XplaClient
	WithAccountNumber(string) XplaClient
	WithBroadcastMode(string) XplaClient
	WithSequence(string) XplaClient
	WithGasLimit(string) XplaClient
	WithGasPrice(string) XplaClient
	WithGasAdjustment(string) XplaClient
	WithFeeAmount(string) XplaClient
	WithSignMode(signing.SignMode) XplaClient
	WithFeeGranter(sdk.AccAddress) XplaClient
	WithTimeoutHeight(string) XplaClient
	WithURL(string) XplaClient
	WithGrpc(string) XplaClient
	WithRpc(string) XplaClient
	WithEvmRpc(string) XplaClient
	WithPagination(types.Pagination) XplaClient
	WithOutputDocument(string) XplaClient
	WithModule(string) XplaClient
	WithMsgType(string) XplaClient
	WithMsg(interface{}) XplaClient
	WithErr(error) XplaClient
	WithUseVP(bool) XplaClient
	WithVPByPath(string) XplaClient
	WithVPByString(string) XplaClient
}

// Methods get params of client.xplaClient.
type GetProvider interface {
	GetChainId() string
	GetPrivateKey() key.PrivateKey
	GetEncoding() params.EncodingConfig
	GetContext() context.Context
	GetLcdURL() string
	GetGrpcUrl() string
	GetGrpcClient() grpc.ClientConn
	GetRpc() string
	GetEvmRpc() string
	GetBroadcastMode() string
	GetAccountNumber() string
	GetSequence() string
	GetGasLimit() string
	GetGasPrice() string
	GetGasAdjustment() string
	GetFeeAmount() string
	GetSignMode() signing.SignMode
	GetFeeGranter() sdk.AccAddress
	GetTimeoutHeight() string
	GetPagination() *query.PageRequest
	GetOutputDocument() string
	GetModule() string
	GetMsg() interface{}
	GetMsgType() string
	GetErr() error
	GetUseVP() bool
	GetVPByte() []byte
}

// Methods handle transaction.
type TxProvider interface {
	CreateAndSignTx() ([]byte, error)
	CreateUnsignedTx() ([]byte, error)
	SignTx(types.SignTxMsg) ([]byte, error)
	MultiSign(types.TxMultiSignMsg) ([]byte, error)
	EncodeTx(types.EncodeTxMsg) (string, error)
	DecodeTx(types.DecodeTxMsg) (string, error)
	ValidateSignatures(types.ValidateSignaturesMsg) (string, error)
}

// Method handles query functions.
type QueryProvider interface {
	Query() (string, error)
}

// Methods handle functions of broadcasting.
type BroadcastProvider interface {
	Broadcast([]byte) (*types.TxRes, error)
	BroadcastBlock([]byte) (*types.TxRes, error)
	BroadcastAsync([]byte) (*types.TxRes, error)
}

// Methods get information from XPLA chain.
type InfoRequestProvider interface {
	LoadAccount(sdk.AccAddress) (authtypes.AccountI, error)
	Simulate(cmclient.TxBuilder) (*sdktx.SimulateResponse, error)
}

// Methods are external functions of each module for sending transaction.
type TxMsgProvider interface {
	// anchor
	RegisterAnchorAcc(types.RegisterAnchorAccMsg) XplaClient
	ChangeAnchorAcc(types.ChangeAnchorAccMsg) XplaClient

	// bank
	BankSend(types.BankSendMsg) XplaClient

	// crisis
	InvariantBroken(types.InvariantBrokenMsg) XplaClient

	// did
	CreateDID(types.CreateDIDMsg) XplaClient
	UpdateDID(types.UpdateDIDMsg) XplaClient
	DeactivateDID(types.DeactivateDIDMsg) XplaClient
	ReplaceDIDMoniker(types.ReplaceDIDMonikerMsg) XplaClient

	// distribution
	FundCommunityPool(types.FundCommunityPoolMsg) XplaClient
	CommunityPoolSpend(types.CommunityPoolSpendMsg) XplaClient
	WithdrawRewards(types.WithdrawRewardsMsg) XplaClient
	WithdrawAllRewards() XplaClient
	SetWithdrawAddr(types.SetWithdrawAddrMsg) XplaClient

	// evm
	EvmSendCoin(types.SendCoinMsg) XplaClient
	DeploySolidityContract(types.DeploySolContractMsg) XplaClient
	InvokeSolidityContract(types.InvokeSolContractMsg) XplaClient

	// feegrant
	FeeGrant(types.FeeGrantMsg) XplaClient
	RevokeFeeGrant(types.RevokeFeeGrantMsg) XplaClient

	// gov
	SubmitProposal(types.SubmitProposalMsg) XplaClient
	GovDeposit(types.GovDepositMsg) XplaClient
	Vote(types.VoteMsg) XplaClient
	WeightedVote(types.WeightedVoteMsg) XplaClient

	// params
	ParamChange(types.ParamChangeMsg) XplaClient

	// private
	InitialAdmin(types.InitialAdminMsg) XplaClient
	AddAdmin(types.AddAdminMsg) XplaClient
	Participate(types.ParticipateMsg) XplaClient
	Accept(types.AcceptMsg) XplaClient
	Deny(types.DenyMsg) XplaClient
	Exile(types.ExileMsg) XplaClient
	Quit(types.QuitMsg) XplaClient

	// slashing
	Unjail() XplaClient

	// staking
	CreateValidator(types.CreateValidatorMsg) XplaClient
	EditValidator(types.EditValidatorMsg) XplaClient
	Delegate(types.DelegateMsg) XplaClient
	Unbond(types.UnbondMsg) XplaClient
	Redelegate(types.RedelegateMsg) XplaClient

	// upgrade
	SoftwareUpgrade(types.SoftwareUpgradeMsg) XplaClient
	CancelSoftwareUpgrade(types.CancelSoftwareUpgradeMsg) XplaClient

	// wasm
	StoreCode(types.StoreMsg) XplaClient
	InstantiateContract(types.InstantiateMsg) XplaClient
	ExecuteContract(types.ExecuteMsg) XplaClient
	ClearContractAdmin(types.ClearContractAdminMsg) XplaClient
	SetContractAdmin(types.SetContractAdminMsg) XplaClient
	Migrate(types.MigrateMsg) XplaClient
}

// Methods are external functions of each module for querying.
type QueryMsgProvider interface {
	// anchor
	AnchorAcc(types.AnchorAccMsg) XplaClient
	AllAggregatedBlocks() XplaClient
	AnchorInfo(types.AnchorInfoMsg) XplaClient
	AnchorBlock(types.AnchorBlockMsg) XplaClient
	AnchorTxBody(types.AnchorTxBodyMsg) XplaClient
	AnchorVerify(types.AnchorVerifyMsg) XplaClient
	AnchorBalances(types.AnchorBalancesMsg) XplaClient
	AnchorParams() XplaClient

	// auth
	AuthParams() XplaClient
	AccAddress(types.QueryAccAddressMsg) XplaClient
	Accounts() XplaClient
	TxsByEvents(types.QueryTxsByEventsMsg) XplaClient
	Tx(types.QueryTxMsg) XplaClient

	// bank
	BankBalances(types.BankBalancesMsg) XplaClient
	DenomMetadata(...types.DenomMetadataMsg) XplaClient
	Total(...types.TotalMsg) XplaClient

	// base
	NodeInfo() XplaClient
	Syncing() XplaClient
	Block(...types.BlockMsg) XplaClient
	ValidatorSet(...types.ValidatorSetMsg) XplaClient

	// did
	GetDID(types.GetDIDMsg) XplaClient
	MonikerByDID(types.MonikerByDIDMsg) XplaClient
	DIDByMoniker(types.DIDByMonikerMsg) XplaClient
	AllDIDs() XplaClient

	// distribution
	DistributionParams() XplaClient
	ValidatorOutstandingRewards(types.ValidatorOutstandingRewardsMsg) XplaClient
	DistCommission(types.QueryDistCommissionMsg) XplaClient
	DistSlashes(types.QueryDistSlashesMsg) XplaClient
	DistRewards(types.QueryDistRewardsMsg) XplaClient
	CommunityPool() XplaClient

	// evidence
	QueryEvidence(...types.QueryEvidenceMsg) XplaClient

	// evm
	CallSolidityContract(types.CallSolContractMsg) XplaClient
	GetTransactionByHash(types.GetTransactionByHashMsg) XplaClient
	GetBlockByHashOrHeight(types.GetBlockByHashHeightMsg) XplaClient
	AccountInfo(types.AccountInfoMsg) XplaClient
	SuggestGasPrice() XplaClient
	EthChainID() XplaClient
	EthBlockNumber() XplaClient
	Web3ClientVersion() XplaClient
	Web3Sha3(types.Web3Sha3Msg) XplaClient
	NetVersion() XplaClient
	NetPeerCount() XplaClient
	NetListening() XplaClient
	EthProtocolVersion() XplaClient
	EthSyncing() XplaClient
	EthAccounts() XplaClient
	EthGetBlockTransactionCount(types.EthGetBlockTransactionCountMsg) XplaClient
	EstimateGas(types.InvokeSolContractMsg) XplaClient
	EthGetTransactionByBlockHashAndIndex(types.GetTransactionByBlockHashAndIndexMsg) XplaClient
	EthGetTransactionReceipt(types.GetTransactionReceiptMsg) XplaClient
	EthNewFilter(types.EthNewFilterMsg) XplaClient
	EthNewBlockFilter() XplaClient
	EthNewPendingTransactionFilter() XplaClient
	EthUninstallFilter(types.EthUninstallFilterMsg) XplaClient
	EthGetFilterChanges(types.EthGetFilterChangesMsg) XplaClient
	EthGetFilterLogs(types.EthGetFilterLogsMsg) XplaClient
	EthGetLogs(types.EthGetLogsMsg) XplaClient
	EthCoinbase() XplaClient

	// feegrant
	QueryFeeGrants(types.QueryFeeGrantMsg) XplaClient

	// gov
	QueryProposal(types.QueryProposalMsg) XplaClient
	QueryProposals(types.QueryProposalsMsg) XplaClient
	QueryDeposit(types.QueryDepositMsg) XplaClient
	QueryVote(types.QueryVoteMsg) XplaClient
	Tally(types.TallyMsg) XplaClient
	GovParams(...types.GovParamsMsg) XplaClient
	Proposer(types.ProposerMsg) XplaClient

	// mint
	MintParams() XplaClient
	Inflation() XplaClient
	AnnualProvisions() XplaClient

	// params
	QuerySubspace(types.SubspaceMsg) XplaClient

	// private
	Admin() XplaClient
	ParticipateState(types.ParticipateStateMsg) XplaClient
	ParticipateSequence(types.ParticipateSequenceMsg) XplaClient
	GenDIDSign(types.GenDIDSignMsg) XplaClient
	IssueVC(types.IssueVCMsg) XplaClient
	GetVP(types.GetVPMsg) XplaClient
	AllUnderReviews() XplaClient
	AllParticipants() XplaClient

	// slashing
	SlashingParams() XplaClient
	SigningInfos(...types.SigningInfoMsg) XplaClient

	// staking
	QueryValidators(...types.QueryValidatorMsg) XplaClient
	QueryDelegation(types.QueryDelegationMsg) XplaClient
	QueryUnbondingDelegation(types.QueryUnbondingDelegationMsg) XplaClient
	QueryRedelegation(types.QueryRedelegationMsg) XplaClient
	HistoricalInfo(types.HistoricalInfoMsg) XplaClient
	StakingPool() XplaClient
	StakingParams() XplaClient

	// upgrade
	UpgradeApplied(types.AppliedMsg) XplaClient
	ModulesVersion(...types.QueryModulesVersionMsg) XplaClient
	Plan() XplaClient

	// wasm
	QueryContract(types.QueryMsg) XplaClient
	ListCode() XplaClient
	ListContractByCode(types.ListContractByCodeMsg) XplaClient
	Download(types.DownloadMsg) XplaClient
	CodeInfo(types.CodeInfoMsg) XplaClient
	ContractInfo(types.ContractInfoMsg) XplaClient
	ContractStateAll(types.ContractStateAllMsg) XplaClient
	ContractHistory(types.ContractHistoryMsg) XplaClient
	Pinned() XplaClient
	LibwasmvmVersion() XplaClient
}

// Method of helper.
type HelperProvider interface {
	EncodedTxbytesToJsonTx([]byte) ([]byte, error)
}

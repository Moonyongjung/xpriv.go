package client

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/Moonyongjung/xpriv.go/core"
	"github.com/Moonyongjung/xpriv.go/core/anchor"
	"github.com/Moonyongjung/xpriv.go/core/auth"
	"github.com/Moonyongjung/xpriv.go/core/bank"
	"github.com/Moonyongjung/xpriv.go/core/base"
	"github.com/Moonyongjung/xpriv.go/core/crisis"
	"github.com/Moonyongjung/xpriv.go/core/distribution"
	"github.com/Moonyongjung/xpriv.go/core/evidence"
	"github.com/Moonyongjung/xpriv.go/core/evm"
	"github.com/Moonyongjung/xpriv.go/core/feegrant"
	"github.com/Moonyongjung/xpriv.go/core/gov"
	"github.com/Moonyongjung/xpriv.go/core/mint"
	"github.com/Moonyongjung/xpriv.go/core/params"
	"github.com/Moonyongjung/xpriv.go/core/slashing"
	"github.com/Moonyongjung/xpriv.go/core/staking"
	"github.com/Moonyongjung/xpriv.go/core/upgrade"
	"github.com/Moonyongjung/xpriv.go/core/wasm"
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/provider"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/util"

	paramsapp "github.com/Moonyongjung/xpla-private-chain/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	grpc1 "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"
)

var _ provider.XplaClient = &xplaClient{}

// The xpla client is a client for performing all functions within the xpla.go library.
// The user mandatorily inputs chain ID.
type xplaClient struct {
	chainId        string
	encodingConfig paramsapp.EncodingConfig
	grpc           grpc1.ClientConn
	context        context.Context
	VP             []byte

	opts provider.Options

	module  string
	msgType string
	msg     interface{}
	err     error

	externalCoreModule
}

// Make new xpla client.
func NewXplaClient(
	chainId string,
) provider.XplaClient {
	var xplac xplaClient
	return xplac.
		WithChainId(chainId).
		WithEncoding(util.MakeEncodingConfig()).
		WithContext(context.Background()).
		UpdateXplacInCoreModule()
}

// Set options of xpla client.
func (xplac *xplaClient) WithOptions(
	options provider.Options,
) provider.XplaClient {
	return xplac.
		WithPrivateKey(options.PrivateKey).
		WithAccountNumber(options.AccountNumber).
		WithBroadcastMode(options.BroadcastMode).
		WithSequence(options.Sequence).
		WithGasLimit(options.GasLimit).
		WithGasPrice(util.DenomRemove(options.GasPrice)).
		WithGasAdjustment(options.GasAdjustment).
		WithFeeAmount(options.FeeAmount).
		WithSignMode(options.SignMode).
		WithFeeGranter(options.FeeGranter).
		WithTimeoutHeight(options.TimeoutHeight).
		WithURL(options.LcdURL).
		WithGrpc(options.GrpcURL).
		WithRpc(options.RpcURL).
		WithEvmRpc(options.EvmRpcURL).
		WithPagination(options.Pagination).
		WithOutputDocument(options.OutputDocument).
		UpdateXplacInCoreModule()
}

// List of core modules.
// If new modules are implemented, regist ModuleExternal structure with
// receiver method in externalCoreModule.
type externalCoreModule struct {
	anchor.AnchorExternal
	auth.AuthExternal
	bank.BankExternal
	base.BaseExternal
	crisis.CrisisExternal
	distribution.DistributionExternal
	evidence.EvidenceExternal
	evm.EvmExternal
	feegrant.FeegrantExternal
	gov.GovExternal
	mint.MintExternal
	params.ParamsExternal
	slashing.SlashingExternal
	staking.StakingExternal
	upgrade.UpgradeExternal
	wasm.WasmExternal
}

// Update xpla client if data in the xplaClient are changed.
func (xplac *xplaClient) UpdateXplacInCoreModule() provider.XplaClient {
	xplac.externalCoreModule = externalCoreModule{
		anchor.NewAnchorExternal(xplac),
		auth.NewAuthExternal(xplac),
		bank.NewBankExternal(xplac),
		base.NewBaseExternal(xplac),
		crisis.NewCrisisExternal(xplac),
		distribution.NewDistributionExternal(xplac),
		evidence.NewEvidenceExternal(xplac),
		evm.NewEvmExternal(xplac),
		feegrant.NewFeegrantExternal(xplac),
		gov.NewGovExternal(xplac),
		mint.NewMintExternal(xplac),
		params.NewParamsExternal(xplac),
		slashing.NewSlashingExternal(xplac),
		staking.NewStakingExternal(xplac),
		upgrade.NewUpgradeExternal(xplac),
		wasm.NewWasmExternal(xplac),
	}
	return VPInputGrpcContext(xplac)
}

// Set chain ID
func (xplac *xplaClient) WithChainId(chainId string) provider.XplaClient {
	xplac.chainId = chainId
	return xplac.UpdateXplacInCoreModule()
}

// Set encoding configuration
func (xplac *xplaClient) WithEncoding(encodingConfig paramsapp.EncodingConfig) provider.XplaClient {
	xplac.encodingConfig = encodingConfig
	return xplac.UpdateXplacInCoreModule()
}

// Set xpla client context
func (xplac *xplaClient) WithContext(ctx context.Context) provider.XplaClient {
	xplac.context = ctx
	return xplac.UpdateXplacInCoreModule()
}

// Set private key
func (xplac *xplaClient) WithPrivateKey(privateKey key.PrivateKey) provider.XplaClient {
	xplac.opts.PrivateKey = privateKey
	return xplac.UpdateXplacInCoreModule()
}

// Set LCD URL
func (xplac *xplaClient) WithURL(lcdURL string) provider.XplaClient {
	xplac.opts.LcdURL = lcdURL
	return xplac.UpdateXplacInCoreModule()
}

// Set GRPC URL to query or broadcast tx
func (xplac *xplaClient) WithGrpc(grpcUrl string) provider.XplaClient {
	connUrl := util.GrpcUrlParsing(grpcUrl)
	c, err := grpc.Dial(
		connUrl, grpc.WithInsecure(),
	)
	if err != nil {
		xplac.err = err
		return xplac.UpdateXplacInCoreModule()
	}
	xplac.grpc = c
	xplac.opts.GrpcURL = grpcUrl
	return xplac.UpdateXplacInCoreModule()
}

// Set RPC URL of tendermint core
func (xplac *xplaClient) WithRpc(rpcUrl string) provider.XplaClient {
	xplac.opts.RpcURL = rpcUrl
	return xplac.UpdateXplacInCoreModule()
}

// Set RPC URL for evm module
func (xplac *xplaClient) WithEvmRpc(evmRpcUrl string) provider.XplaClient {
	xplac.opts.EvmRpcURL = evmRpcUrl
	return xplac.UpdateXplacInCoreModule()
}

// Set broadcast mode
func (xplac *xplaClient) WithBroadcastMode(broadcastMode string) provider.XplaClient {
	xplac.opts.BroadcastMode = broadcastMode
	return xplac.UpdateXplacInCoreModule()
}

// Set account number
func (xplac *xplaClient) WithAccountNumber(accountNumber string) provider.XplaClient {
	xplac.opts.AccountNumber = accountNumber
	return xplac.UpdateXplacInCoreModule()
}

// Set account sequence
func (xplac *xplaClient) WithSequence(sequence string) provider.XplaClient {
	xplac.opts.Sequence = sequence
	return xplac.UpdateXplacInCoreModule()
}

// Set gas limit
func (xplac *xplaClient) WithGasLimit(gasLimit string) provider.XplaClient {
	xplac.opts.GasLimit = gasLimit
	return xplac.UpdateXplacInCoreModule()
}

// Set Gas price
func (xplac *xplaClient) WithGasPrice(gasPrice string) provider.XplaClient {
	xplac.opts.GasPrice = gasPrice
	return xplac.UpdateXplacInCoreModule()
}

// Set Gas adjustment
func (xplac *xplaClient) WithGasAdjustment(gasAdjustment string) provider.XplaClient {
	xplac.opts.GasAdjustment = gasAdjustment
	return xplac.UpdateXplacInCoreModule()
}

// Set fee amount
func (xplac *xplaClient) WithFeeAmount(feeAmount string) provider.XplaClient {
	xplac.opts.FeeAmount = feeAmount
	return xplac.UpdateXplacInCoreModule()
}

// Set transaction sign mode
func (xplac *xplaClient) WithSignMode(signMode signing.SignMode) provider.XplaClient {
	xplac.opts.SignMode = signMode
	return xplac.UpdateXplacInCoreModule()
}

// Set fee granter
func (xplac *xplaClient) WithFeeGranter(feeGranter sdk.AccAddress) provider.XplaClient {
	xplac.opts.FeeGranter = feeGranter
	return xplac.UpdateXplacInCoreModule()
}

// Set timeout block height
func (xplac *xplaClient) WithTimeoutHeight(timeoutHeight string) provider.XplaClient {
	xplac.opts.TimeoutHeight = timeoutHeight
	return xplac.UpdateXplacInCoreModule()
}

// Set pagination
func (xplac *xplaClient) WithPagination(pagination types.Pagination) provider.XplaClient {
	emptyPagination := types.Pagination{}
	if pagination != emptyPagination {
		pageReq, err := core.ReadPageRequest(pagination)
		if err != nil {
			xplac.err = err
		}
		core.PageRequest = pageReq
	} else {
		core.PageRequest = core.DefaultPagination()
	}

	return xplac.UpdateXplacInCoreModule()
}

// Set output document name
func (xplac *xplaClient) WithOutputDocument(outputDocument string) provider.XplaClient {
	xplac.opts.OutputDocument = outputDocument
	return xplac.UpdateXplacInCoreModule()
}

// Set module name
func (xplac *xplaClient) WithModule(module string) provider.XplaClient {
	xplac.module = module
	return xplac.UpdateXplacInCoreModule()
}

// Set message type of modules
func (xplac *xplaClient) WithMsgType(msgType string) provider.XplaClient {
	xplac.msgType = msgType
	return xplac.UpdateXplacInCoreModule()
}

// Set message
func (xplac *xplaClient) WithMsg(msg interface{}) provider.XplaClient {
	xplac.msg = msg
	return xplac.UpdateXplacInCoreModule()
}

// Set error
func (xplac *xplaClient) WithErr(err error) provider.XplaClient {
	xplac.err = err
	return xplac.UpdateXplacInCoreModule()
}

// Set trigger use VP
func (xplac *xplaClient) WithUseVP(useVP bool) provider.XplaClient {
	xplac.opts.UseVP = useVP
	return xplac.UpdateXplacInCoreModule()
}

// Set Verifiable Presentation by file path
func (xplac *xplaClient) WithVPByPath(vpPath string) provider.XplaClient {
	if xplac.GetUseVP() {
		if xplac.VP == nil {
			f, err := os.Open(vpPath)
			if err != nil {
				xplac.err = err
			}
			defer f.Close()

			jsonByte, err := io.ReadAll(f)
			if err != nil {
				xplac.err = err
			}

			xplac.VP = jsonByte
		}
	}

	return xplac.UpdateXplacInCoreModule()
}

// Set Verifiable Presentation by string
func (xplac *xplaClient) WithVPByString(vp string) provider.XplaClient {
	if xplac.GetUseVP() {
		if xplac.VP == nil {
			jsonByte, err := json.Marshal(vp)
			if err != nil {
				xplac.err = err
			}

			xplac.VP = jsonByte
		}
	}

	return xplac.UpdateXplacInCoreModule()
}

// Get parameters of the xpla client
func (xplac *xplaClient) GetChainId() string                    { return xplac.chainId }
func (xplac *xplaClient) GetPrivateKey() key.PrivateKey         { return xplac.opts.PrivateKey }
func (xplac *xplaClient) GetEncoding() paramsapp.EncodingConfig { return xplac.encodingConfig }
func (xplac *xplaClient) GetContext() context.Context           { return xplac.context }
func (xplac *xplaClient) GetLcdURL() string                     { return xplac.opts.LcdURL }
func (xplac *xplaClient) GetGrpcUrl() string                    { return xplac.opts.GrpcURL }
func (xplac *xplaClient) GetGrpcClient() grpc1.ClientConn       { return xplac.grpc }
func (xplac *xplaClient) GetRpc() string                        { return xplac.opts.RpcURL }
func (xplac *xplaClient) GetEvmRpc() string                     { return xplac.opts.EvmRpcURL }
func (xplac *xplaClient) GetBroadcastMode() string              { return xplac.opts.BroadcastMode }
func (xplac *xplaClient) GetAccountNumber() string              { return xplac.opts.AccountNumber }
func (xplac *xplaClient) GetSequence() string                   { return xplac.opts.Sequence }
func (xplac *xplaClient) GetGasLimit() string                   { return xplac.opts.GasLimit }
func (xplac *xplaClient) GetGasPrice() string                   { return xplac.opts.GasPrice }
func (xplac *xplaClient) GetGasAdjustment() string              { return xplac.opts.GasAdjustment }
func (xplac *xplaClient) GetFeeAmount() string                  { return xplac.opts.FeeAmount }
func (xplac *xplaClient) GetSignMode() signing.SignMode         { return xplac.opts.SignMode }
func (xplac *xplaClient) GetFeeGranter() sdk.AccAddress         { return xplac.opts.FeeGranter }
func (xplac *xplaClient) GetTimeoutHeight() string              { return xplac.opts.TimeoutHeight }
func (xplac *xplaClient) GetPagination() *query.PageRequest     { return core.PageRequest }
func (xplac *xplaClient) GetOutputDocument() string             { return xplac.opts.OutputDocument }
func (xplac *xplaClient) GetModule() string                     { return xplac.module }
func (xplac *xplaClient) GetMsgType() string                    { return xplac.msgType }
func (xplac *xplaClient) GetMsg() interface{}                   { return xplac.msg }
func (xplac *xplaClient) GetErr() error                         { return xplac.err }
func (xplac *xplaClient) GetUseVP() bool                        { return xplac.opts.UseVP }
func (xplac *xplaClient) GetVPByte() []byte                     { return xplac.VP }

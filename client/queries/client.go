package queries

import (
	"context"

	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"
	"github.com/Moonyongjung/xpla-private-chain/app/params"
	cmclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/gogo/protobuf/grpc"
	"github.com/gogo/protobuf/proto"
)

var out []byte
var res proto.Message
var err error

// Query internal XPLA client
type IXplaClient struct {
	Ixplac    ModuleClient
	QueryType uint8
}

type ModuleClient interface {
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
	GetVPByte() []byte
}

func NewIXplaClient(moduleClient ModuleClient, qt uint8) *IXplaClient {
	return &IXplaClient{Ixplac: moduleClient, QueryType: qt}
}

// Print protobuf message by using cosmos sdk codec.
func printProto(i IXplaClient, toPrint proto.Message) ([]byte, error) {
	out, err := i.Ixplac.GetEncoding().Marshaler.MarshalJSON(toPrint)
	if err != nil {
		return nil, util.LogErr(errors.ErrFailedToMarshal, err)
	}
	return out, nil
}

// Print object by using cosmos sdk legacy amino.
func printObjectLegacy(i IXplaClient, toPrint interface{}) ([]byte, error) {
	out, err := i.Ixplac.GetEncoding().Amino.MarshalJSON(toPrint)
	if err != nil {
		return nil, util.LogErr(errors.ErrFailedToMarshal, err)
	}
	return out, nil
}

// For auth module and gov module, make cosmos sdk client for querying.
func clientForQuery(i IXplaClient) (cmclient.Context, error) {
	client, err := cmclient.NewClientFromNode(i.Ixplac.GetRpc())
	if err != nil {
		return cmclient.Context{}, util.LogErr(errors.ErrSdkClient, err)
	}

	clientCtx, err := util.NewClient()
	if err != nil {
		return cmclient.Context{}, err
	}

	clientCtx = clientCtx.
		WithNodeURI(i.Ixplac.GetRpc()).
		WithClient(client)

	return clientCtx, nil
}

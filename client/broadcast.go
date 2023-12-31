package client

import (
	mevm "github.com/Moonyongjung/xpriv.go/core/evm"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

var xplaTxRes types.TxRes

// Broadcast the transaction.
// Default broadcast mode is "sync" if not xpla client has broadcast mode option.
// The broadcast method is determined according to the broadcast mode option of the xpla client.
// For evm transaction broadcast, use a separate method in this function.
func (xplac *xplaClient) Broadcast(txBytes []byte) (*types.TxRes, error) {

	if xplac.GetModule() == mevm.EvmModule {
		return xplac.broadcastEvm(txBytes)

	} else {
		broadcastMode := xplac.GetBroadcastMode()
		switch {
		case broadcastMode == "block":
			return xplac.BroadcastBlock(txBytes)
		case broadcastMode == "async":
			return xplac.BroadcastAsync(txBytes)
		case broadcastMode == "sync":
			return broadcastTx(xplac, txBytes, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
		default:
			return broadcastTx(xplac, txBytes, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
		}
	}
}

// Broadcast the transaction with mode "block".
// It takes precedence over the option of the xpla client.
func (xplac *xplaClient) BroadcastBlock(txBytes []byte) (*types.TxRes, error) {
	if xplac.GetModule() == mevm.EvmModule {
		return xplac.broadcastEvm(txBytes)
	}
	return broadcastTx(xplac, txBytes, txtypes.BroadcastMode_BROADCAST_MODE_BLOCK)
}

// Broadcast the transaction with mode "Async".
// It takes precedence over the option of the xpla client.
func (xplac *xplaClient) BroadcastAsync(txBytes []byte) (*types.TxRes, error) {
	if xplac.GetModule() == mevm.EvmModule {
		return xplac.broadcastEvm(txBytes)
	}
	return broadcastTx(xplac, txBytes, txtypes.BroadcastMode_BROADCAST_MODE_ASYNC)
}

// Broadcast the transaction which is evm transaction by using ethclient of go-ethereum.
func (xplac *xplaClient) broadcastEvm(txBytes []byte) (*types.TxRes, error) {
	if xplac.GetEvmRpc() == "" {
		return nil, util.LogErr(errors.ErrNotSatisfiedOptions, "evm JSON-RPC URL must exist")
	}
	evmClient, err := util.NewEvmClient(xplac.GetEvmRpc(), xplac.GetContext())
	if err != nil {
		return nil, err
	}
	broadcastMode := xplac.GetBroadcastMode()
	return broadcastTxEvm(xplac, txBytes, broadcastMode, evmClient)
}

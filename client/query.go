package client

import (
	"github.com/Moonyongjung/xpriv.go/controller"
	"github.com/Moonyongjung/xpriv.go/core"

	mevm "github.com/Moonyongjung/xpriv.go/core/evm"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"
)

// Query transactions and xpla blockchain information.
// Execute a query of functions for all modules.
// After module query messages are generated, it receives query messages/information to the xpla client receiver and transmits a query message.
func (xplac *xplaClient) Query() (string, error) {
	if xplac.GetErr() != nil {
		return "", xplac.GetErr()
	}

	if xplac.GetGrpcUrl() == "" && xplac.GetLcdURL() == "" {
		if xplac.GetModule() == mevm.EvmModule {
			if xplac.GetEvmRpc() == "" {
				return "", util.LogErr(errors.ErrNotSatisfiedOptions, "evm JSON-RPC URL must exist")
			}

		} else {
			return "", util.LogErr(errors.ErrNotSatisfiedOptions, "at least one of the gRPC URL or LCD URL must exist for query")
		}
	}
	queryClient := core.NewIxplaClient(xplac, setQueryType(xplac))

	return controller.Controller().Get(xplac.GetModule()).NewQueryRouter(*queryClient)
}

func setQueryType(xplac *xplaClient) uint8 {
	// Default query type is gRPC, not LCD.
	if xplac.GetGrpcUrl() != "" {
		return types.QueryGrpc
	} else {
		return types.QueryLcd
	}
}

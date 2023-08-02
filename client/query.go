package client

import (
	"github.com/Moonyongjung/xpriv.go/client/queries"
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

// Query transactions and xpla blockchain information.
// Execute a query of functions for all modules.
// After module query messages are generated, it receives query messages/information to the xpla client receiver and transmits a query message.
func (xplac *XplaClient) Query() (string, error) {
	if xplac.Err != nil {
		return "", xplac.Err
	}

	if xplac.GetGrpcUrl() == "" && xplac.GetLcdURL() == "" {
		if xplac.Module == mevm.EvmModule {
			if xplac.GetEvmRpc() == "" {
				return "", util.LogErr(errors.ErrNotSatisfiedOptions, "evm JSON-RPC URL must exist")
			}

		} else {
			return "", util.LogErr(errors.ErrNotSatisfiedOptions, "at least one of the gRPC URL or LCD URL must exist for query")
		}
	}

	qt := setQueryType(xplac)
	xplac = VPInputGrpcContext(xplac)
	ixplaClient := queries.NewIXplaClient(xplac, qt)

	if xplac.Module == manchor.AnchorModule {
		return ixplaClient.QueryAnchor()

	} else if xplac.Module == mauth.AuthModule {
		return ixplaClient.QueryAuth()

	} else if xplac.Module == mbank.BankModule {
		return ixplaClient.QueryBank()

	} else if xplac.Module == mbase.Base {
		return ixplaClient.QueryBase()

	} else if xplac.Module == mdid.DidModule {
		return ixplaClient.QueryDID()

	} else if xplac.Module == mdist.DistributionModule {
		return ixplaClient.QueryDistribution()

	} else if xplac.Module == mevidence.EvidenceModule {
		return ixplaClient.QueryEvidence()

	} else if xplac.Module == mevm.EvmModule {
		return ixplaClient.QueryEvm()

	} else if xplac.Module == mfeegrant.FeegrantModule {
		return ixplaClient.QueryFeegrant()

	} else if xplac.Module == mgov.GovModule {
		return ixplaClient.QueryGov()

	} else if xplac.Module == mmint.MintModule {
		return ixplaClient.QueryMint()

	} else if xplac.Module == mparams.ParamsModule {
		return ixplaClient.QueryParams()

	} else if xplac.Module == mpriv.PrivateModule {
		return ixplaClient.QueryPrivate()

	} else if xplac.Module == mslashing.SlashingModule {
		return ixplaClient.QuerySlashing()

	} else if xplac.Module == mstaking.StakingModule {
		return ixplaClient.QueryStaking()

	} else if xplac.Module == mupgrade.UpgradeModule {
		return ixplaClient.QueryUpgrade()

	} else if xplac.Module == mwasm.WasmModule {
		return ixplaClient.QueryWasm()

	} else {
		return "", util.LogErr(errors.ErrInvalidRequest, "invalid module")
	}
}

func setQueryType(xplac *XplaClient) uint8 {
	// Default query type is gRPC, not LCD.
	if xplac.Opts.GrpcURL != "" {
		return types.QueryGrpc
	} else {
		return types.QueryLcd
	}
}

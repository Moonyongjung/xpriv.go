package controller

import (
	"sync"

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
)

var once sync.Once
var cc *coreController

// Controller is able to contol modules in the core package.
// Route Tx & Query logic by message type.
// If need to add new modules of XPLA, insert NewCoreModule in the core controller.
type coreController struct {
	cores map[string]core.CoreModule
}

func init() {
	Controller()
}

// Set core controller only once as singleton, and get core controller.
func Controller() *coreController {
	once.Do(func() {
		cc = NewCoreController(
			anchor.NewCoreModule(),
			auth.NewCoreModule(),
			bank.NewCoreModule(),
			base.NewCoreModule(),
			crisis.NewCoreModule(),
			distribution.NewCoreModule(),
			evidence.NewCoreModule(),
			evm.NewCoreModule(),
			feegrant.NewCoreModule(),
			gov.NewCoreModule(),
			mint.NewCoreModule(),
			params.NewCoreModule(),
			slashing.NewCoreModule(),
			staking.NewCoreModule(),
			upgrade.NewCoreModule(),
			wasm.NewCoreModule(),
		)
	})
	return cc
}

// Register routing info of core modules in the hash map.
func NewCoreController(coreModules ...core.CoreModule) *coreController {
	m := make(map[string]core.CoreModule)
	for _, coreModule := range coreModules {
		m[coreModule.Name()] = coreModule
	}

	return &coreController{
		cores: m,
	}
}

// Get info of each module by its name.
func (c coreController) Get(moduleName string) core.CoreModule {
	return c.cores[moduleName]
}

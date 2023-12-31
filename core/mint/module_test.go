package mint_test

import (
	"github.com/Moonyongjung/xpriv.go/core/mint"
)

func (s *IntegrationTestSuite) TestCoreModule() {
	c := mint.NewCoreModule()

	// test get name
	s.Require().Equal(mint.MintModule, c.Name())

	// test tx
	_, err := c.NewTxRouter(nil, "", nil)
	s.Require().Error(err)
}

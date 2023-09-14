package base_test

import (
	"github.com/Moonyongjung/xpriv.go/core/base"
)

func (s *IntegrationTestSuite) TestCoreModule() {
	c := base.NewCoreModule()

	// test get name
	s.Require().Equal(base.Base, c.Name())

	// test tx
	_, err := c.NewTxRouter(nil, "", nil)
	s.Require().Error(err)
}

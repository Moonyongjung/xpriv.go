package auth_test

import (
	"github.com/Moonyongjung/xpriv.go/core/auth"
)

func (s *IntegrationTestSuite) TestCoreModule() {
	c := auth.NewCoreModule()

	// test get name
	s.Require().Equal(auth.AuthModule, c.Name())

	// test tx
	_, err := c.NewTxRouter(nil, "", nil)
	s.Require().Error(err)
}

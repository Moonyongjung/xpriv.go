package evidence_test

import (
	"github.com/Moonyongjung/xpriv.go/core/evidence"
)

func (s *IntegrationTestSuite) TestCoreModule() {
	c := evidence.NewCoreModule()

	// test get name
	s.Require().Equal(evidence.EvidenceModule, c.Name())

	// test tx
	_, err := c.NewTxRouter(nil, "", nil)
	s.Require().Error(err)
}

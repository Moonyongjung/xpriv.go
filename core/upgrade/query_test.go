package upgrade_test

import (
	"strings"
	"testing"

	"github.com/Moonyongjung/xpriv.go/client"
	"github.com/Moonyongjung/xpriv.go/provider"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/Moonyongjung/xpriv.go/util/testutil"
	"github.com/Moonyongjung/xpriv.go/util/testutil/network"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
)

var (
	validatorNumber = 1
)

type IntegrationTestSuite struct {
	suite.Suite

	xplac provider.XplaClient
	apis  []string

	cfg     network.Config
	network *network.Network
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.network = network.New(s.T(), s.cfg)
	s.Require().NoError(s.network.WaitForNextBlock())

	s.xplac = client.NewXplaClient(testutil.TestChainId)
	s.apis = []string{
		s.network.Validators[0].APIAddress,
		s.network.Validators[0].AppConfig.GRPC.Address,
	}
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestModulesVersion() {
	for i, api := range s.apis {
		if i == 0 {
			s.xplac.WithURL(api)
		} else {
			s.xplac.WithGrpc(api)
		}

		res, err := s.xplac.ModulesVersion().Query()
		s.Require().NoError(err)

		var queryModuleVersionsResponse upgradetypes.QueryModuleVersionsResponse
		jsonpb.Unmarshal(strings.NewReader(res), &queryModuleVersionsResponse)

		s.Require().Equal(24, len(queryModuleVersionsResponse.ModuleVersions))
	}
	s.xplac = provider.ResetXplac(s.xplac)
}

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = validatorNumber
	suite.Run(t, NewIntegrationTestSuite(cfg))
}

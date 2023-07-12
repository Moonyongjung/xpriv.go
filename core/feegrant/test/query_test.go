package feegrant_test

import (
	"context"
	"time"

	mfeegrant "github.com/Moonyongjung/xpla-private-chain.go/core/feegrant"
	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/util"
	"github.com/Moonyongjung/xpla-private-chain.go/util/testutil"
	xapp "github.com/Moonyongjung/xpla-private-chain/app"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
)

type TestSuite struct {
	suite.Suite

	ctx     sdk.Context
	context context.Context
	app     *xapp.XplaApp
	keeper  keeper.Keeper
}

func (suite *TestSuite) SetupTest() {
	checkTx := false
	app := testutil.Setup(checkTx, 5)
	ctx := app.BaseApp.NewContext(checkTx, tmproto.Header{
		Time: time.Now(),
	})

	suite.app = app
	suite.ctx = ctx
	suite.context = sdk.WrapSDKContext(ctx)
	suite.keeper = suite.app.FeeGrantKeeper
}

func (suite *TestSuite) TestFeeAllowance() {

	m1, _ := key.NewMnemonic()
	m2, _ := key.NewMnemonic()

	k1, _ := key.NewPrivKey(m1)
	k2, _ := key.NewPrivKey(m2)

	granter, _ := util.GetAddrByPrivKey(k1)
	grantee, _ := util.GetAddrByPrivKey(k2)

	var req *feegrant.QueryAllowanceRequest

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
		preRun    func()
		postRun   func(_ *feegrant.QueryAllowanceResponse)
	}{
		{
			"fail: invalid granter",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: "invalid_granter",
					Grantee: grantee.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantMsg(QueryFeeGrantMsg)
				req = &msg
			},
			true,
			func() {},
			func(*feegrant.QueryAllowanceResponse) {},
		},
		{
			"fail: invalid grantee",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: granter.String(),
					Grantee: "invalid_grantee",
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantMsg(QueryFeeGrantMsg)
				req = &msg
			},
			true,
			func() {},
			func(*feegrant.QueryAllowanceResponse) {},
		},
		{
			"fail: no grants",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: granter.String(),
					Grantee: grantee.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantMsg(QueryFeeGrantMsg)
				req = &msg
			},
			true,
			func() {},
			func(*feegrant.QueryAllowanceResponse) {},
		},
		{
			"valid query: expect single grant",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: granter.String(),
					Grantee: grantee.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantMsg(QueryFeeGrantMsg)
				req = &msg
			},
			false,
			func() {
				suite.grantFeeAllowance(granter, grantee)
			},
			func(response *feegrant.QueryAllowanceResponse) {
				suite.Require().Equal(response.Allowance.Granter, granter.String())
				suite.Require().Equal(response.Allowance.Grantee, grantee.String())
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()
			tc.preRun()
			resp, err := suite.keeper.Allowance(suite.context, req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (suite *TestSuite) TestFeeAllowances() {

	m1, _ := key.NewMnemonic()
	m2, _ := key.NewMnemonic()

	k1, _ := key.NewPrivKey(m1)
	k2, _ := key.NewPrivKey(m2)

	granter, _ := util.GetAddrByPrivKey(k1)
	grantee, _ := util.GetAddrByPrivKey(k2)

	var req *feegrant.QueryAllowancesRequest

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
		preRun    func()
		postRun   func(_ *feegrant.QueryAllowancesResponse)
	}{
		{
			"fail: invalid granter",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Grantee: "invalid_grantee",
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantsByGranteeMsg(QueryFeeGrantMsg)
				req = &msg
			},
			true,
			func() {},
			func(*feegrant.QueryAllowancesResponse) {},
		},
		{
			"fail: no grants",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Grantee: grantee.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantsByGranteeMsg(QueryFeeGrantMsg)
				req = &msg
			},
			false,
			func() {},
			func(resp *feegrant.QueryAllowancesResponse) {
				suite.Require().Equal(len(resp.Allowances), 0)
			},
		},
		{
			"valid query: expect single grant",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Grantee: grantee.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantsByGranteeMsg(QueryFeeGrantMsg)
				req = &msg
			},
			false,
			func() {
				suite.grantFeeAllowance(granter, grantee)
			},
			func(response *feegrant.QueryAllowancesResponse) {
				suite.Require().Equal(len(response.Allowances), 1)
				suite.Require().Equal(response.Allowances[0].Granter, granter.String())
				suite.Require().Equal(response.Allowances[0].Grantee, grantee.String())
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()
			tc.preRun()
			resp, err := suite.keeper.Allowances(suite.context, req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (suite *TestSuite) TestFeeAllowancesByGranter() {

	m1, _ := key.NewMnemonic()
	m2, _ := key.NewMnemonic()

	k1, _ := key.NewPrivKey(m1)
	k2, _ := key.NewPrivKey(m2)

	granter, _ := util.GetAddrByPrivKey(k1)
	grantee, _ := util.GetAddrByPrivKey(k2)

	var req *feegrant.QueryAllowancesByGranterRequest

	testCases := []struct {
		name      string
		malleate  func()
		expectErr bool
		preRun    func()
		postRun   func(_ *feegrant.QueryAllowancesByGranterResponse)
	}{
		{
			"fail: invalid granter",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: "invalid_granter",
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantsByGranterMsg(QueryFeeGrantMsg)
				req = &msg
			},
			true,
			func() {},
			func(*feegrant.QueryAllowancesByGranterResponse) {},
		},
		{
			"fail: no grants",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: granter.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantsByGranterMsg(QueryFeeGrantMsg)
				req = &msg
			},
			false,
			func() {},
			func(resp *feegrant.QueryAllowancesByGranterResponse) {
				suite.Require().Equal(len(resp.Allowances), 0)
			},
		},
		{
			"valid query: expect single grant",
			func() {
				QueryFeeGrantMsg := types.QueryFeeGrantMsg{
					Granter: granter.String(),
				}
				msg, _ := mfeegrant.MakeQueryFeeGrantsByGranterMsg(QueryFeeGrantMsg)
				req = &msg
			},
			false,
			func() {
				suite.grantFeeAllowance(granter, grantee)
			},
			func(response *feegrant.QueryAllowancesByGranterResponse) {
				suite.Require().Equal(len(response.Allowances), 1)
				suite.Require().Equal(response.Allowances[0].Granter, granter.String())
				suite.Require().Equal(response.Allowances[0].Grantee, grantee.String())
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.malleate()
			tc.preRun()
			resp, err := suite.keeper.AllowancesByGranter(suite.context, req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (suite *TestSuite) grantFeeAllowance(granter, grantee sdk.AccAddress) {
	exp := suite.ctx.BlockTime().AddDate(1, 0, 0)
	err := suite.app.FeeGrantKeeper.GrantAllowance(suite.ctx, granter, grantee, &feegrant.BasicAllowance{
		SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("axpla", 555)),
		Expiration: &exp,
	})
	suite.Require().NoError(err)
}

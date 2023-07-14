package client

import (
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/util"
	"github.com/Moonyongjung/xpriv.go/util/testutil"

	xapp "github.com/Moonyongjung/xpla-private-chain/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type TestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *xapp.XplaApp
}

const (
	unsignedTxPath = "../util/testutil/test_files/unsignedTx.json"
	signedTxPath   = "../util/testutil/test_files/signedTx.json"
)

func (suite *TestSuite) SetupTest() {
	checkTx := false
	app := testutil.Setup(checkTx, 5)
	ctx := app.BaseApp.NewContext(checkTx, tmproto.Header{})

	suite.app = app
	suite.ctx = ctx
}

func (suite *TestSuite) TestSimulateCreateUnsignedTx() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 4)
	from := accounts[0]
	to := accounts[1]
	coins := randomSendCoins(r, suite.ctx, from, suite.app.BankKeeper, suite.app.AccountKeeper)

	fee := util.FromUint64ToString(util.MulUint64(util.FromStringToUint64(types.DefaultGasLimit), util.FromStringToUint64(types.DefaultGasPrice)))

	xplac := NewXplaClient(testutil.TestChainId)
	xplac.WithOptions(
		Options{
			GasLimit:  types.DefaultGasLimit,
			FeeAmount: fee,
		},
	)

	bankSendMsg := types.BankSendMsg{
		FromAddress: from.Address.String(),
		ToAddress:   to.Address.String(),
		Amount:      coins.String(),
	}

	txbytes, err := xplac.BankSend(bankSendMsg).CreateUnsignedTx()
	suite.Require().NoError(err)

	clientCtx, err := util.NewClient()
	suite.Require().NoError(err)

	_, _, newTx, err := readTxAndInitContexts(clientCtx, unsignedTxPath)
	suite.Require().NoError(err)

	newTxbytes, err := xplac.EncodingConfig.TxConfig.TxEncoder()(newTx)
	suite.Require().NoError(err)
	suite.Require().Equal(txbytes, newTxbytes)
}

func (suite *TestSuite) TestSimulateSignTx() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 4)
	from := accounts[0]

	xplac := NewXplaClient(testutil.TestChainId)
	xplac.WithOptions(
		Options{
			GasLimit:   types.DefaultGasLimit,
			GasPrice:   types.DefaultGasPrice,
			PrivateKey: from.PrivKey,
		},
	)

	signTxMsg := types.SignTxMsg{
		UnsignedFileName: unsignedTxPath,
		SignatureOnly:    false,
		MultisigAddress:  "",
		FromAddress:      from.Address.String(),
		Overwrite:        false,
		Amino:            false,
	}

	txbytes, err := xplac.SignTx(signTxMsg)
	suite.Require().NoError(err)

	clientCtx, err := util.NewClient()
	suite.Require().NoError(err)

	_, _, newTx, err := readTxAndInitContexts(clientCtx, signedTxPath)
	suite.Require().NoError(err)

	newTxbytes, err := xplac.EncodingConfig.TxConfig.TxJSONEncoder()(newTx)
	suite.Require().NoError(err)
	suite.Require().Equal(txbytes, newTxbytes)
}

func (suite *TestSuite) TestSimulateSignatureOnly() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 4)
	from := accounts[0]
	to := accounts[1]
	m1 := accounts[2]
	m2 := accounts[3]

	accList := []simtypes.Account{from, to, m1, m2}
	xplac := NewXplaClient(testutil.TestChainId)
	xplac.WithOptions(
		Options{
			GasLimit: types.DefaultGasLimit,
			GasPrice: types.DefaultGasPrice,
		},
	)

	for i, acc := range accList {
		xplac.WithOptions(
			Options{
				PrivateKey: acc.PrivKey,
			},
		)

		signTxMsg := types.SignTxMsg{
			UnsignedFileName: unsignedTxPath,
			SignatureOnly:    true,
			MultisigAddress:  "",
			FromAddress:      acc.Address.String(),
			Overwrite:        false,
			Amino:            false,
		}

		txbytes, err := xplac.SignTx(signTxMsg)
		suite.Require().NoError(err)

		signPath := "../util/testutil/test_files/signature" + util.FromIntToString(i) + ".json"

		suite.Require().Equal(txbytes, convertJson(signPath))
	}
}

func (suite *TestSuite) TestSimulateCreateAndSignTx() {
	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, 4)
	from := accounts[0]
	to := accounts[1]
	coins := randomSendCoins(r, suite.ctx, from, suite.app.BankKeeper, suite.app.AccountKeeper)

	xplac := NewXplaClient(testutil.TestChainId)
	xplac.WithPrivateKey(from.PrivKey)

	bankSendMsg := types.BankSendMsg{
		FromAddress: from.Address.String(),
		ToAddress:   to.Address.String(),
		Amount:      coins.String(),
	}

	txbytes, err := xplac.BankSend(bankSendMsg).CreateAndSignTx()
	suite.Require().NoError(err)

	clientCtx, err := util.NewClient()
	suite.Require().NoError(err)

	_, _, newTx, err := readTxAndInitContexts(clientCtx, signedTxPath)
	suite.Require().NoError(err)

	newTxbytes, err := xplac.EncodingConfig.TxConfig.TxEncoder()(newTx)
	suite.Require().NoError(err)
	suite.Require().Equal(txbytes, newTxbytes)
}

func (suite *TestSuite) TestSimulateEncodeAndDecodeTx() {
	xplac := NewXplaClient(testutil.TestChainId)

	encodeTxMsg := types.EncodeTxMsg{
		FileName: unsignedTxPath,
	}

	encodeRes, err := xplac.EncodeTx(encodeTxMsg)
	suite.Require().NoError(err)
	encoded := "Cp8BCpwBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnwKK3hwbGExbDAza21hNHZ2OXFjdmhnY3hmMmdhMHJudjdkcWN1bWF4ZXhzc2gSK3hwbGExaDR4MmpsbnFrenEyazh3cmZ6dnR0bDlwNWdjZmZ6NHhlNWNqMmMaIAoFYXhwbGESFzgzNjIyODM5MDIyNDM4MjI3MDU2NTgyEiMSIQobCgVheHBsYRISMjEyNTAwMDAwMDAwMDAwMDAwEJChDw=="
	suite.Require().Equal(encodeRes, encoded)

	decodeTxMsg := types.DecodeTxMsg{
		EncodedByteString: encoded,
	}

	decodeRes, err := xplac.DecodeTx(decodeTxMsg)
	suite.Require().NoError(err)

	clientCtx, err := util.NewClient()
	suite.Require().NoError(err)

	_, _, newTx, err := readTxAndInitContexts(clientCtx, unsignedTxPath)
	suite.Require().NoError(err)

	newTxbytes, err := xplac.EncodingConfig.TxConfig.TxJSONEncoder()(newTx)
	suite.Require().NoError(err)
	suite.Require().Equal([]byte(decodeRes), newTxbytes)
}

func (suite *TestSuite) TestSimulateValidateSignature() {
	xplac := NewXplaClient(testutil.TestChainId)

	validateSignaturesMsg := types.ValidateSignaturesMsg{
		FileName: signedTxPath,
		Offline:  true,
	}

	res, err := xplac.ValidateSignatures(validateSignaturesMsg)
	suite.Require().NoError(err)
	suite.Require().Equal(res, "success validate")
}

func (suite *TestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
	accounts := testutil.RandomAccounts(r, n)

	initAmt := suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 200000)
	initCoins := sdk.NewCoins(sdk.NewCoin("axpla", initAmt))

	// add coins to the accounts
	for _, account := range accounts {
		acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, account.Address)
		suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
		suite.Require().NoError(testutil.FundAccount(suite.app.BankKeeper, suite.ctx, account.Address, initCoins))
	}

	return accounts
}

func randomSendCoins(
	r *rand.Rand, ctx sdk.Context, account simtypes.Account, bk bankkeeper.Keeper, ak banktypes.AccountKeeper,
) sdk.Coins {
	acc := ak.GetAccount(ctx, account.Address)
	if acc == nil {
		return nil
	}

	spendable := bk.SpendableCoins(ctx, acc.GetAddress())
	sendCoins := simtypes.RandSubsetCoins(r, spendable)
	if sendCoins.Empty() {
		return nil
	}

	return sendCoins
}

func convertJson(filePath string) []byte {
	bytes, _ := os.ReadFile(filePath)
	temp := strings.Replace(string(bytes), " ", "", -1)
	temp = strings.Replace(temp, "\n", "", -1)
	return []byte(temp)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

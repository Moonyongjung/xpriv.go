package types

import (
	xplatypes "github.com/Moonyongjung/xpla-private-chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetChainConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(xplatypes.Bech32PrefixAccAddr, xplatypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(xplatypes.Bech32PrefixValAddr, xplatypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(xplatypes.Bech32PrefixConsAddr, xplatypes.Bech32PrefixConsPub)
	config.SetCoinType(xplatypes.CoinType)
	config.Seal()
}

func init() {
	SetChainConfig()
}

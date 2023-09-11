package anchor

import (
	"github.com/Moonyongjung/xpriv.go/provider"
	"github.com/Moonyongjung/xpriv.go/types"
)

type AnchorExternal struct {
	Xplac provider.XplaClient
}

func NewAnchorExternal(xplac provider.XplaClient) (e AnchorExternal) {
	e.Xplac = xplac
	return e
}

// Tx

// Register anchor account of the validator in the private chain
func (e AnchorExternal) RegisterAnchorAcc(registerAnchorAccMsg types.RegisterAnchorAccMsg) provider.XplaClient {
	msg, err := MakeRegisterAnchorAccMsg(registerAnchorAccMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorRegisterAnchorAccMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Change anchor account of the validator in the private chain
func (e AnchorExternal) ChangeAnchorAcc(changeAnchorAccMsg types.ChangeAnchorAccMsg) provider.XplaClient {
	msg, err := MakeChangeAnchorAccMsg(changeAnchorAccMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorChangeAnchorAccMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query

// Query mapping info of the anchor account.
func (e AnchorExternal) AnchorAcc(anchorAccMsg types.AnchorAccMsg) provider.XplaClient {
	msg, err := MakeAnchorAccMsg(anchorAccMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorQueryAnchorAccMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query aggregated blocks are saved in the state DB.
func (e AnchorExternal) AllAggregatedBlocks() provider.XplaClient {
	msg, err := MakeAllAggregatedBlocksMsg()
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorAllAggregatedBlocksMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query anchoring info includes from/end height and tx hash are saved in the state DB.
func (e AnchorExternal) AnchorInfo(anchorInfoMsg types.AnchorInfoMsg) provider.XplaClient {
	msg, err := MakeAnchorInfoMsg(anchorInfoMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorAnchorInfoMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query Anchored block is recorded in the anchor contract of the public chain.
func (e AnchorExternal) AnchorBlock(anchorBlockMsg types.AnchorBlockMsg) provider.XplaClient {
	msg, err := MakeAnchorBlockMsg(anchorBlockMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorAnchorBlockMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query anchoring transaction body in the public chain.
func (e AnchorExternal) AnchorTxBody(anchorTxBodyMsg types.AnchorTxBodyMsg) provider.XplaClient {
	msg, err := MakeAnchorTxBodyMsg(anchorTxBodyMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorAnchorTxBodyMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Check the consistency of the block in the private chain.
func (e AnchorExternal) AnchorVerify(anchorVerifyMsg types.AnchorVerifyMsg) provider.XplaClient {
	msg, err := MakeAnchorVerifyMsg(anchorVerifyMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorVerifyMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query balances of the anchor account in the public chain.
func (e AnchorExternal) AnchorBalances(anchorBalancesMsg types.AnchorBalancesMsg) provider.XplaClient {
	msg, err := MakeAnchorBalancesMsg(anchorBalancesMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorAnchorBalancesMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query params of anchor module
func (e AnchorExternal) AnchorParams() provider.XplaClient {
	msg, err := MakeAnchorParamsMsg()
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(AnchorModule).
		WithMsgType(AnchorParamsMsgType).
		WithMsg(msg)
	return e.Xplac
}

package did

import (
	"github.com/Moonyongjung/xpriv.go/provider"
	"github.com/Moonyongjung/xpriv.go/types"
)

type DidExternal struct {
	Xplac provider.XplaClient
}

func NewDidExternal(xplac provider.XplaClient) (e DidExternal) {
	e.Xplac = xplac
	return e
}

// Tx

// Create new DID.
func (e DidExternal) CreateDID(createDIDMsg types.CreateDIDMsg) provider.XplaClient {
	msg, err := MakeCreateDIDMsg(createDIDMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidCreateDidMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Update existed DID.
func (e DidExternal) UpdateDID(updateDIDMsg types.UpdateDIDMsg) provider.XplaClient {
	msg, err := MakeUpdateDIDMsg(updateDIDMsg, e.Xplac.GetLcdURL(), e.Xplac.GetGrpcUrl(), e.Xplac.GetGrpcClient(), e.Xplac.GetPrivateKey(), e.Xplac.GetContext())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidUpdateDidMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Deactivate existed DID.
func (e DidExternal) DeactivateDID(deactivateDIDMsg types.DeactivateDIDMsg) provider.XplaClient {
	msg, err := MakeDeactivateDIDMsg(deactivateDIDMsg, e.Xplac.GetLcdURL(), e.Xplac.GetGrpcUrl(), e.Xplac.GetGrpcClient(), e.Xplac.GetPrivateKey(), e.Xplac.GetContext())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidDeactivateDidMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Replace DID moniker
func (e DidExternal) ReplaceDIDMoniker(replaceDIDMonikerMsg types.ReplaceDIDMonikerMsg) provider.XplaClient {
	msg, err := MakeReplaceDIDMonikerMsg(replaceDIDMonikerMsg, e.Xplac.GetLcdURL(), e.Xplac.GetGrpcUrl(), e.Xplac.GetGrpcClient(), e.Xplac.GetPrivateKey(), e.Xplac.GetContext())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidReplaceDidMonikerMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query

// Query DID info.
func (e DidExternal) GetDID(getDIDMsg types.GetDIDMsg) provider.XplaClient {
	msg, err := MakeGetDIDMsg(getDIDMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidGetDidMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query moniker by DID.
func (e DidExternal) MonikerByDID(monikerByDIDMsg types.MonikerByDIDMsg) provider.XplaClient {
	msg, err := MakeMonikerByDIDMsg(monikerByDIDMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidMonikerByDidMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query DID by moniker.
func (e DidExternal) DIDByMoniker(didByMonikerMsg types.DIDByMonikerMsg) provider.XplaClient {
	msg, err := MakeDIDByMonikerMsg(didByMonikerMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidDidByMonikerMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query all DIDs are activated.
func (e DidExternal) AllDIDs() provider.XplaClient {
	msg, err := MakeAllDIDsMsg()
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(DidModule).
		WithMsgType(DidAllDidsMsgType).
		WithMsg(msg)
	return e.Xplac
}

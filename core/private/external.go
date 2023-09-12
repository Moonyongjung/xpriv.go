package private

import (
	"github.com/Moonyongjung/xpriv.go/provider"
	"github.com/Moonyongjung/xpriv.go/types"
)

type PrivateExternal struct {
	Xplac provider.XplaClient
}

func NewPrivateExternal(xplac provider.XplaClient) (e PrivateExternal) {
	e.Xplac = xplac
	return e
}

// Tx

// Set the initial administrator of the private chain
func (e PrivateExternal) InitialAdmin(initialAdminMsg types.InitialAdminMsg) provider.XplaClient {
	msg, err := MakeInitialAdminMsg(initialAdminMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateInitialAdminMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Enroll additional admin of the private chain
func (e PrivateExternal) AddAdmin(addAdminMsg types.AddAdminMsg) provider.XplaClient {
	msg, err := MakeAddAdminMsg(addAdminMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateAddAdminMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Request to participate to the private chain.
func (e PrivateExternal) Participate(participateMsg types.ParticipateMsg) provider.XplaClient {
	msg, err := MakeParticipateMsg(participateMsg, e.Xplac.GetLcdURL(), e.Xplac.GetGrpcUrl(), e.Xplac.GetGrpcClient(), e.Xplac.GetPrivateKey(), e.Xplac.GetContext())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateParticipateMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Accept a participant to join the private chain
func (e PrivateExternal) Accept(acceptMsg types.AcceptMsg) provider.XplaClient {
	msg, err := MakeAcceptMsg(acceptMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateAcceptMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Deny a participant to join the private chain
func (e PrivateExternal) Deny(denyMsg types.DenyMsg) provider.XplaClient {
	msg, err := MakeDenyMsg(denyMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateDenyMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Exile a participant from private chain
func (e PrivateExternal) Exile(exileMsg types.ExileMsg) provider.XplaClient {
	msg, err := MakeExileMsg(exileMsg, e.Xplac.GetPrivateKey())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateExileMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Quit from private chain
func (e PrivateExternal) Quit(quitMsg types.QuitMsg) provider.XplaClient {
	msg, err := MakeQuitMsg(quitMsg, e.Xplac.GetLcdURL(), e.Xplac.GetGrpcUrl(), e.Xplac.GetGrpcClient(), e.Xplac.GetPrivateKey(), e.Xplac.GetContext())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateQuitMsgType).
		WithMsg(msg)
	return e.Xplac
}

// Query

// Query all administrators in the private chain.
func (e PrivateExternal) Admin() provider.XplaClient {
	msg, err := MakeQueryAdminMsg()
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateQueryAdminMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Query participate state of the DID.
func (e PrivateExternal) ParticipateState(participateStateMsg types.ParticipateStateMsg) provider.XplaClient {
	msg, err := MakeParticipateStateMsg(participateStateMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateParticipateStateMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Query participate sequence of the DID.
func (e PrivateExternal) ParticipateSequence(participateSequenceMsg types.ParticipateSequenceMsg) provider.XplaClient {
	msg, err := MakeParticipateSequenceMsg(participateSequenceMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateParticipateSequenceMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Gen the DID signature from DID key.
func (e PrivateExternal) GenDIDSign(genDIDSignMsg types.GenDIDSignMsg) provider.XplaClient {
	msg, err := MakeGenDIDSignMsg(genDIDSignMsg, e.Xplac.GetLcdURL(), e.Xplac.GetGrpcUrl(), e.Xplac.GetGrpcClient(), e.Xplac.GetContext())
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateGenDIDSignMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Get the verifiable credential which is able to use on the private chain for access control.
// Only the VC owner can query by using DID signature
func (e PrivateExternal) IssueVC(issueVCMsg types.IssueVCMsg) provider.XplaClient {
	msg, err := MakeIssueVCMsg(issueVCMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateIssueVCMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Get the verifiable presentation which is able to use on the private chain for access control.
// Participants can acess the private chain when this presentation submit to chain as verifier.
// Only the VP owner can query by using DID signature
func (e PrivateExternal) GetVP(getVPMsg types.GetVPMsg) provider.XplaClient {
	msg, err := MakeGetVPMsg(getVPMsg)
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateGetVPMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Query to check all participate state under review
func (e PrivateExternal) AllUnderReviews() provider.XplaClient {
	msg, err := MakeAllUnderReviewsMsg()
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateAllUnderReviewsMsgType).
		WithMsg(msg)

	return e.Xplac
}

// Query all participants
func (e PrivateExternal) AllParticipants() provider.XplaClient {
	msg, err := MakeAllParticipantsMsg()
	if err != nil {
		return provider.ResetModuleAndMsgXplac(e.Xplac).WithErr(err)
	}
	e.Xplac.WithModule(PrivateModule).
		WithMsgType(PrivateAllParticipantsMsgType).
		WithMsg(msg)

	return e.Xplac
}

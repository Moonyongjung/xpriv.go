package queries

import (
	mpriv "github.com/Moonyongjung/xpla-private-chain.go/core/private"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"

	privtypes "github.com/Moonyongjung/xpla-private-chain/x/private/types"
)

// Query client for private module.
func (i IXplaClient) QueryPrivate() (string, error) {
	if i.QueryType == types.QueryGrpc {
		return queryByGrpcPrivate(i)
	} else {
		return queryByLcdPrivate(i)
	}

}

func queryByGrpcPrivate(i IXplaClient) (string, error) {
	queryClient := privtypes.NewQueryClient(i.Ixplac.GetGrpcClient())

	switch {
	// Admins
	case i.Ixplac.GetMsgType() == mpriv.PrivateQueryAdminMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryAdminRequest)
		res, err = queryClient.Admin(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// Particpate state of the DID
	case i.Ixplac.GetMsgType() == mpriv.PrivateParticipateStateMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryParticipateStateRequest)
		res, err = queryClient.ParticipateState(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// Gen DID signature
	case i.Ixplac.GetMsgType() == mpriv.PrivateGenDIDSignMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(string)

		var didSigBase64 privtypes.DidSigBase64
		didSigBase64.Base64Sig = convertMsg

		res = &didSigBase64

		// Issue VC
	case i.Ixplac.GetMsgType() == mpriv.PrivateIssueVCMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryIssueVCRequest)

		res, err = queryClient.IssueVC(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// Get VP
	case i.Ixplac.GetMsgType() == mpriv.PrivateGetVPMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryGetVPRequest)

		res, err = queryClient.GetVP(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}
	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err = printProto(i, res)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

const (
	privateAdminLabel            = "admin"
	privateParticipateStateLabel = "participate_state"
	privateIssueVCLabel          = "issue_vc"
	privateGetVPLabel            = "get_vp"
)

func queryByLcdPrivate(i IXplaClient) (string, error) {
	url := "/xpla/private/v1beta1/"

	switch {
	// Admins
	case i.Ixplac.GetMsgType() == mpriv.PrivateQueryAdminMsgType:
		url = url + util.MakeQueryLabels(privateAdminLabel)

		// Participate state of the DID
	case i.Ixplac.GetMsgType() == mpriv.PrivateParticipateStateMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryParticipateStateRequest)

		url = url + util.MakeQueryLabels(privateParticipateStateLabel, convertMsg.DidBase64)

		// Gen DID sign
	case i.Ixplac.GetMsgType() == mpriv.PrivateGenDIDSignMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(string)

		return convertMsg, nil

		// Issue VC
	case i.Ixplac.GetMsgType() == mpriv.PrivateIssueVCMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryIssueVCRequest)

		url = url + util.MakeQueryLabels(privateIssueVCLabel)

		bodyByte, err := i.Ixplac.GetEncoding().Marshaler.MarshalJSON(convertMsg.Body)
		if err != nil {
			return "", util.LogErr(errors.ErrParse, err)
		}

		out, err := util.CtxHttpClient("POST", i.Ixplac.GetLcdURL()+url, bodyByte, i.Ixplac.GetContext())
		if err != nil {
			return "", err
		}

		return string(out), nil

		// Get VP
	case i.Ixplac.GetMsgType() == mpriv.PrivateGetVPMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryGetVPRequest)

		url = url + util.MakeQueryLabels(privateGetVPLabel)

		bodyByte, err := i.Ixplac.GetEncoding().Marshaler.MarshalJSON(convertMsg.Body)
		if err != nil {
			return "", util.LogErr(errors.ErrParse, err)
		}

		out, err := util.CtxHttpClient("POST", i.Ixplac.GetLcdURL()+url, bodyByte, i.Ixplac.GetContext())
		if err != nil {
			return "", err
		}

		return string(out), nil

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err := util.CtxHttpClient("GET", i.Ixplac.GetLcdURL()+url, nil, i.Ixplac.GetContext())
	if err != nil {
		return "", err
	}

	return string(out), nil
}
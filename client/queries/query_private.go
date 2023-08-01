package queries

import (
	mpriv "github.com/Moonyongjung/xpriv.go/core/private"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

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

		// Particpate sequence of the DID
	case i.Ixplac.GetMsgType() == mpriv.PrivateParticipateSequenceMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryParticipateSequenceRequest)
		res, err = queryClient.ParticipateSequence(
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

		// All under reviews
	case i.Ixplac.GetMsgType() == mpriv.PrivateAllUnderReviewsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryAllUnderReviewsRequest)
		res, err = queryClient.AllUnderReviews(
			i.Ixplac.GetContext(),
			&convertMsg,
		)
		if err != nil {
			return "", util.LogErr(errors.ErrGrpcRequest, err)
		}

		// All participants
	case i.Ixplac.GetMsgType() == mpriv.PrivateAllParticipantsMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryAllParticipantsRequest)
		res, err = queryClient.AllParticipants(
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
	privateAdminLabel               = "admin"
	privateParticipateStateLabel    = "participate_state"
	privateParticipateSequenceLabel = "participate_sequence"
	privateIssueVCLabel             = "issue_vc"
	privateGetVPLabel               = "get_vp"
	privateAllUnderReviewsLabel     = "all_under_reviews"
	privateAllParticipantsLabel     = "all_participants"
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

		url = url + util.MakeQueryLabels(privateParticipateStateLabel, convertMsg.Did)

		// Participate sequence of the DID
	case i.Ixplac.GetMsgType() == mpriv.PrivateParticipateSequenceMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(privtypes.QueryParticipateSequenceRequest)

		url = url + util.MakeQueryLabels(privateParticipateSequenceLabel, convertMsg.Did)

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

		// all under reviews
	case i.Ixplac.GetMsgType() == mpriv.PrivateAllUnderReviewsMsgType:
		url = url + util.MakeQueryLabels(privateAllUnderReviewsLabel)

		// all participants
	case i.Ixplac.GetMsgType() == mpriv.PrivateAllParticipantsMsgType:
		url = url + util.MakeQueryLabels(privateAllParticipantsLabel)

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err := util.CtxHttpClient("GET", i.Ixplac.GetLcdURL()+url, nil, i.Ixplac.GetContext())
	if err != nil {
		return "", err
	}

	return string(out), nil
}

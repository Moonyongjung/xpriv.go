package queries

import (
	mdid "github.com/Moonyongjung/xpriv.go/core/did"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	didtypes "github.com/Moonyongjung/xpla-private-chain/x/did/types"
)

// Query client for DID module.
func (i IXplaClient) QueryDID() (string, error) {
	if i.QueryType == types.QueryGrpc {
		return queryByGrpcDID(i)
	} else {
		return queryByLcdDID(i)
	}

}

func queryByGrpcDID(i IXplaClient) (string, error) {
	queryClient := didtypes.NewQueryClient(i.Ixplac.GetGrpcClient())

	switch {
	// Get DID
	case i.Ixplac.GetMsgType() == mdid.DidGetDidMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(didtypes.QueryDIDRequest)
		res, err = queryClient.DID(
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
	didGetDIDLabel = "dids"
)

func queryByLcdDID(i IXplaClient) (string, error) {
	url := "/xpla/did/v1beta1/"

	switch {
	// Get DID
	case i.Ixplac.GetMsgType() == mdid.DidGetDidMsgType:
		convertMsg, _ := i.Ixplac.GetMsg().(didtypes.QueryDIDRequest)
		url = url + util.MakeQueryLabels(didGetDIDLabel, convertMsg.DidBase64)

	default:
		return "", util.LogErr(errors.ErrInvalidMsgType, i.Ixplac.GetMsgType())
	}

	out, err := util.CtxHttpClient("GET", i.Ixplac.GetLcdURL()+url, nil, i.Ixplac.GetContext())
	if err != nil {
		return "", err
	}

	return string(out), nil
}

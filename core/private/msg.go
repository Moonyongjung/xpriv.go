package private

import (
	"context"

	"github.com/Moonyongjung/xpriv.go/core"
	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/gogo/protobuf/grpc"

	privtypes "github.com/Moonyongjung/xpla-private-chain/x/private/types"
)

// (Tx) make msg - initial admin
func MakeInitialAdminMsg(initialAdminMsg types.InitialAdminMsg, privKey key.PrivateKey) (privtypes.MsgInitialAdmin, error) {
	return parseInitialAdminArgs(initialAdminMsg, privKey)
}

// (Tx) make msg - add admin
func MakeAddAdminMsg(addAdminMsg types.AddAdminMsg, privKey key.PrivateKey) (privtypes.MsgAddAdmin, error) {
	return parseAddAdminArgs(addAdminMsg, privKey)
}

// (Tx) make msg - participate
func MakeParticipateMsg(participateMsg types.ParticipateMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (privtypes.MsgParticipate, error) {
	return parseParticipateArgs(participateMsg, lcdUrl, grpcUrl, grpcConn, privKey, ctx)
}

// (Tx) make msg - accept
func MakeAcceptMsg(acceptMsg types.AcceptMsg, privKey key.PrivateKey) (privtypes.MsgAccept, error) {
	return parseAcceptArgs(acceptMsg, privKey)
}

// (Tx) make msg - deny
func MakeDenyMsg(denyMsg types.DenyMsg, privKey key.PrivateKey) (privtypes.MsgDeny, error) {
	return parseDenyArgs(denyMsg, privKey)
}

// (Tx) make msg - exile
func MakeExileMsg(exileMsg types.ExileMsg, privKey key.PrivateKey) (privtypes.MsgExile, error) {
	return parseExileArgs(exileMsg, privKey)
}

// (Tx) make msg - quit
func MakeQuitMsg(quitMsg types.QuitMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (privtypes.MsgQuit, error) {
	return parseQuitArgs(quitMsg, lcdUrl, grpcUrl, grpcConn, privKey, ctx)
}

// (Query) make msg - query admin
func MakeQueryAdminMsg() (privtypes.QueryAdminRequest, error) {
	return privtypes.QueryAdminRequest{}, nil
}

// (Query) make msg - participate state
func MakeParticipateStateMsg(participateStateMsg types.ParticipateStateMsg) (privtypes.QueryParticipateStateRequest, error) {
	return privtypes.QueryParticipateStateRequest{
		Did: participateStateMsg.DID,
	}, nil
}

// (Query) make msg - participate sequence
func MakeParticipateSequenceMsg(participateSequenceMsg types.ParticipateSequenceMsg) (privtypes.QueryParticipateSequenceRequest, error) {
	return privtypes.QueryParticipateSequenceRequest{
		Did: participateSequenceMsg.DID,
	}, nil
}

// (Query) make msg - gen DID signature
func MakeGenDIDSignMsg(genDIDSignMsg types.GenDIDSignMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, ctx context.Context) (string, error) {
	return parseGenDIDSignArgs(genDIDSignMsg, lcdUrl, grpcUrl, grpcConn, ctx)
}

// (Query) make msg - issue vc
func MakeIssueVCMsg(issueVCMsg types.IssueVCMsg) (privtypes.QueryIssueVCRequest, error) {
	return parseIssueVCArgs(issueVCMsg)
}

// (Query) make msg - get vp
func MakeGetVPMsg(getVPMsg types.GetVPMsg) (privtypes.QueryGetVPRequest, error) {
	return parseGetVPArgs(getVPMsg)
}

// (Query) make msg - all under reviews
func MakeAllUnderReviewsMsg() (privtypes.QueryAllUnderReviewsRequest, error) {
	return privtypes.QueryAllUnderReviewsRequest{
		Pagination: core.PageRequest,
	}, nil
}

// (Query) make msg - all participants
func MakeAllParticipantsMsg() (privtypes.QueryAllParticipantsRequest, error) {
	return privtypes.QueryAllParticipantsRequest{
		Pagination: core.PageRequest,
	}, nil
}

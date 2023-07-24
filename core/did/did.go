package did

import (
	"context"

	"github.com/Moonyongjung/xpriv.go/key"
	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/gogo/protobuf/grpc"

	didtypes "github.com/Moonyongjung/xpla-private-chain/x/did/types"
)

// (Tx) make msg - create DID
func MakeCreateDIDMsg(createDIDMsg types.CreateDIDMsg, privKey key.PrivateKey) (didtypes.MsgCreateDID, error) {
	return parseCreateDIDArgs(createDIDMsg, privKey)
}

// (Tx) make msg - update DID
func MakeUpdateDIDMsg(updateDIDMsg types.UpdateDIDMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (didtypes.MsgUpdateDID, error) {
	return parseUpdateDIDArgs(updateDIDMsg, lcdUrl, grpcUrl, grpcConn, privKey, ctx)
}

// (Tx) make msg - deactivate DID
func MakeDeactivateDIDMsg(deactivateDIDMsg types.DeactivateDIDMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (didtypes.MsgDeactivateDID, error) {
	return parseDeactivateDIDArgs(deactivateDIDMsg, lcdUrl, grpcUrl, grpcConn, privKey, ctx)
}

// (Tx) make msg - replace DID moniker
func MakeReplaceDIDMonikerMsg(replaceDIDMonikerMsg types.ReplaceDIDMonikerMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (didtypes.MsgReplaceDIDMoniker, error) {
	return parseReplaceDIDMonikerArgs(replaceDIDMonikerMsg, lcdUrl, grpcUrl, grpcConn, privKey, ctx)
}

// (Query) - get DID
func MakeGetDIDMsg(getDIDMsg types.GetDIDMsg) (didtypes.QueryDIDRequest, error) {
	return didtypes.QueryDIDRequest{
		Did: getDIDMsg.DIDOrMoniker,
	}, nil
}

// (Query) - moniker by DID
func MakeMonikerByDIDMsg(monikerByDIDMsg types.MonikerByDIDMsg) (didtypes.QueryMonikerByDIDRequest, error) {
	return didtypes.QueryMonikerByDIDRequest{
		Did: monikerByDIDMsg.DID,
	}, nil
}

// (Query) - did by moniker
func MakeDIDByMonikerMsg(didByMonikerMsg types.DIDByMonikerMsg) (didtypes.QueryDIDByMonikerRequest, error) {
	return didtypes.QueryDIDByMonikerRequest{
		Moniker: didByMonikerMsg.Moniker,
	}, nil
}

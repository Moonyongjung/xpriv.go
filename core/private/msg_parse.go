package did

import (
	"context"
	"encoding/base64"

	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"
	"github.com/gogo/protobuf/grpc"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	xcrypto "github.com/Moonyongjung/xpla-private-chain/crypto"
	didcrypto "github.com/Moonyongjung/xpla-private-chain/x/did/client/crypto"
	didtypes "github.com/Moonyongjung/xpla-private-chain/x/did/types"
	privtypes "github.com/Moonyongjung/xpla-private-chain/x/private/types"
)

// Parsing - initial admin
func parseInitialAdminArgs(initialAdminMsg types.InitialAdminMsg, privKey key.PrivateKey) (privtypes.MsgInitialAdmin, error) {
	did, didKey, err := privtypes.ParseDIDKey(initialAdminMsg.InitAdminDIDKey)
	if err != nil {
		return privtypes.MsgInitialAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	base64Proof, err := makeBase64Proof(did, didKey, did, initialAdminMsg.DIDKeyPath, initialAdminMsg.DIDPassphrase)
	if err != nil {
		return privtypes.MsgInitialAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgInitialAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgInitialAdmin(didKey, base64Proof, fromAddress.String()), nil
}

// Parsing - add admin
func parseAddAdminArgs(addAdminMsg types.AddAdminMsg, privKey key.PrivateKey) (privtypes.MsgAddAdmin, error) {
	newAdminDID, newAdminDIDKey, err := privtypes.ParseDIDKey(addAdminMsg.NewAdminDIDKey)
	if err != nil {
		return privtypes.MsgAddAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	initAdminDID, initAdminDIDKey, err := privtypes.ParseDIDKey(addAdminMsg.InitAdminDIDKey)
	if err != nil {
		return privtypes.MsgAddAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	base64Proof, err := makeBase64Proof(initAdminDID, initAdminDIDKey, newAdminDID, addAdminMsg.InitAdminDIDKeyPath, addAdminMsg.InitAdminDIDPassphrase)
	if err != nil {
		return privtypes.MsgAddAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgAddAdmin{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgAddAdmin(addAdminMsg.NewAdminAddress, newAdminDIDKey, initAdminDIDKey, base64Proof, fromAddress.String()), nil
}

// Parsing - participate
func parseParticipateArgs(participantMsg types.ParticipateMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (privtypes.MsgParticipate, error) {
	did, didKey, err := privtypes.ParseDIDKey(participantMsg.ParticipantDIDKey)
	if err != nil {
		return privtypes.MsgParticipate{}, util.LogErr(errors.ErrParse, err)
	}

	didDocumentWithSeq, err := util.GetDIDDocByQueryClient(did, lcdUrl, grpcUrl, grpcConn, ctx)
	if err != nil {
		return privtypes.MsgParticipate{}, util.LogErr(errors.ErrHttpRequest, err)
	}

	didSeq := util.FromUint64ToString(didDocumentWithSeq.Sequence)

	didSigBase64, err := createSign(did, didKey, did, didSeq, participantMsg.DIDKeyPath, participantMsg.DIDPassphrase)
	if err != nil {
		return privtypes.MsgParticipate{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgParticipate{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgParticipate(didKey, didSigBase64, fromAddress.String()), nil
}

// Parsing - accept
func parseAcceptArgs(acceptMsg types.AcceptMsg, privKey key.PrivateKey) (privtypes.MsgAccept, error) {
	participantDID, err := didtypes.ParseDID(acceptMsg.ParticipantDID)
	if err != nil {
		return privtypes.MsgAccept{}, util.LogErr(errors.ErrParse, err)
	}

	adminDID, adminDIDKey, err := privtypes.ParseDIDKey(acceptMsg.AdminDIDKey)
	if err != nil {
		return privtypes.MsgAccept{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgAccept{}, util.LogErr(errors.ErrParse, err)
	}

	base64Proof, err := makeBase64Proof(adminDID, adminDIDKey, participantDID, acceptMsg.AdminDIDKeyPath, acceptMsg.AdminDIDPassphrase)
	if err != nil {
		return privtypes.MsgAccept{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgAccept(participantDID, adminDIDKey, base64Proof, fromAddress.String()), nil
}

// Parsing - deny
func parseDenyArgs(denyMsg types.DenyMsg, privKey key.PrivateKey) (privtypes.MsgDeny, error) {
	participandDID, err := didtypes.ParseDID(denyMsg.ParticipantDID)
	if err != nil {
		return privtypes.MsgDeny{}, util.LogErr(errors.ErrParse, err)
	}

	adminDID, err := didtypes.ParseDID(denyMsg.AdminDID)
	if err != nil {
		return privtypes.MsgDeny{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgDeny{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgDeny(participandDID, adminDID, fromAddress.String()), nil
}

// Parsing - exile
func parseExileArgs(exileMsg types.ExileMsg, privKey key.PrivateKey) (privtypes.MsgExile, error) {
	participantDID, err := didtypes.ParseDID(exileMsg.ParticipantDID)
	if err != nil {
		return privtypes.MsgExile{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgExile{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgExile(participantDID, fromAddress.String()), nil
}

// Parsing - quit
func parseQuitArgs(quitMsg types.QuitMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (privtypes.MsgQuit, error) {
	did, didKey, err := privtypes.ParseDIDKey(quitMsg.ParticipantDIDKey)
	if err != nil {
		return privtypes.MsgQuit{}, util.LogErr(errors.ErrParse, err)
	}

	document, err := util.GetDIDDocByQueryClient(did, lcdUrl, grpcUrl, grpcConn, ctx)
	if err != nil {
		return privtypes.MsgQuit{}, util.LogErr(errors.ErrInvalidRequest, err)
	}

	didSeq := util.FromUint64ToString(document.Sequence)

	didSigBase64, err := createSign(did, didKey, did, didSeq, quitMsg.DIDKeyPath, quitMsg.DIDPassphrase)
	if err != nil {
		return privtypes.MsgQuit{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return privtypes.MsgQuit{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.NewMsgQuit(didKey, didSigBase64, fromAddress.String()), nil
}

// Parsing - participate state
func parseParticipateStateArgs(participateStateMsg types.ParticipateStateMsg) (privtypes.QueryParticipateStateRequest, error) {
	did, err := didtypes.ParseDID(participateStateMsg.DID)
	if err != nil {
		return privtypes.QueryParticipateStateRequest{}, util.LogErr(errors.ErrParse, err)
	}

	return privtypes.QueryParticipateStateRequest{
		DidBase64: base64.StdEncoding.EncodeToString([]byte(did)),
	}, nil
}

// Parsing - gen DID signature
func parseGenDIDSignArgs(genDIDSignMsg types.GenDIDSignMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, ctx context.Context) (string, error) {
	did, didKey, err := privtypes.ParseDIDKey(genDIDSignMsg.DIDKey)
	if err != nil {
		return "", util.LogErr(errors.ErrParse, err)
	}

	document, err := util.GetDIDDocByQueryClient(did, lcdUrl, grpcUrl, grpcConn, ctx)
	if err != nil {
		return "", util.LogErr(errors.ErrParse, err)
	}

	didSeq := util.FromUint64ToString(document.Sequence)

	didSigBase64, err := createSign(did, didKey, did, didSeq, genDIDSignMsg.DIDKeyPath, genDIDSignMsg.DIDPassphrase)
	if err != nil {
		return "", util.LogErr(errors.ErrParse, err)
	}

	return didSigBase64, nil
}

// Parsing - issue vc
func parseIssueVCArgs(issueVCMsg types.IssueVCMsg) (privtypes.QueryIssueVCRequest, error) {
	did, didKey, err := privtypes.ParseDIDKey(issueVCMsg.DIDKey)
	if err != nil {
		return privtypes.QueryIssueVCRequest{}, util.LogErr(errors.ErrParse, err)
	}

	didBase64 := base64.StdEncoding.EncodeToString([]byte(did))
	methodKey := privtypes.FromDidKeyToMethodKey(didKey)

	return privtypes.QueryIssueVCRequest{
		Body: &privtypes.VerifiableBody{
			DidBase64:    didBase64,
			DidSigBase64: issueVCMsg.DIDSignBase64,
			MethodKey:    methodKey,
		},
	}, nil
}

// Parsing - get vp
func parseGetVPArgs(getVPMsg types.GetVPMsg) (privtypes.QueryGetVPRequest, error) {
	did, didKey, err := privtypes.ParseDIDKey(getVPMsg.DIDKey)
	if err != nil {
		return privtypes.QueryGetVPRequest{}, util.LogErr(errors.ErrParse, err)
	}

	didBase64 := base64.StdEncoding.EncodeToString([]byte(did))
	methodKey := privtypes.FromDidKeyToMethodKey(didKey)

	return privtypes.QueryGetVPRequest{
		Body: &privtypes.VerifiableBody{
			DidBase64:    didBase64,
			DidSigBase64: getVPMsg.DIDSignBase64,
			MethodKey:    methodKey,
		},
	}, nil
}

func makeBase64Proof(adminDid, adminDidKey, participantDid, didKeyPath, didKeyPassphrase string) (string, error) {
	return createSign(adminDid, adminDidKey, participantDid, privtypes.ProofSequence, didKeyPath, didKeyPassphrase)
}

// create the DID signature by using DID private key from the key store directory
func createSign(signingDid, signingDidKey, targetDid, sigSeq, didKeyPath, didKeyPassphrase string) (string, error) {
	verificationMethodID, err := didtypes.ParseVerificationMethodID(signingDidKey, signingDid)
	if err != nil {
		return "", err
	}

	didPrivKey, err := getPrivKeyFromKeyStore(didKeyPath, didKeyPassphrase, verificationMethodID)
	if err != nil {
		return "", err
	}

	proof, err := privtypes.Sign(targetDid, sigSeq, didPrivKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(proof), nil
}

func getPrivKeyFromKeyStore(didKeyPath, didPassphrase string, verificationMethodID string) (secp256k1.PrivKey, error) {
	ks, err := didcrypto.NewKeyStore(didKeyPath)
	if err != nil {
		return nil, util.LogErr(errors.ErrInvalidRequest, err)
	}

	didPrivKeyBytes, err := ks.LoadByAddress(string(verificationMethodID), didPassphrase)
	if err != nil {
		return nil, util.LogErr(errors.ErrInvalidRequest, err)
	}

	return xcrypto.PrivKeyFromBytes(didPrivKeyBytes)
}

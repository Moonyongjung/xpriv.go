package did

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Moonyongjung/xpla-private-chain.go/key"
	"github.com/Moonyongjung/xpla-private-chain.go/types"
	"github.com/Moonyongjung/xpla-private-chain.go/types/errors"
	"github.com/Moonyongjung/xpla-private-chain.go/util"
	"github.com/gogo/protobuf/grpc"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	xcrypto "github.com/Moonyongjung/xpla-private-chain/crypto"
	didcrypto "github.com/Moonyongjung/xpla-private-chain/x/did/client/crypto"
	didtypes "github.com/Moonyongjung/xpla-private-chain/x/did/types"
)

// Parsing - create DID
func parseCreateDIDArgs(createDIDMsg types.CreateDIDMsg, privKey key.PrivateKey) (didtypes.MsgCreateDID, error) {
	if createDIDMsg.SaveDIDKeyPath == "" {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrNotFound, "indicate directory for saving DID key")
	}

	didPrivKey, err := didcrypto.GenSecp256k1PrivKey(createDIDMsg.DIDMnemonic, createDIDMsg.DIDPassphrase)
	if err != nil {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrParse, err)
	}

	pubKey := xcrypto.PubKeyBytes(xcrypto.DerivePubKey(didPrivKey))
	did := didtypes.NewDID(pubKey)
	verificationMethodID := didtypes.NewVerificationMethodID(did, "key1")
	verificationMethod := didtypes.NewVerificationMethod(verificationMethodID, didtypes.ES256K_2019, did, pubKey)
	verificationMethods := []*didtypes.VerificationMethod{
		&verificationMethod,
	}
	relationship := didtypes.NewVerificationRelationship(verificationMethods[0].Id)
	authentications := []didtypes.VerificationRelationship{
		relationship,
	}
	doc := didtypes.NewDIDDocument(did, didtypes.WithVerificationMethods(verificationMethods), didtypes.WithAuthentications(authentications))

	sig, err := didtypes.Sign(&doc, didtypes.InitialSequence, didPrivKey)
	if err != nil {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrParse, err)
	}

	msg := didtypes.NewMsgCreateDID(did, doc, verificationMethodID, sig, fromAddress.String())
	if err := msg.ValidateBasic(); err != nil {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrInvalidRequest, err)
	}

	ks, err := didcrypto.NewKeyStore(createDIDMsg.SaveDIDKeyPath)
	if err != nil {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrParse, err)
	}

	_, err = ks.Save(verificationMethodID, didPrivKey[:], createDIDMsg.DIDPassphrase)
	if err != nil {
		return didtypes.MsgCreateDID{}, util.LogErr(errors.ErrParse, err)
	}

	return msg, nil
}

// Parsing - update DID
func parseUpdateDIDArgs(updateDIDMsg types.UpdateDIDMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (didtypes.MsgUpdateDID, error) {
	did, err := didtypes.ParseDID(updateDIDMsg.DID)
	if err != nil {
		return didtypes.MsgUpdateDID{}, util.LogErr(errors.ErrInvalidMsgType, err)
	}

	verificationMethodID, err := didtypes.ParseVerificationMethodID(did+"#"+updateDIDMsg.KeyID, did)
	if err != nil {
		return didtypes.MsgUpdateDID{}, util.LogErr(errors.ErrInvalidMsgType, err)
	}

	doc, err := readDIDDocFrom(updateDIDMsg.DIDDocumentPath)
	if err != nil {
		return didtypes.MsgUpdateDID{}, util.LogErr(errors.ErrInvalidRequest, err)
	}

	didPrivKey, err := getPrivKeyFromKeyStore(updateDIDMsg.DIDKeyPath, updateDIDMsg.DIDPassphrase, verificationMethodID)
	if err != nil {
		return didtypes.MsgUpdateDID{}, util.LogErr(errors.ErrInvalidRequest, err)
	}

	sign, err := signUsingCurrentSeq(did, lcdUrl, grpcUrl, grpcConn, ctx, didPrivKey, doc)
	if err != nil {
		return didtypes.MsgUpdateDID{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return didtypes.MsgUpdateDID{}, util.LogErr(errors.ErrParse, err)
	}

	return didtypes.NewMsgUpdateDID(did, doc, verificationMethodID, sign, fromAddress.String()), nil
}

// Parsing - deactivate DID
func parseDeactivateDIDArgs(deactivateDIDMsg types.DeactivateDIDMsg, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, privKey key.PrivateKey, ctx context.Context) (didtypes.MsgDeactivateDID, error) {
	did, err := didtypes.ParseDID(deactivateDIDMsg.DID)
	if err != nil {
		return didtypes.MsgDeactivateDID{}, util.LogErr(errors.ErrParse, err)
	}

	verificationMethodID, err := didtypes.ParseVerificationMethodID(did+"#"+deactivateDIDMsg.KeyID, did)
	if err != nil {
		return didtypes.MsgDeactivateDID{}, util.LogErr(errors.ErrParse, err)
	}

	didPrivKey, err := getPrivKeyFromKeyStore(deactivateDIDMsg.DIDKeyPath, deactivateDIDMsg.DIDPassphrase, verificationMethodID)
	if err != nil {
		return didtypes.MsgDeactivateDID{}, util.LogErr(errors.ErrInvalidRequest, err)
	}

	doc := didtypes.DIDDocument{
		Id: did,
	}

	sign, err := signUsingCurrentSeq(did, lcdUrl, grpcUrl, grpcConn, ctx, didPrivKey, doc)
	if err != nil {
		return didtypes.MsgDeactivateDID{}, util.LogErr(errors.ErrParse, err)
	}

	fromAddress, err := util.GetAddrByPrivKey(privKey)
	if err != nil {
		return didtypes.MsgDeactivateDID{}, util.LogErr(errors.ErrParse, err)
	}

	return didtypes.NewMsgDeactivateDID(did, verificationMethodID, sign, fromAddress.String()), nil
}

// Parsing - get DID
func parseGetDIDArgs(getDIDMsg types.GetDIDMsg) (didtypes.QueryDIDRequest, error) {
	did, err := didtypes.ParseDID(getDIDMsg.DID)
	if err != nil {
		return didtypes.QueryDIDRequest{}, util.LogErr(errors.ErrParse, err)
	}

	didBase64 := base64.StdEncoding.EncodeToString([]byte(did))

	return didtypes.QueryDIDRequest{
		DidBase64: didBase64,
	}, nil
}

func readDIDDocFrom(path string) (didtypes.DIDDocument, error) {
	var doc didtypes.DIDDocument

	file, err := os.Open(path)
	if err != nil {
		return doc, err
	}
	defer file.Close()

	// Use gogoproto's jsonpb to handle camelCase and custom types as well as snake_case.
	if err := jsonpb.Unmarshal(file, &doc); err != nil {
		return doc, fmt.Errorf("fail to decode DIDDocument JSON: %w", err)
	}
	if !doc.Valid() {
		return doc, fmt.Errorf("invalid DID document")
	}

	return doc, nil
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

func signUsingCurrentSeq(did, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, ctx context.Context, didPrivKey secp256k1.PrivKey, doc didtypes.DIDDocument) ([]byte, error) {
	didBase64 := base64.StdEncoding.EncodeToString([]byte(did))
	var didRes didtypes.QueryDIDResponse
	if grpcUrl != "" {
		queryClient := didtypes.NewQueryClient(grpcConn)
		res, err := queryClient.DID(
			ctx,
			&didtypes.QueryDIDRequest{
				DidBase64: didBase64,
			},
		)
		if err != nil {
			return nil, err
		}

		didRes = *res

	} else {
		url := lcdUrl + "/xpla/did/v1beta1/dids/" + didBase64

		out, err := util.CtxHttpClient("GET", url, nil, ctx)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(out, &didRes)
	}

	return didtypes.Sign(&doc, didRes.DidDocumentWithSeq.Sequence, didPrivKey)
}

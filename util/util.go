package util

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Moonyongjung/xpla-private-chain.go/types"
	didtypes "github.com/Moonyongjung/xpla-private-chain/x/did/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/gogo/protobuf/jsonpb"
)

func GetAddrByPrivKey(priv cryptotypes.PrivKey) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromHex(priv.PubKey().Address().String())
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func GasLimitAdjustment(gasUsed uint64, gasAdjustment string) (string, error) {
	gasAdj, err := strconv.ParseFloat(gasAdjustment, 64)
	if err != nil {
		return "", err
	}
	return FromIntToString(int(gasAdj * float64(gasUsed))), nil
}

func GrpcUrlParsing(normalUrl string) string {
	if strings.Contains(normalUrl, "http://") || strings.Contains(normalUrl, "https://") {
		parsedUrl := strings.Split(normalUrl, "://")
		return parsedUrl[1]
	} else {
		return normalUrl
	}
}

func DenomAdd(amount string) string {
	if strings.Contains(amount, types.XplaDenom) {
		return amount
	} else {
		return amount + types.XplaDenom
	}
}

func DenomRemove(amount string) string {
	if strings.Contains(amount, types.XplaDenom) {
		returnAmount := strings.Split(amount, types.XplaDenom)
		return returnAmount[0]
	} else {
		return amount
	}
}

func ConvertEvmChainId(chainId string) (*big.Int, error) {
	conv1 := strings.Split(chainId, "_")
	conv2 := strings.Split(conv1[1], "-")
	returnChainId, err := FromStringToBigInt(conv2[0])
	if err != nil {
		return nil, err
	}
	return returnChainId, nil
}

func Bech32toValidatorAddress(validators []string) ([]sdk.ValAddress, error) {
	vals := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		vals[i] = addr
	}
	return vals, nil
}

func MakeQueryLcdUrl(metadata string) string {
	return "/" + strings.Replace(metadata, "query.proto", "", -1)
}

func MakeQueryLabels(labels ...string) string {
	return strings.Join(labels, "/")
}

func GetDIDDocByQueryClient(did, lcdUrl, grpcUrl string, grpcConn grpc.ClientConn, ctx context.Context) (didtypes.DIDDocumentWithSeq, error) {
	didBase64 := base64.StdEncoding.EncodeToString([]byte(did))
	var didRes didtypes.QueryDIDResponse
	if grpcUrl != "" {
		didQueryclient := didtypes.NewQueryClient(grpcConn)
		didQueryMsg := didtypes.QueryDIDRequest{
			DidBase64: didBase64,
		}
		didRes, err := didQueryclient.DID(ctx, &didQueryMsg)
		if err != nil {
			return didtypes.DIDDocumentWithSeq{}, err
		}

		if didRes.DidDocumentWithSeq.Empty() {
			return didtypes.DIDDocumentWithSeq{}, fmt.Errorf("DID is empty")
		}
		if didRes.DidDocumentWithSeq.Deactivated() {
			return didtypes.DIDDocumentWithSeq{}, fmt.Errorf("DID is deactivate")
		}

	} else {
		url := lcdUrl + "/xpla/did/v1beta1/dids/" + didBase64

		out, err := CtxHttpClient("GET", url, nil, ctx)
		if err != nil {
			return didtypes.DIDDocumentWithSeq{}, err
		}
		jsonpb.Unmarshal(strings.NewReader(string(out)), &didRes)
	}

	return didRes.DidDocumentWithSeq, nil
}

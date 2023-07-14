package evidence

import (
	"encoding/hex"

	"github.com/Moonyongjung/xpriv.go/types"
	"github.com/Moonyongjung/xpriv.go/types/errors"
	"github.com/Moonyongjung/xpriv.go/util"

	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
)

// Parsing - evidence
func parseQueryEvidenceArgs(queryEvidenceMsg types.QueryEvidenceMsg) (evidencetypes.QueryEvidenceRequest, error) {
	decodedHash, err := hex.DecodeString(queryEvidenceMsg.Hash)
	if err != nil {
		return evidencetypes.QueryEvidenceRequest{}, util.LogErr(errors.ErrParse, err)
	}

	return evidencetypes.QueryEvidenceRequest{
		EvidenceHash: decodedHash,
	}, nil
}

package types

type RegisterAnchorAccMsg struct {
	AnchorAccountAddr string
	ValidatorAddr     string
}

type ChangeAnchorAccMsg struct {
	AnchorAccountAddr string
	ValidatorAddr     string
}

type AnchorAccMsg struct {
	ValidatorAddr string
}

type AnchorInfoMsg struct {
	PrivChainHeight string
}

type AnchorBlockMsg struct {
	PrivChainHeight    string
	AnchorContractAddr string
}

type AnchorTxBodyMsg struct {
	PrivChainHeight string
}

type AnchorVerifyMsg struct {
	PrivChainHeight    string
	AnchorContractAddr string
}

type AnchorBalancesMsg struct {
	ValidatorAddr string
}

package types

type CreateDIDMsg struct {
	DIDMnemonic    string
	DIDPassphrase  string
	SaveDIDKeyPath string
	Moniker        string
}

type UpdateDIDMsg struct {
	DID             string
	KeyID           string
	DIDDocumentPath string
	DIDPassphrase   string
	DIDKeyPath      string
}

type DeactivateDIDMsg struct {
	DID           string
	KeyID         string
	DIDPassphrase string
	DIDKeyPath    string
}

type ReplaceDIDMonikerMsg struct {
	DID           string
	KeyId         string
	DIDPassphrase string
	DIDKeyPath    string
	NewMoniker    string
}

type GetDIDMsg struct {
	DIDOrMoniker string
}

type MonikerByDIDMsg struct {
	DID string
}

type DIDByMonikerMsg struct {
	Moniker string
}

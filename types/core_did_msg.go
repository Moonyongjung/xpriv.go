package types

type CreateDIDMsg struct {
	DIDMnemonic    string
	DIDPassphrase  string
	SaveDIDKeyPath string
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

type GetDIDMsg struct {
	DID string
}

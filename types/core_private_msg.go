package types

type InitialAdminMsg struct {
	InitAdminDIDKey string
	DIDPassphrase   string
	DIDKeyPath      string
}

type AddAdminMsg struct {
	NewAdminDIDKey         string
	NewAdminAddress        string
	InitAdminDIDKey        string
	InitAdminDIDPassphrase string
	InitAdminDIDKeyPath    string
}

type ParticipateMsg struct {
	ParticipantDIDKey string
	DIDPassphrase     string
	DIDKeyPath        string
}

type AcceptMsg struct {
	ParticipantDID     string
	AdminDIDKey        string
	AdminDIDPassphrase string
	AdminDIDKeyPath    string
}

type DenyMsg struct {
	ParticipantDID string
	AdminDID       string
}

type ExileMsg struct {
	ParticipantDID string
}

type QuitMsg struct {
	ParticipantDIDKey string
	DIDPassphrase     string
	DIDKeyPath        string
}

type ParticipateStateMsg struct {
	DID string
}

type GenDIDSignMsg struct {
	DIDKey        string
	DIDPassphrase string
	DIDKeyPath    string
}

type IssueVCMsg struct {
	DIDKey        string
	DIDSignBase64 string
}

type GetVPMsg struct {
	DIDKey        string
	DIDSignBase64 string
}

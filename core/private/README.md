# Private module

## Note

- The passphrase is used in the private module is not sended to the chain.
- Parameters such as DID or DIDKey can be replaced to moniker.


## Usage
### (Tx) Initial administrator of the private chain
```go
// Set initial admin of the private chain
initialAdminMsg := types.InitialAdminMsg{
    InitAdminDIDKey: "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    // InitAdminDIDKey: "admindid#key1",
    DIDPassphrase:   "passphrase",
    DIDKeyPath:      "DID/KEY/DIRECTORY",
}

txbytes, err := xplac.InitialAdmin(initialAdminMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Add administrator
```go
// Add administrator
// Init admin can only add another admins
addAdminMsg := types.AddAdminMsg{
    // NewAdminDIDKey: "admindid#key1",
    NewAdminDIDKey:         "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    NewAdminAddress:        "xpriv1k9vu5k7nzq8rfdqu5prvrxlpsrh9y5kxxvla77",
    // InitAdminDIDKey: "initadmindid#key1",
    InitAdminDIDKey:        "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    InitAdminDIDPassphrase: "passphrase",
    InitAdminDIDKeyPath:    "DID/KEY/DIRECTORY",
}

txbytes, err := xplac.AddAdmin(addAdminMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Participate
```go
// Users try to participate to the private chain
participateMsg := types.ParticipateMsg{
    // ParticipantDIDKey: "participantdid#key1",
    ParticipantDIDKey: "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    DIDPassphrase:     "passphrase",
    DIDKeyPath:        "DID/KEY/DIRECTORY",
}

txbytes, err := xplac.Participate(participateMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Accept
```go
// Admin accepts the participant to join private chain
acceptMsg := types.AcceptMsg{
    // ParticipantDIDKey: "participantdid",
    ParticipantDID:     "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty",
    // AdminDIDKey: "admindid#key1",
    AdminDIDKey:        "did:xpla:AGX4EWyvuqA1ivpwbstRu1vSgnTXAqyM3agQvbjstRcp#key1",
    AdminDIDPassphrase: "passphrase",
    AdminDIDKeyPath:    "DID/KEY/DIRECTORY",
}

txbytes, err := xplac.Accept(acceptMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Deny
```go
// Admin denies the participant to join private chain
denyMsg := types.DenyMsg{
    // ParticipantDIDKey: "participantdid",
    ParticipantDID: "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty",
    AdminDID:       "did:xpla:AGX4EWyvuqA1ivpwbstRu1vSgnTXAqyM3agQvbjstRcp",
}
txbytes, err := xplac.Deny(denyMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Exile
```go
// Admin exile the participant of the private chain
exileMsg := types.ExileMsg{
    // ParticipantDIDKey: "participantdid",
    ParticipantDID: "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty",
}

txbytes, err := xplac.Exile(exileMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Quit
```go
// The participant quit of the private chain
quitMsg := types.QuitMsg{
    // ParticipantDIDKey: "participantdid#key1",
    ParticipantDIDKey: "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    DIDPassphrase:     "passphrase",
    DIDKeyPath:        "DID/KEY/DIRECTORY",
}
txbytes, err := xplac.Quit(quitMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Query) Get admin list
```go
// Get Admin list
res, err = xplac.Admin().Query()
```

### (Query) Participate state
```go
// Get participate state
participateStateMsg := types.ParticipateStateMsg{
    // DID: "didmoniker"
    DID: "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty",
}

res, err = xplac.ParticipateState(participateStateMsg).Query()
```

### (Query) Participate sequence
```go
// Get participate sequence
participateSequenceMsg := types.ParticipateSequenceMsg{
    // DID: "didmoniker"
    DID: "did:xpla:HZPozx8eoQrv9vLi26SuZ5gKtKgMduCpectAdYVn3ikP",
}

res, err = xplac.ParticipateSequence(participateSequenceMsg).Query()
```

### (Query) Generate DID signature by using DID sequence
```go
// Generate DID signature
// Sign includes current DID sequence which is increased when user uses DID by sending tx
// Only the owner of the DID can receive DID signature
genDIDSignMsg := types.GenDIDSignMsg{
    // DIDKey: "didmoniker#key1"
    DIDKey:        "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    DIDPassphrase: passphrase,
    DIDKeyPath:    config.CommonParams.DefaultDIDKeypath,
}

res, err = xplac.GenDIDSign(genDIDSignMsg).Query()
```

### (Query) Issue Verfiable Credential
```go
// Issue VC
// "DIDSignBase64" is result by querying "GenDIDSign"
issueVCMsg := types.IssueVCMsg{
    // DIDKey: "didmoniker#key1"
    DIDKey:        "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    DIDSignBase64: "u6+LFHzyluJ2pXyggZPTCp6c6sGVL13BvqGAwjAjUClV+3C6ivSoObUvCvnegMA4BbFkD9nUPP/2Wtsk51xMwA==",
}

res, err = xplac.IssueVC(issueVCMsg).Query()
```

### (Query) Get Verifiable Presentation
```go
// Get the VP
// "DIDSignBase64" is result by querying "GenDIDSign"
getVPMsg := types.GetVPMsg{
    // DIDKey: "didmoniker#key1"
    DIDKey:        "did:xpla:EyAhwxY8KYKNqfZKFoWs9GT1jchFNwrs8MMfeyssmqty#key1",
    DIDSignBase64: "u6+LFHzyluJ2pXyggZPTCp6c6sGVL13BvqGAwjAjUClV+3C6ivSoObUvCvnegMA4BbFkD9nUPP/2Wtsk51xMwA==",
}

res, err = xplac.GetVP(getVPMsg).Query()
```

### (Query) All under reviews state of participants
```go
// all under reviews
res, err := xplac.AllUnderReviews().Query()
```

### (Query) All participants
```go
// all participants
res, err = xplac.AllParticipants().Query()
```


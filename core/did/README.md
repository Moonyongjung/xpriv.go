# DID module

## Note
The passphrase is used in the DID module is not sended to the chain.

## Usage
### (Tx) Create DID
```go
// Set info to create DID
createDIDMsg := types.CreateDIDMsg{
    DIDMnemonic:    "catalog appear keep human ...",
    DIDPassphrase:  "passphrase",
    SaveDIDKeyPath: "/DID/KEY/DIRECTORY",
    // moniker is option
    Moniker:        "didMoniker",
}

txbytes, err := xplac.CreateDID(createDIDMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Update DID
```go
// Update exist DID 
updateDIDMSg := types.UpdateDIDMsg{
    DID:             "did:xpla:mM54wt4G3KBJaJSXEfFv87BuxRHQ2K9WS3m7crFReCL",
    KeyID:           "key1",
    DIDPassphrase:   "passphrase",
    DIDKeyPath:      "DID/KEY/DIRECTORY",
    DIDDocumentPath: "NEW/DID/DOC/DIRECTORY",
}

txbytes, err := xplac.UpdateDID(updateDIDMSg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Deactivate DID
```go
// Deactivate exist DID
deactivateDIDMsg := types.DeactivateDIDMsg{
    DID:           "did:xpla:mM54wt4G3KBJaJSXEfFv87BuxRHQ2K9WS3m7crFReCL",
    KeyID:         "key1",
    DIDPassphrase: "passphrase",
    DIDKeyPath:    "DID/KEY/DIRECTORY",
}

txbytes, err := xplac.DeactivateDID(deactivateDIDMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Replace DID Moniker
```go
// Replace DID moniker
replaceDIDMonikerMsg := types.ReplaceDIDMonikerMsg{
    DID:           "did:xpla:mM54wt4G3KBJaJSXEfFv87BuxRHQ2K9WS3m7crFReCL",
    KeyId:         "key1",
    DIDPassphrase: "passphrase",
    DIDKeyPath:    "DID/KEY/DIRECTORY",
    NewMoniker:    "newDidMoniker",
}

txbytes, err := xplac.ReplaceDIDMoniker(replaceDIDMonikerMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Query) Get DID info
```go
// Get DID info
getDIDMsg := types.GetDIDMsg{
    DID: "did:xpla:mM54wt4G3KBJaJSXEfFv87BuxRHQ2K9WS3m7crFReCL",
}

res, err := xplac.GetDID(getDIDMsg).Query()
```

### (Query) Moniker by DID
```go
// Get moniker by DID
monikerByDIDMsg := types.MonikerByDIDMsg{
    DID: "did:xpla:mM54wt4G3KBJaJSXEfFv87BuxRHQ2K9WS3m7crFReCL",
}

res, err := xplac.MonikerByDID(monikerByDIDMsg).Query()
```

### (Query) DID by moniker
```go
// Get DID by moniker
didByMonikerMsg := types.DIDByMonikerMsg{
    Moniker: "didMoniker",
}

res, err := xplac.DIDByMoniker(didByMonikerMsg).Query()
```


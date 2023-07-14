# Private module
## Usage
### (Tx) Create DID
```go
// Set info to create DID
createDIDMsg := types.CreateDIDMsg{
    DIDMnemonic:    "catalog appear keep human ...",
    DIDPassphrase:  "passphrase",
    SaveDIDKeyPath: "/DID/KEY/DIRECTORY",
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

### (Query) Get DID info
```go
// Get DID info
getDIDMsg := types.GetDIDMsg{
    DID: "did:xpla:mM54wt4G3KBJaJSXEfFv87BuxRHQ2K9WS3m7crFReCL",
}

res, err = xplac.GetDID(getDIDMsg).Query()
```



# Anchor module

## Usage

### (Tx) Register anchor account
```go
// register anchor account with validator of the private chain
// the validator send anchor tx to public chain by using mapped anchor account
registerAnchorAccMsg := types.RegisterAnchorAccMsg{
    AnchorAccountAddr: "xpla148f0qtq5zrujl8ed02lq4e9x2v7zavxzses5hc",
    ValidatorAddr:     "xprivvaloper1chmr83fzt5gyz6k72pw9czut8lg6crg6kpz95z",
}

txbytes, err := xplac.RegisterAnchorAcc(registerAnchorAccMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Tx) Change anchor account
```go
// change anchor account of the validator
changeAnchorAccMsg := types.ChangeAnchorAccMsg{
    AnchorAccountAddr: "xpla148f0qtq5zrujl8ed02lq4e9x2v7zavxzses5hc",
    ValidatorAddr:     "xprivvaloper1chmr83fzt5gyz6k72pw9czut8lg6crg6kpz95z",
}

txbytese, err := xplac.ChangeAnchorAcc(changeAnchorAccMsg).CreateAndSignTx()
res, err := xplac.Broadcast(txbytes)
```

### (Query) Anchor account 
```go
// query anchor account of the validator
anchorAccMsg := types.AnchorAccMsg{
    ValidatorAddr: "xprivvaloper1jmf9krhvv9l0ds6ughst5ffd30dvmjf57y9hdd",
}

res, err = xplac.AnchorAcc(anchorAccMsg).Query()
```

### (Query) All aggregated blocks
```go
// query all aggregated blocks for anchoring
res, err = xplac.AllAggregatedBlocks().Query()
```

### (Query) Anchor info
```go
// query anchoring info
anchorInfoMsg := types.AnchorInfoMsg{
    PrivChainHeight: "20",
}?

res, err = xplac.AnchorInfo(anchorInfoMsg).Query()
```

### (Query) Anchor block
```go
// query anchoring block info in the anchor contract
anchorBlockMsg := types.AnchorBlockMsg{
    PrivChainHeight:    "20",
    AnchorContractAddr: "xpla1fyr2mptjswz4w6xmgnpgm93x0q4s4wdl6srv3rtz3utc4f6fmxeqajvryg",
}

res, err = xplac.AnchorBlock(anchorBlockMsg).Query()
```

### (Query) Anchor tx body
```go
// query anchoring transaction in the public chain
anchorTxBodyMsg := types.AnchorTxBodyMsg{
    PrivChainHeight: "20",
}

res, err = xplac.AnchorTxBody(anchorTxBodyMsg).Query()
```

### (Query) Verify
```go
// check the consistency of the block in the private chain
anchorVerifyMsg := types.AnchorVerifyMsg{
    PrivChainHeight:    "20",
    AnchorContractAddr: "xpla1fyr2mptjswz4w6xmgnpgm93x0q4s4wdl6srv3rtz3utc4f6fmxeqajvryg",
}

res, err = xplac.AnchorVerify(anchorVerifyMsg).Query()
```

### (Query) Anchor balances
```go
// query balances of the anchot account in the public chain
anchorBalancesMsg := types.AnchorBalancesMsg{
    ValidatorAddr: "xprivvaloper1jmf9krhvv9l0ds6ughst5ffd30dvmjf57y9hdd",
}

res, err = xplac.AnchorBalances(anchorBalancesMsg).Query()
```

### (Query) Params
```go
// Get params of anchor module
res, err := xplac.AnchorParams().Query()
```
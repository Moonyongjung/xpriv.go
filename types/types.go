package types

const (
	// New mnemonic entropy size
	DefaultEntropySize = 256
	// Xpla base denomination
	XplaDenom = "axpriv"
	// Xpla default key algorithm name
	DefaultXplaKeyAlgo = "eth_secp256k1"
	// Xpla tool default name
	XplaToolDefaultName = "xpriv"

	// query method type
	QueryGrpc = 1
	QueryLcd  = 2

	DefaultGasLimit      = "250000"
	DefaultGasPrice      = "850000000000"
	DefaultGasAdjustment = "1.75"
	DefaultAccNum        = 0
	DefaultAccSeq        = 0
)

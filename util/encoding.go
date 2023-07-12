package util

import (
	"github.com/Moonyongjung/xpla-private-chain/app"
	"github.com/Moonyongjung/xpla-private-chain/app/params"
)

func MakeEncodingConfig() params.EncodingConfig {
	return app.MakeEncodingConfig()
}

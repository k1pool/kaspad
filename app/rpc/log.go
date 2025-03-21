package rpc

import (
	"github.com/k1pool/kaspad/infrastructure/logger"
	"github.com/k1pool/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)

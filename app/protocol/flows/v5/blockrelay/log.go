package blockrelay

import (
	"github.com/k1pool/kaspad/infrastructure/logger"
	"github.com/k1pool/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)

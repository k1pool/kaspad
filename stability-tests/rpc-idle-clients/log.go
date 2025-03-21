package main

import (
	"github.com/k1pool/kaspad/infrastructure/logger"
	"github.com/k1pool/kaspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)

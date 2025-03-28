package main

import (
	"github.com/k1pool/kaspad/infrastructure/logger"
	"github.com/k1pool/kaspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)

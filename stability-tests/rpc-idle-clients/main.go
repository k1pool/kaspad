package main

import (
	"time"

	"github.com/k1pool/kaspad/stability-tests/common"
	"github.com/k1pool/kaspad/stability-tests/common/rpc"
	"github.com/k1pool/kaspad/util/panics"
	"github.com/k1pool/kaspad/util/profiling"
	"github.com/pkg/errors"
)

func main() {
	defer panics.HandlePanic(log, "rpc-idle-clients-main", nil)
	err := parseConfig()
	if err != nil {
		panic(errors.Wrap(err, "error parsing configuration"))
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())

	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	numRPCClients := cfg.NumClients
	clients := make([]*rpc.Client, numRPCClients)
	for i := uint32(0); i < numRPCClients; i++ {
		rpcClient, err := rpc.ConnectToRPC(&cfg.Config, cfg.NetParams())
		if err != nil {
			panic(errors.Wrap(err, "error connecting to RPC server"))
		}
		clients[i] = rpcClient
	}

	const testDuration = 30 * time.Second
	select {
	case <-time.After(testDuration):
	}
	for _, client := range clients {
		client.Close()
	}
}

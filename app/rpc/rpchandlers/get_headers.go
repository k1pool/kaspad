package rpchandlers

import (
	"github.com/k1pool/kaspad/app/appmessage"
	"github.com/k1pool/kaspad/app/rpc/rpccontext"
	"github.com/k1pool/kaspad/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}

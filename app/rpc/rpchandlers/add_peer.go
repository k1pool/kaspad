package rpchandlers

import (
	"github.com/k1pool/kaspad/app/appmessage"
	"github.com/k1pool/kaspad/app/rpc/rpccontext"
	"github.com/k1pool/kaspad/infrastructure/network/netadapter/router"
	"github.com/k1pool/kaspad/util/network"
)

// HandleAddPeer handles the respectively named RPC command
func HandleAddPeer(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if context.Config.SafeRPC {
		log.Warn("AddPeer RPC command called while node in safe RPC mode -- ignoring.")
		response := appmessage.NewAddPeerResponseMessage()
		response.Error =
			appmessage.RPCErrorf("AddPeer RPC command called while node in safe RPC mode")
		return response, nil
	}

	AddPeerRequest := request.(*appmessage.AddPeerRequestMessage)
	address, err := network.NormalizeAddress(AddPeerRequest.Address, context.Config.ActiveNetParams.DefaultPort)
	if err != nil {
		errorMessage := &appmessage.AddPeerResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Could not parse address: %s", err)
		return errorMessage, nil
	}

	context.ConnectionManager.AddConnectionRequest(address, AddPeerRequest.IsPermanent)

	response := appmessage.NewAddPeerResponseMessage()
	return response, nil
}

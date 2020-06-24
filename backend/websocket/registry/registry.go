package registry

import (
	"sync"

	"github.com/bseto/arcade/backend/websocket/identifier"
)

// Registry defines an interface in which `Registry`'s should provide
type Registry interface {
	Register(send chan []byte, clientID identifier.Client)
	Unregister(clientID identifier.Client)

	//SendToSameHub(clientID identifier.Client, message []byte)
	//SendToCaller(clientID identifier.Client, message []byte)
}

// RegistryProvider will
type RegistryProvider struct {
	lookupLock sync.RWMutex
	//lookupMap map[id.Client.HubName]map[id.Client.ClientUUID](chan []byte)
}

func GetRegistryProvider() *RegistryProvider {
	return &RegistryProvider{}
}

//func GetRegistryProvider() RegistryProvider {
//return RegistryProvider {
//lookupMap: make(
//map[id.Client.HubName]map[id.Client.ClientUUID](chan []byte),
//),
//}
//}

func (r *RegistryProvider) Register(
	send chan []byte,
	clientID identifier.Client,
) {
	//stub
	return
}

func (r *RegistryProvider) Unregister(
	clientID identifier.Client,
) {
	// stub
	return
}

func (r *RegistryProvider) SendToSameHub(
	clientID identifier.Client,
	message []byte,
) {
	// stub
	return

}
func (r *RegistryProvider) SendToCaller(
	clientID identifier.Client,
	message []byte,
) {
	//stub
	return
}

// package registry will be in charge of handling the send channel.
// this package should have no logic in regards to handling hub
// related things. The hub is only used in this package for identifying
// the channel
package registry

import (
	"sync"

	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
)

// Registry defines an interface in which `Registry`'s should provide
type Registry interface {
	Register(send chan []byte, clientID identifier.Client)
	Unregister(clientID identifier.Client) (registryEmpty bool)

	SendToSameHub(clientID identifier.Client, message []byte)
	SendToCaller(clientID identifier.Client, message []byte)
}

// RegistryProvider will provide the actual registry functionality
type RegistryProvider struct {
	lookupLock sync.RWMutex
	lookupMap  map[identifier.ClientUUIDStruct](chan []byte)
}

func GetRegistryProvider() *RegistryProvider {
	return &RegistryProvider{
		lookupMap: make(
			map[identifier.ClientUUIDStruct](chan []byte),
		),
	}
}

// Register should take the send chan and fill in the lookupMap
// This function should be threadsafe
func (r *RegistryProvider) Register(
	send chan []byte,
	clientID identifier.Client,
) {
	r.lookupLock.Lock()
	defer r.lookupLock.Unlock()
	log.Infof("registering: %v", clientID)

	_, ok := r.lookupMap[clientID.ClientUUID]

	if ok == true {
		log.Errorf(
			"the client already exists: %v : %v",
			clientID,
		)
		return
	}

	r.lookupMap[clientID.ClientUUID] = send
}

func (r *RegistryProvider) Unregister(
	clientID identifier.Client,
) (registryEmpty bool) {
	r.lookupLock.Lock()
	defer r.lookupLock.Unlock()
	log.Infof("unregistering: %v", clientID)

	delete(r.lookupMap, clientID.ClientUUID)
	if len(r.lookupMap) == 0 {
		registryEmpty = true
	}
	return
}

func (r *RegistryProvider) SendToSameHub(
	clientID identifier.Client,
	message []byte,
) {
	r.lookupLock.Lock()
	defer r.lookupLock.Unlock()

	for _, clientChannel := range r.lookupMap {
		clientChannel <- message
	}
	return

}

func (r *RegistryProvider) SendToCaller(
	clientID identifier.Client,
	message []byte,
) {
	r.lookupLock.Lock()
	defer r.lookupLock.Unlock()

	sendChannel, ok := r.lookupMap[clientID.ClientUUID]
	if ok != true {
		log.Errorf("could not find channel for ID: %v", clientID)
		return
	}

	sendChannel <- message
}

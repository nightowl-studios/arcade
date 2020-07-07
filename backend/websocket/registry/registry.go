// package registry will be used for two purposes
// one purpose is to indirectly get access to the send channel
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

	GetClientSlice() []*identifier.UserDetails
	GetClientUserDetail(
		clientID identifier.Client,
	) *identifier.UserDetails

	SendToSameHub(clientID identifier.Client, message []byte)
	SendToCaller(clientID identifier.Client, message []byte)
	SendToSameHubExceptCaller(clientID identifier.Client, message []byte)
}

// RegistryProvider will provide the actual registry functionality
type RegistryProvider struct {
	lookupLock sync.RWMutex
	lookupMap  map[identifier.ClientUUIDStruct]UserDetails
}

type UserDetails struct {
	send        (chan []byte)
	userDetails *identifier.UserDetails
}

func GetRegistryProvider() *RegistryProvider {
	return &RegistryProvider{
		lookupMap: make(
			map[identifier.ClientUUIDStruct]UserDetails,
		),
	}
}

func (r *RegistryProvider) GetClientUserDetail(
	clientID identifier.Client,
) *identifier.UserDetails {
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	return r.lookupMap[clientID.ClientUUID].userDetails
}

func (r *RegistryProvider) GetClientSlice() []*identifier.UserDetails {
	var clients []*identifier.UserDetails
	for _, userDetails := range r.lookupMap {
		clients = append(clients, userDetails.userDetails)
	}
	return clients
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

	r.lookupMap[clientID.ClientUUID] = UserDetails{
		send: send,
		userDetails: &identifier.UserDetails{
			ClientUUID: clientID.ClientUUID,
		},
	}
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
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	for _, clientDetails := range r.lookupMap {
		clientDetails.send <- message
	}
	return

}

func (r *RegistryProvider) SendToCaller(
	clientID identifier.Client,
	message []byte,
) {
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	clientDetails, ok := r.lookupMap[clientID.ClientUUID]
	if ok != true {
		log.Errorf("could not find channel for ID: %v", clientID)
		return
	}

	clientDetails.send <- message
}

// SendToSameHubExceptCaller will send to everyone in the hub, except for the caller
func (r *RegistryProvider) SendToSameHubExceptCaller(
	clientID identifier.Client,
	message []byte,
) {
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	for _, clientDetails := range r.lookupMap {
		if clientDetails.userDetails.ClientUUID == clientID.ClientUUID {
			// don't send if the clientUUID match the caller
			continue
		}
		clientDetails.send <- message
	}

	return
}

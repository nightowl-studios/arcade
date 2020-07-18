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
	Register(send chan []byte, clientID identifier.ClientUUIDStruct)
	Unregister(clientID identifier.ClientUUIDStruct) (registryEmpty bool)

	GetClientSlice() []*identifier.UserDetails
	GetClientUserDetail(
		clientID identifier.ClientUUIDStruct,
	) *identifier.UserDetails

	SendToSameHub(clientID identifier.ClientUUIDStruct, message []byte)
	SendToCaller(clientID identifier.ClientUUIDStruct, message []byte)
	SendToSameHubExceptCaller(clientID identifier.ClientUUIDStruct, message []byte)
}

// RegistryProvider will provide the actual registry functionality
type RegistryProvider struct {
	lookupLock   sync.RWMutex
	lookupMap    map[identifier.ClientUUIDStruct]UserDetails
	unregistered map[identifier.ClientUUIDStruct]UserDetails
}

type UserDetails struct {
	send        (chan []byte)
	userDetails *identifier.UserDetails
}

func GetRegistryProvider() *RegistryProvider {
	return &RegistryProvider{
		lookupMap:    make(map[identifier.ClientUUIDStruct]UserDetails),
		unregistered: make(map[identifier.ClientUUIDStruct]UserDetails),
	}
}

func (r *RegistryProvider) GetClientUserDetail(
	clientID identifier.ClientUUIDStruct,
) *identifier.UserDetails {
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	return r.lookupMap[clientID].userDetails
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
// This function will look for the client in the unregistered map
// if it exists, it means this user is trying to reconnect
func (r *RegistryProvider) Register(
	send chan []byte,
	clientID identifier.ClientUUIDStruct,
) {
	r.lookupLock.Lock()
	defer r.lookupLock.Unlock()
	log.Infof("registering: %v", clientID)

	_, ok := r.lookupMap[clientID]

	if ok == true {
		log.Errorf(
			"the client already exists: %v : %v",
			clientID,
		)
		return
	}

	userDetails, ok := r.unregistered[clientID]
	if ok == true {
		// someone is reconnecting. Move them from the unregistered map to the
		// lookupMap
		userDetails.send = send
		delete(r.unregistered, userDetails.userDetails.ClientUUID)
	} else {
		// create a new user
		userDetails = UserDetails{
			send: send,
			userDetails: &identifier.UserDetails{
				ClientUUID: clientID,
			},
		}
	}

	r.lookupMap[clientID] = userDetails
}

// Unregister will take the client and move it to the unregistered map
// if the main lookupMap is empty, it'll return true
func (r *RegistryProvider) Unregister(
	clientID identifier.ClientUUIDStruct,
) (registryEmpty bool) {
	r.lookupLock.Lock()
	defer r.lookupLock.Unlock()
	log.Infof("unregistering: %v", clientID)

	userDetails, ok := r.lookupMap[clientID]
	if ok {
		userDetails.send = nil
		r.unregistered[clientID] = userDetails
	}

	delete(r.lookupMap, clientID)
	if len(r.lookupMap) == 0 {
		registryEmpty = true
	}
	return
}

func (r *RegistryProvider) SendToSameHub(
	clientID identifier.ClientUUIDStruct,
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
	clientID identifier.ClientUUIDStruct,
	message []byte,
) {
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	clientDetails, ok := r.lookupMap[clientID]
	if ok != true {
		log.Errorf("could not find channel for ID: %v", clientID)
		return
	}

	clientDetails.send <- message
}

// SendToSameHubExceptCaller will send to everyone in the hub, except for the caller
func (r *RegistryProvider) SendToSameHubExceptCaller(
	clientID identifier.ClientUUIDStruct,
	message []byte,
) {
	r.lookupLock.RLock()
	defer r.lookupLock.RUnlock()

	for _, clientDetails := range r.lookupMap {
		if clientDetails.userDetails.ClientUUID == clientID {
			// don't send if the clientUUID match the caller
			continue
		}
		clientDetails.send <- message
	}

	return
}

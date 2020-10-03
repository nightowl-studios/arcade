package registry

import (
	"testing"

	"github.com/bseto/arcade/backend/websocket/identifier"
)

var TestSequentialRegisterProvider = []struct {
	testName       string
	inputClientIDs []identifier.Client
	// no defined expected because all inputs are expected as output
}{
	{
		testName: "single client",
		inputClientIDs: []identifier.Client{
			identifier.Client{
				HubName:    identifier.HubNameStruct{"1"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
		},
	},
	{
		testName: "two clients",
		inputClientIDs: []identifier.Client{
			identifier.Client{
				HubName:    identifier.HubNameStruct{"1"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
		},
	},
	{
		testName: "multiple clients",
		inputClientIDs: []identifier.Client{
			identifier.Client{
				HubName:    identifier.HubNameStruct{"1a"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"1a"},
				ClientUUID: identifier.ClientUUIDStruct{"2"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"1a"},
				ClientUUID: identifier.ClientUUIDStruct{"3ab"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2"},
				ClientUUID: identifier.ClientUUIDStruct{"a"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2"},
				ClientUUID: identifier.ClientUUIDStruct{"b"},
			},
		},
	},
}

func TestSequentialRegister(t *testing.T) {
	// dummyChannel for use for test
	dummyChannel := make(chan []byte)
	defer func() {
		// cleanup the dummyChannel at the end of the test
		close(dummyChannel)
	}()

	for _, testVal := range TestSequentialRegisterProvider {
		t.Run(testVal.testName, func(t *testing.T) {
			reg := GetRegistryProvider()

			// fill up the registry with the inputClientID
			for _, inputClientID := range testVal.inputClientIDs {
				reg.Register(dummyChannel, inputClientID.ClientUUID)
			}

			// now check to see if all the clients are in the map
			for _, inputClientID := range testVal.inputClientIDs {
				_, ok := reg.lookupMap[inputClientID.ClientUUID]
				if !ok {
					t.Errorf(
						"unable to find: hubID: %v, clientID: %v",
						inputClientID.HubName.HubName,
						inputClientID.ClientUUID.UUID,
					)
				}
			}
		})
	}
}

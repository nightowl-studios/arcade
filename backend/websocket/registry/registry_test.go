package registry

import (
	"sync"
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
				reg.Register(dummyChannel, inputClientID)
			}

			// now check to see if all the clients are in the map
			for _, inputClientID := range testVal.inputClientIDs {
				_, ok := reg.lookupMap[inputClientID.HubName][inputClientID.ClientUUID]
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

// TestThreadedRegister will divide the registration up into 2 threads to test
// threadsafety. The provider is defined below because of how large it is
func TestThreadedRegister(t *testing.T) {
	// dummyChannel for use for test
	dummyChannel := make(chan []byte)
	defer func() {
		// cleanup the dummyChannel at the end of the test
		close(dummyChannel)
	}()

	for _, testVal := range TestThreadedRegisterProvider {
		t.Run(testVal.testName, func(t *testing.T) {
			reg := GetRegistryProvider()
			var wg sync.WaitGroup

			registerThread := func(wg *sync.WaitGroup, clientIDs []identifier.Client) {
				defer wg.Done()
				for _, inputClientID := range clientIDs {
					reg.Register(dummyChannel, inputClientID)
				}
			}

			IDLength := len(testVal.inputClientIDs)
			wg.Add(1)
			go registerThread(&wg, testVal.inputClientIDs[0:IDLength/2])
			wg.Add(1)
			go registerThread(&wg, testVal.inputClientIDs[IDLength/2:])

			// Will wait for the above two threads to finish
			wg.Wait()

			// now check to see if all the clients are in the map
			for _, inputClientID := range testVal.inputClientIDs {
				_, ok := reg.lookupMap[inputClientID.HubName][inputClientID.ClientUUID]
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

var TestThreadedRegisterProvider = []struct {
	testName       string
	inputClientIDs []identifier.Client
	// no defined expected because all inputs are expected as output
}{
	{
		testName: "hard mode multiple clients",
		inputClientIDs: []identifier.Client{
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2a"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2a"},
				ClientUUID: identifier.ClientUUIDStruct{"2"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"2a"},
				ClientUUID: identifier.ClientUUIDStruct{"3ab"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"3"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"3"},
				ClientUUID: identifier.ClientUUIDStruct{"a"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"3"},
				ClientUUID: identifier.ClientUUIDStruct{"b"},
			},
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
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"1"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"2"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"3"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"4"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"5"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"6"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"7"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"8"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"9"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"10"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"11"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"12"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"13"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"14"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"15"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"16"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"17"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"18"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"19"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"20"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"21"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"22"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"23"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"24"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"25"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"26"},
			},
			identifier.Client{
				HubName:    identifier.HubNameStruct{"lotsofusers"},
				ClientUUID: identifier.ClientUUIDStruct{"27"},
			},
		},
	},
}

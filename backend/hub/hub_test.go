package hub

import (
	"crypto/rand"
	"testing"

	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/dgrijalva/jwt-go"
)

var TestParseTokenProvider = []struct {
	Name           string
	Client         identifier.Client
	OverrideSecret bool
	ExpectedErr    error
}{
	{
		Name: "Simple Test Case",
		Client: identifier.Client{
			ClientUUID: identifier.ClientUUIDStruct{"ABCD"},
			HubName:    identifier.HubNameStruct{"123"},
		},
		ExpectedErr:    nil,
		OverrideSecret: false,
	},
	{
		Name: "Invalid Token",
		Client: identifier.Client{
			ClientUUID: identifier.ClientUUIDStruct{"ABCD"},
			HubName:    identifier.HubNameStruct{"123"},
		},
		ExpectedErr:    jwt.ErrSignatureInvalid,
		OverrideSecret: true,
	},
}

func TestParseToken(t *testing.T) {
	secret := make([]byte, 60)
	_, err := rand.Read(secret)
	if err != nil {
		t.Fatalf("unable to setup secret: %v", err)
	}
	for _, testVal := range TestParseTokenProvider {

		t.Run(testVal.Name, func(t *testing.T) {
			claims := JSONWebToken{
				Client: testVal.Client,
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedString, err := token.SignedString(secret)
			tokenMessage := JSONWebTokenMessage{
				Token:         signedString,
				ContainsToken: true,
			}

			var parseSecret []byte
			if testVal.OverrideSecret == true {
				parseSecret = []byte("dummypassword")
			} else {
				parseSecret = secret
			}
			client, err := ParseToken(tokenMessage, parseSecret)

			//if errors.Is(err, testVal.ExpectedErr) != true {
			if err != nil {
				if err.Error() != testVal.ExpectedErr.Error() {
					t.Errorf("expected err: %v, got: %v", testVal.ExpectedErr, err)
				}
				// No need to run any further
				return
			}

			if testVal.Client.ClientUUID.UUID != client.ClientUUID.UUID {
				t.Errorf(
					"clientUUID mismatch. Expected: %v, got: %v",
					testVal.Client.ClientUUID.UUID,
					client.ClientUUID.UUID,
				)
			}

		})

	}
}
// package identifier provides functionality to uniquely identify websocket
// clients
package identifier

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Client struct should contain all information about the client connected via
// websocket - These details do not change. If the user connects to a different
// websocket connection, they will be given a new Client detail
// ClientUUID is a completely unique identifier
// HubName is the name of the hub they are connected to
type Client struct {
	ClientUUID ClientUUIDStruct `json:"clientUUID"`
	HubName    HubNameStruct    `json:"hubName"`

	// this Client struct may also include things like a jwt token or something
	// later so that we can code in some reconnect functionality
}

// UserDetails should only be held by the registry
// These are details that can change and follow the "user"
// with the exception of the ClientUUID which is needed so
// the frontend can ID this user
type UserDetails struct {
	ClientUUID   ClientUUIDStruct `json:"clientUUID"`
	NickNameLock sync.RWMutex     `json:"-"`
	NickName     string           `json:"nickname"`
	JoinOrder    int              `json:"joinOrder"`
}

func (u *UserDetails) GetNickName() string {
	u.NickNameLock.RLock()
	defer u.NickNameLock.RUnlock()
	return u.NickName
}

func (u *UserDetails) ChangeNickName(newName string) {
	u.NickNameLock.Lock()
	defer u.NickNameLock.Unlock()
	u.NickName = newName
}

type ClientUUIDStruct struct {
	UUID string
}

type HubNameStruct struct {
	HubName string
}

func CreateClientUUID() string {
	return uuid.NewV4().String()
}

// CreateHubName will return a randomly generated string suitable for a
// hubname. The `nameLength` will control the length of the output
func CreateHubName(nameLength int) string {
	return RandStringBytesMaskImprSrcSB(nameLength)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandStringBytesMaskImprSrcSB is a copy-paste function from SO
// not changing the function name so it's easier to find and credit
// to this link: https://stackoverflow.com/a/31832326
func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return strings.ToUpper(sb.String())
}

package websocket

import (
	"net/http"
	"time"

	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/gorilla/websocket"
)

// WebsocketClient should be created and used to upgrade the http request
// to create a websocket connection
type WebsocketClient interface {
	InsertHandler(handler WebsocketHandler)

	// Upgrade will upgrade the http request. It is also in the Upgrade
	// function that the WebsocketHandler handleAuthentication function will
	// be called
	Upgrade(w http.ResponseWriter, r *http.Request) error

	// Close will close the websocket connection and stop any internal processes
	Close()
}

type WebsocketHandler interface {
	HandleMessage(
		messageType int,
		message []byte,
		clientID identifier.Client,
		err error,
	)

	HandleAuthentication(
		w http.ResponseWriter,
		r *http.Request,
		conn *websocket.Conn,
		send chan []byte,
	) (identifier.Client, error)

	SignalClose()
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var defaultUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins to initialize connections
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	handler WebsocketHandler
}

func GetClient(handler WebsocketHandler) Client {
	return Client{
		handler: handler,
	}
}

func (c *Client) InsertHandler(handler WebsocketHandler) {
	// stub
	return
}

// Upgrade will upgrade the http request. It is also in the Upgrade
// function that the WebsocketHandler handleAuthentication function will
// be called
func (c *Client) Upgrade(
	w http.ResponseWriter,
	r *http.Request,
) (err error) {
	// stub
	return
}

// Close will close the websocket connection and stop any internal processes
func (c *Client) Close() {
	// stub
	return
}

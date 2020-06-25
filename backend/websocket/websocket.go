package websocket

import (
	"bytes"
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
	conn *websocket.Conn
	send chan []byte

	handler WebsocketHandler

	clientID identifier.Client
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

// readPump is a function in charge of reading from the websocket. No other
// process should read from the connection when the readPump is running.
// this function is meant to be called as a go routine
func (c *Client) readPump() {
	defer func() {
		c.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		c.handler.HandleMessage(
			messageType,
			message,
			c.clientID,
			err,
		)
	}
}

// writePump is a function in charge of writing to the websocket. If any messages
// need to be sent to the websocket, it needs to be done through this function
// This function is meant to be called as a go routine
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.Close()

		// cleanup local variables
		ticker.Stop()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			c.sendMessages(message)
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				// ask the frontend client to close the channel
				// this should trigger the readPump to unregister and close
				// which will trigger the `ok <- c.send` to be false
				c.Close()
				continue
			}
		}
	}
}

func (c *Client) sendMessages(message []byte) {
	// stub
	return
}

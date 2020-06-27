package websocket

import (
	"bytes"
	"net/http"
	"time"

	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/gorilla/websocket"
)

// WebsocketClient should be created and used to upgrade the http request
// to create a websocket connection
type WebsocketClient interface {
	// Upgrade will upgrade the http request. It is also in the Upgrade
	// function that the WebsocketHandler handleAuthentication function will
	// be called
	Upgrade(w http.ResponseWriter, r *http.Request) error

	// Close will close the websocket connection and stop any internal processes
	Close()
}

// WebsocketHandler allows for different API handlers to be used within a websocket
// context. As long as these functions are defined, the WebsocketHandler will
// have access to reading (via HandleMessage) and writing (via `send chan []byte`)
// to the websocket.
// Other functionality such as database references, or translations or anything
// should be stored within the struct that implements this interface
type WebsocketHandler interface {

	// HandleMessage is where the WebsocketHandler should route the messages.
	// Essentially, this function is to become the router for this Game / or API
	HandleMessage(
		messageType int,
		message []byte,
		clientID identifier.Client,
		err error,
	)

	// HandleAuthentication is called during the websocket upgrade.
	// If there is any authentication handshake, it should be done here.
	// The `writePump` will be started prior to calling this function,
	// so sending messages via the send channel is the proper way to
	// send outgoing messages. As for incoming messages, the `readPump` will
	// not be started (to avoid calling the HandleMessage function prior to
	// proper authentication) so reading the incoming messages has to be done
	// manually
	// The WebsocketClient will abort if this function returns a non nil error
	HandleAuthentication(
		w http.ResponseWriter,
		r *http.Request,
		conn *websocket.Conn,
		send chan []byte,
	) (identifier.Client, error)

	// SignalClose is called by the WebsocketClient when the websocket
	// client is about to close.
	// Make sure to do all cleanup in this SignalClose - ie, cleaning up
	// the send channel
	SignalClose(caller identifier.Client)

	// Upgrader allows the WebsocketHandler to decide the properties of the
	// websocket upgrade
	Upgrader() *websocket.Upgrader
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

// Client is the struct that will implement the WebsocketClient interface
type Client struct {
	conn *websocket.Conn
	send chan []byte

	handler WebsocketHandler

	clientID identifier.Client
}

// Close will close the websocket connection and stop any internal processes
func (c *Client) Close() {
	// stub
	return
}

// upgrade will upgrade the http request. It is also in the Upgrade
// function that the WebsocketHandler handleAuthentication function will
// be called

// Upgrade will upgrade the http request to a websocket.
// The Upgrade function will use the Upgrader from the handler
// and call the handler's Authentication function
func (c *Client) Upgrade(
	w http.ResponseWriter,
	r *http.Request,
) error {
	conn, err := c.handler.Upgrader().Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	c.conn = conn

	go c.writePump()

	id, err := c.handler.HandleAuthentication(w, r, conn, c.send)
	if err != nil {
		log.Errorf("unable to handle auth, exiting: %v", err)
		return err
	}
	c.clientID = id

	go c.readPump()

	return err
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
			log.Errorf("unable to read from websocket, closing socket: %v", err)
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

// Utility Functions

func GetClient(handler WebsocketHandler) Client {
	return Client{
		send:    make(chan []byte),
		handler: handler,
	}
}

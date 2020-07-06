package websocket

import (
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/bseto/arcade/backend/hub"
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

	Dial(hostname string)

	// Close will close the websocket connection and stop any internal processes
	Close()

	RegisterCloseListener(listener WebsocketCloseListener)
}

// WebsocketCloseListener are notified when the websocket is closed
type WebsocketCloseListener interface {
	WebsocketClose(clientID identifier.Client)
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
	connWriteLock sync.Mutex
	conn          *websocket.Conn
	send          chan []byte

	hub hub.Hub

	clientID identifier.Client

	// closeState tracks whether or not this websocket client is
	// already cleaning up.
	// This is needed since either one of the read or write pump can trigger
	// a close
	closeState     bool
	closeStateLock sync.RWMutex

	// listeners
	closeListenerLock sync.RWMutex
	closeListener     []WebsocketCloseListener
}

func GetClient(
	hubInstance hub.Hub,
) Client {
	return Client{
		send:       make(chan []byte),
		hub:        hubInstance,
		closeState: false,
	}
}

// GetDialClient will return a client where it's meant for dialing
// a websocket server. Not meant to be used for anything other than
// testing purposes at the moment
// This function will automatically start the writePump, but will not start
// the read pump
func DialClient(hostname string) (Client, chan []byte, *websocket.Conn, error) {
	sendChan := make(chan []byte)

	retClient := Client{
		send:       sendChan,
		closeState: false,
	}

	conn, _, err := websocket.DefaultDialer.Dial(hostname, nil)
	if err != nil {
		log.Errorf("unable to dial connection: %v", err)
		return retClient, sendChan, conn, err
	}
	retClient.conn = conn
	go retClient.writePump()

	return retClient, sendChan, conn, nil
}

func (c *Client) RegisterCloseListener(listener WebsocketCloseListener) {
	c.closeListenerLock.Lock()
	defer c.closeListenerLock.Unlock()
	c.closeListener = append(c.closeListener, listener)
}

func (c *Client) NotifyClose() {
	c.closeListenerLock.RLock()
	defer c.closeListenerLock.RUnlock()
	for _, listener := range c.closeListener {
		listener.WebsocketClose(c.clientID)
	}
}

// Close will close the websocket connection and stop any internal processes
func (c *Client) Close() {
	c.closeStateLock.Lock()
	if c.closeState == true {
		// This websocket has already triggered a Close cleanup - exiting
		c.closeStateLock.Unlock()
		return
	}
	c.closeState = true
	c.closeStateLock.Unlock()

	// stop anyone from trying to use the websocket before we actually
	// close the websocket
	c.NotifyClose()

	c.connWriteLock.Lock()
	defer c.connWriteLock.Unlock()

	c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
		websocket.CloseNormalClosure,
		"",
	))

	time.Sleep(time.Second * 1)
	c.conn.Close()

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
	conn, err := c.hub.Upgrader().Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	c.conn = conn

	go c.writePump()

	id, err := c.hub.HandleAuthentication(w, r, conn, c.send)
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
			log.Infof("websocket has been closed from client: %v, %v", c.clientID, err)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		c.hub.HandleMessage(
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
	c.connWriteLock.Lock()
	defer c.connWriteLock.Unlock()
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	err := c.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Errorf("unable to send message: %v : clientID: %v", err, c.clientID)
	}

	return
}

// package scribble will define a `WebsocketHandler` to handle the backend for
// the scribble game
package scribble

import (
	"net/http"

	ws "github.com/bseto/arcade/backend/websocket"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/websocket"
)

type API struct {
	handlers map[string]ws.WebsocketHandler
	registry registry.Registry
}

func GetScribbleAPI(reg registry.Registry) *API {
	return &API{
		handlers: make(map[string]ws.WebsocketHandler),
		registry: reg,
	}
}

func (a *API) HandleMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	err error,
) {
	// stub
	return
}

func (a *API) HandleAuthentication(
	w http.ResponseWriter,
	r *http.Request,
	conn *websocket.Conn,
	send chan []byte,
) (client identifier.Client, err error) {
	// stub
	return
}

func (a *API) SignalClose() {
	// stub
	return
}

func (a *API) Upgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow all origins to connect
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

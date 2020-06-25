package main

import (
	"net/http"

	"github.com/bseto/arcade/backend/game/scribble"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/mux"
)

func main() {
	initializeLogging()
}

func initializeRoutes() {
	r := mux.NewRouter()
	reg := registry.GetRegistryProvider()
	scribbleAPI := scribble.GetScribbleAPI(reg)

	r.PathPrefix("/ws/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wsClient := websocket.GetClient(scribbleAPI)
		wsClient.Upgrade(w, r)
	})
}

func initializeLogging() {
	config := log.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      log.Info,
		ConsoleJSONFormat: false,
		EnableFile:        true,
		FileLevel:         log.Debug,
		FileJSONFormat:    true,
		FileLocation:      "log.log",
	}
	err := log.NewLogger(config, log.InstanceLogrusLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log: %v", err.Error())
	}
	log.Infof("logging initialized")
}

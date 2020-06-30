package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/bseto/arcade/backend/game/scribble"
	"github.com/bseto/arcade/backend/hub"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/mux"
)

var port *int = flag.Int("port", 8080, "defines the port to listen and serve on")

func main() {
	flag.Parse()
	initializeLogging()
	initializeRoutes()
}

func initializeRoutes() {
	r := mux.NewRouter()
	reg := registry.GetRegistryProvider()
	scribbleAPI := scribble.GetScribbleAPI(reg)
	hub := hub.GetHub(reg)
	hub.SetupRoutes(r)

	r.PathPrefix("/ws/{hubID}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wsClient := websocket.GetClient(scribbleAPI)
		wsClient.Upgrade(w, r)
	})

	address := fmt.Sprintf(":%v", *port)
	log.Infof("starting server on: %v", address)
	err := http.ListenAndServe(address, r)
	log.Fatalf("unable to listen and serve: %v", err)
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

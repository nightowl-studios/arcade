package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/hubapi"
	"github.com/bseto/arcade/backend/websocket"
)

// TestTwoConnectionsScenario will ensure that two clients can connect
// This is probably the most hack test written ever
func TestTwoConnectionsScenario(t *testing.T) {
	go initializeRoutes()

	time.Sleep(time.Second * 1)

	websocketClient1, send1, conn1, err := websocket.DialClient("ws://localhost:8081/ws/1")
	if err != nil {
		t.Fatalf("unable to create websocket connection")
	}
	defer websocketClient1.Close()

	_, msg, err := conn1.ReadMessage()
	if err != nil {
		t.Error("unable to read message")
	}

	var gameMsg game.Message
	err = json.Unmarshal(msg, &gameMsg)
	if err != nil {
		t.Errorf("unable to unmarshal game message: %v", err)
	}

	if gameMsg.API != "hub" {
		t.Errorf("expected api: hub, got: %v", gameMsg.API)
	}

	var hubReply hubapi.HubAPIReply
	err = json.Unmarshal(gameMsg.Payload, &hubReply)

	if len(hubReply.ConnectedClients) != 1 {
		t.Errorf("expected 1 connected client, have: %v", len(hubReply.ConnectedClients))
	}

	lobbyDetails := `
	{
		"api":"hub",
			"payload":{
				"requestLobbyDetails": true,
				"changeNameTo": "Gordon"
			}
		}
	`

	send1 <- []byte(lobbyDetails)

	_, msg, err = conn1.ReadMessage()
	if err != nil {
		t.Error("unable to read message")
	}

	err = json.Unmarshal(msg, &gameMsg)
	if err != nil {
		t.Errorf("unable to unmarshal game message: %v", err)
	}

	if gameMsg.API != "hub" {
		t.Errorf("expected api: hub, got: %v", gameMsg.API)
	}

	err = json.Unmarshal(gameMsg.Payload, &hubReply)

	if len(hubReply.ConnectedClients) != 1 {
		t.Errorf("expected 1 connected client, have: %v", len(hubReply.ConnectedClients))
	}
	if hubReply.ConnectedClients[0].GetNickName() != "Gordon" {
		t.Errorf("expected: Gordon, got :%v", hubReply.ConnectedClients[0].GetNickName())
	}

	// Connect the second client now

	websocketClient2, send2, conn2, err := websocket.DialClient("ws://localhost:8081/ws/1")
	if err != nil {
		t.Fatalf("unable to create websocket connection")
	}
	defer websocketClient2.Close()

	_, msg, err = conn2.ReadMessage()
	if err != nil {
		t.Error("unable to read message")
	}
	err = json.Unmarshal(msg, &gameMsg)
	if err != nil {
		t.Errorf("unable to unmarshal game message: %v", err)
	}
	if gameMsg.API != "hub" {
		t.Errorf("expected api: hub, got: %v", gameMsg.API)
	}
	err = json.Unmarshal(gameMsg.Payload, &hubReply)
	if len(hubReply.ConnectedClients) != 2 {
		t.Errorf("expected 2 connected client, have: %v", len(hubReply.ConnectedClients))
	}
	_, msg, err = conn1.ReadMessage()
	if err != nil {
		t.Error("unable to read message")
	}
	err = json.Unmarshal(msg, &gameMsg)
	if err != nil {
		t.Errorf("unable to unmarshal game message: %v", err)
	}
	if gameMsg.API != "hub" {
		t.Errorf("expected api: hub, got: %v", gameMsg.API)
	}
	err = json.Unmarshal(gameMsg.Payload, &hubReply)
	if len(hubReply.ConnectedClients) != 2 {
		t.Errorf("expected 1 connected client, have: %v", len(hubReply.ConnectedClients))
	}

	send2 <- []byte(lobbyDetails)

	_, msg, err = conn2.ReadMessage()
	if err != nil {
		t.Error("unable to read message")
	}
	err = json.Unmarshal(msg, &gameMsg)
	if err != nil {
		t.Errorf("unable to unmarshal game message: %v", err)
	}
	if gameMsg.API != "hub" {
		t.Errorf("expected api: hub, got: %v", gameMsg.API)
	}
	err = json.Unmarshal(gameMsg.Payload, &hubReply)
	if len(hubReply.ConnectedClients) != 2 {
		t.Errorf("expected 2 connected client, have: %v", len(hubReply.ConnectedClients))
	}
	if hubReply.ConnectedClients[0].GetNickName() != "Gordon" {
		t.Errorf("expected: Gordon, got :%v", hubReply.ConnectedClients[0].GetNickName())
	}
	if hubReply.ConnectedClients[1].GetNickName() != "Gordon" {
		t.Errorf("expected: Gordon, got :%v", hubReply.ConnectedClients[1].GetNickName())
	}

	_, msg, err = conn1.ReadMessage()
	if err != nil {
		t.Error("unable to read message")
	}
	err = json.Unmarshal(msg, &gameMsg)
	if err != nil {
		t.Errorf("unable to unmarshal game message: %v", err)
	}
	if gameMsg.API != "hub" {
		t.Errorf("expected api: hub, got: %v", gameMsg.API)
	}
	err = json.Unmarshal(gameMsg.Payload, &hubReply)
	if len(hubReply.ConnectedClients) != 2 {
		t.Errorf("expected 2 connected client, have: %v", len(hubReply.ConnectedClients))
	}
	if hubReply.ConnectedClients[0].GetNickName() != "Gordon" {
		t.Errorf("expected: Gordon, got :%v", hubReply.ConnectedClients[0].GetNickName())
	}
	if hubReply.ConnectedClients[1].GetNickName() != "Gordon" {
		t.Errorf("expected: Gordon, got :%v", hubReply.ConnectedClients[1].GetNickName())
	}

}

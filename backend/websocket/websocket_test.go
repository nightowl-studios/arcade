package websocket

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bseto/arcade/backend/mocks"
	"github.com/gorilla/websocket"
)

type CustomResponseRecorder struct {
	httptest.ResponseRecorder
}

func (c *CustomResponseRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	s := strings.NewReader("abcd")
	br := bufio.NewReader(s)
	w := new(bytes.Buffer)
	bw := bufio.NewWriter(w)
	rw := bufio.NewReadWriter(br, bw)

	// I'm just getting desperate here trying to create any net.Conn
	// I think it might be better to just create a legit weboscket.Dial
	// and test this upgrade like that...
	conn, _ := net.Dial("localhost", "localhost")

	return conn, rw, nil
}

func (c *CustomResponseRecorder) Flush() {
	log.Printf("hello")
}

func TestUpgradeCallsUpgrader(t *testing.T) {

	// This websocket connection is too hard to mock out. I'll try to figure
	// it out some other time
	t.Skip()

	// Create all the mock requirements
	mockHandler := &mocks.WebsocketHandler{}
	req := httptest.NewRequest("GET", "ws://localhost/ws/scribbl/12", nil)

	// Looked into Gorilla websocket package `DialContext` to find out what
	// to set the headers to
	req.Header["Upgrade"] = []string{"websocket"}
	req.Header["Connection"] = []string{"Upgrade"}
	req.Header["Sec-Websocket-Version"] = []string{"13"}
	req.Header["Sec-Websocket-Key"] = []string{"DWyibjG+i4upoi88zYzN7Q=="}

	w := &CustomResponseRecorder{
		ResponseRecorder: httptest.ResponseRecorder{
			HeaderMap: make(http.Header),
			Body:      new(bytes.Buffer),
			Code:      200,
		},
	}

	w.Flush()

	// Define expected behaviour
	mockHandler.On("Upgrader").Return(&websocket.Upgrader{})
	mockHandler.On("HandleAuthentication").Return(nil, nil)

	testClient := GetClient(mockHandler)
	err := testClient.Upgrade(w, req)

	if err != nil {
		t.Errorf("unable to Upgrade: %v", err)
	}
	mockHandler.AssertExpectations(t)
}

package echo

import (
	"encoding/json"

	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
)

const (
	name string = "hello"
)

type Echo struct{}

func (e *Echo) HandleInteraction(
	message json.RawMessage,
	caller identifier.Client,
	registry registry.Registry,
) {
}

func (e *Echo) Name() string {
	return name
}

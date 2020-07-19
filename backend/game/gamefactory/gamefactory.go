package gamefactory

import (
	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble"
	"github.com/bseto/arcade/backend/websocket/registry"
)

// GameFactory should give a new game router out to anyone who requests it
type GameFactory interface {
	GetAvailableGames() []string
	GetGame(string, registry.Registry) game.GameRouter
}

type gameFactory struct {
	// the reason we have a map to functions that provide game routers
	// is because we do not want to store the gamerouter here. We want to be
	// able to grab a new instance and give it to a hub that requests it
	routers map[string]func(reg registry.Registry) game.GameRouter
}

func GetGameFactory() *gameFactory {
	routersMap := CreateGameRoutersMap(
		scribble.GetScribbleRouter,
	)
	return &gameFactory{
		routers: routersMap,
	}
}

func (g *gameFactory) GetAvailableGames() []string {
	var gameList []string
	for key, _ := range g.routers {
		gameList = append(gameList, key)
	}
	return gameList
}

func (g *gameFactory) GetGame(gameName string, reg registry.Registry) game.GameRouter {
	return g.routers[gameName](reg)
}

func CreateGameRoutersMap(routers ...func(registry.Registry) game.GameRouter) map[string]func(registry.Registry) game.GameRouter {
	routerMap := make(map[string]func(registry.Registry) game.GameRouter)
	for _, router := range routers {
		// kinda hacky just to get the router name...
		routerMap[router(nil).RouterName()] = router
	}
	return routerMap
}

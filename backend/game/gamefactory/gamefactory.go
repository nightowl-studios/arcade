package gamefactory

import (
	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble"
)

// GameFactory should give a new game router out to anyone who requests it
type GameFactory interface {
	GetAvailableGames() []string
	GetGame(string) game.GameRouter
}

type gameFactory struct {
	// the reason we have a map to functions that provide game routers
	// is because we do not want to store the gamerouter here. We want to be
	// able to grab a new instance and give it to a hub that requests it
	routers map[string]func() game.GameRouter
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

func (g *gameFactory) GetGame(gameName string) game.GameRouter {
	return g.routers[gameName]()
}

func CreateGameRoutersMap(routers ...func() game.GameRouter) map[string]func() game.GameRouter {
	routerMap := make(map[string]func() game.GameRouter)
	for _, router := range routers {
		routerMap[router().RouterName()] = router
	}
	return routerMap
}

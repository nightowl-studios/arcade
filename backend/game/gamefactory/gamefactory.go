package gamefactory

import (
	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/game/scribble"
)

type GameFactory interface {
	GetAvailableGames() []string
	GetGame(string) game.GameRouter
}

type gameFactory struct {
	routers map[string]game.GameRouter
}

func GetGameFactory() *gameFactory {
	routersMap := CreateGameRoutersMap(
		scribble.GetScribbleRouter(),
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
	return g.routers[gameName]
}

func CreateGameRoutersMap(routers ...game.GameRouter) map[string]game.GameRouter {
	routerMap := make(map[string]game.GameRouter)
	for _, router := range routers {
		routerMap[router.RouterName()] = router
	}
	return routerMap
}

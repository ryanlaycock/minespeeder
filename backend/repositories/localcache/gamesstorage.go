package localcache

import (
	"sync"

	"github.com/ryanlaycock/minespeeder/domain/games"
	"fmt"
)

type LocalCache struct {
	Games sync.Map
}

func NewLocalCache() *LocalCache {
	return &LocalCache{
		Games: sync.Map{},
	}
}

func (l *LocalCache) GetGame(gameId string) (games.Game, error) {
	game, ok := l.Games.Load(gameId)
	if !ok {
		return games.Game{}, fmt.Errorf("game with id %s not found", gameId)
	}
	return game.(games.Game), nil
}

func (l *LocalCache) CreateGame(gameId string) (games.Game, error) {
	game := games.Game{
		Boards: map[string]games.Board{},
	}
	l.Games.Store(gameId, game)
	return game, nil
}

func (l *LocalCache) StartGame(gameId string) (games.Game, error) {
	game, err := l.GetGame(gameId)
	if err != nil {
		return games.Game{}, err
	}
	for _, board := range game.Boards {
		for _, tile := range board.Tiles {
			tile.State = games.Hidden
		}
	}
	l.Games.Store(gameId, game)
	return game, nil
}

func (l *LocalCache) AddBoard(gameId string, boardId string, boardOptions games.BoardOptions) (games.Game, error) {
	game, err := l.GetGame(gameId)
	if err != nil {
		return games.Game{}, err
	}
	board := games.Board{
		Tiles: []games.Tile{},
	}
	for x := 0; x < boardOptions.Width; x++ {
		for y := 0; y < boardOptions.Height; y++ {
			tile := games.Tile{
				State: games.Hidden,
				XPos:  x,
				YPos:  y,
			}
			board.Tiles = append(board.Tiles, tile)
		}
	}
	game.Boards[boardId] = board
	l.Games.Store(gameId, game)
	return game, nil
}

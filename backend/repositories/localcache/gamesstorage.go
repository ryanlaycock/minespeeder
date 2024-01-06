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

func (l *LocalCache) GetBoard(gameId string, boardId string) (games.Board, error) {
	game, err := l.GetGame(gameId)
	if err != nil {
		return games.Board{}, err
	}

	board, ok := game.Boards[boardId]
	if !ok {
		return games.Board{}, fmt.Errorf("board with id %s not found", boardId)
	}
	return board, nil
}

func (l *LocalCache) StoreBoard(gameId string, boardId string, board games.Board) (games.Board, error) {
	game, err := l.GetGame(gameId)
	if err != nil {
		return games.Board{}, err
	}

	game.Boards[boardId] = board
	l.Games.Store(gameId, game)
	return board, nil
}

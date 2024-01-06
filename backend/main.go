package main

import (
	"fmt"

	"github.com/ryanlaycock/minespeeder/api"
	"github.com/ryanlaycock/minespeeder/domain/games"
	"github.com/ryanlaycock/minespeeder/repositories/localcache"
)

func main() {
	gamesStorage := localcache.NewLocalCache()
	gamesManager := games.NewGamesManager(gamesStorage)
	m := api.NewMineSpeederServer(*gamesManager)

	// Temp setup to return an example board
	gamesManager.CreateGame("game1")
	gamesManager.CreateBoard("game1", "board1", games.BoardOptions{
		Width:  8, // TODO Fix these being the wrong way around
		Height: 6,
		NumberOfBombs: 3,
	})

	fmt.Println("Starting server")
	m.ListenAndServe()
	fmt.Println("Server started")
}

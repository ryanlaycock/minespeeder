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

	fmt.Println("Starting server")
	m.ListenAndServe()
	fmt.Println("Server closed")
}

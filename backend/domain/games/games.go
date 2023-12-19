package games

type GamesStorage interface {
	GetGame(gameId string) (Game, error)
	CreateGame(gameId string) (Game, error)
	AddBoard(gameId string, boardId string, boardOptions BoardOptions) (Game, error)
	StartGame(gameId string) (Game, error)
}

type GamesManager struct {
	GamesStorage GamesStorage
}

func (gm *GamesManager) GetGame(gameId string) (Game, error) {
	return gm.GamesStorage.GetGame(gameId)
}

func (gm *GamesManager) CreateGame(gameId string) (Game, error) {
	return gm.GamesStorage.CreateGame(gameId)
}

func (gm *GamesManager) AddBoard(gameId string, boardId string, boardOptions BoardOptions) (Game, error) {
	return gm.GamesStorage.AddBoard(gameId, boardId, boardOptions)
}

func (gm *GamesManager) StartGame(gameId string) (Game, error) {
	return gm.GamesStorage.StartGame(gameId)
}

func NewGamesManager(
	storage GamesStorage,
) *GamesManager {
	return &GamesManager{
		GamesStorage: storage,
	}
}

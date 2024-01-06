package games

import (
	"fmt"
	"math/rand"
	"strconv"
)

type GamesStorage interface {
	GetGame(gameId string) (Game, error)
	CreateGame(gameId string) (Game, error)

	GetBoard(gameId string, boardId string) (Board, error)
	StoreBoard(gameId string, boardId string, board Board) (Board, error)
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

func (gm *GamesManager) GetBoard(gameId string, boardId string) (Board, error) {
	return gm.GamesStorage.GetBoard(gameId, boardId)
}

func (gm *GamesManager) CreateBoard(gameId string, boardId string, boardOptions BoardOptions) (Board, error) {
	_, err := gm.GamesStorage.GetGame(gameId)
	if err != nil {
		return Board{}, err
	}

	board := createBoard(boardOptions)
	board.Height = boardOptions.Height
	board.Width = boardOptions.Width
	board.NumberOfBombs = boardOptions.NumberOfBombs

	board, err = gm.GamesStorage.StoreBoard(gameId, boardId, board)
	if err != nil {
		return Board{}, err
	}
	return board, nil
}

func (gm *GamesManager) ApplyAction(gameId string, boardId string, action Action) (Board, error) {
	board, err := gm.GamesStorage.GetBoard(gameId, boardId)
	if err != nil {
		return Board{}, err
	}
	fmt.Println(board)

	err = board.ApplyAction(action)
	if err != nil {
		return Board{}, err
	}

	board, err = gm.GamesStorage.StoreBoard(gameId, boardId, board)
	if err != nil {
		return Board{}, err
	}
	fmt.Println(board)
	return board, nil
}

func (b *Board) ApplyAction(action Action) error {
	tile, err := b.GetTile(action.XPos, action.YPos)
	if err != nil {
		return err
	}

	switch action.ActionType {
	case RevealAction:
		tile.CurrentState = tile.Value
		if tile.Value == Empty {
			b.RevealEmptyNeighbourTile(action.XPos, action.YPos)
		}
	case FlagAction:
		tile.CurrentState = Flag
	}

	b.SetTile(action.XPos, action.YPos, *tile)
	return nil
}

func (b *Board) RevealEmptyNeighbourTile(xPos int, yPos int) {
	tile, err := b.GetTile(xPos, yPos)
	if err != nil {
		return
	}

	if tile.Value != Empty {
		return
	}

	tile.CurrentState = Empty
	b.SetTile(xPos, yPos, *tile)

	for x := xPos - 1; x <= xPos+1; x++ {
		for y := yPos - 1; y <= yPos+1; y++ {
			if x == xPos && y == yPos {
				continue
			}
			if x < 0 || x >= 10 || y < 0 || y >= 10 {
				continue
			}
			b.RevealEmptyNeighbourTile(x, y)
		}
	}
}

func (b *Board) GetTile(xPos int, yPos int) (*Tile, error) {
	for _, tile := range b.Tiles {
		if tile.XPos == xPos && tile.YPos == yPos {
			return &tile, nil
		}
	}
	return &Tile{}, fmt.Errorf("tile not found")
}

func (b *Board) SetTile(xPos int, yPos int, tile Tile) {
	for i, t := range b.Tiles {
		if t.XPos == xPos && t.YPos == yPos {
			b.Tiles[i] = tile
			return
		}
	}
}

func createBoard(boardOptions BoardOptions) Board {
	b := make([][]Tile, boardOptions.Width)
	for i := 0; i < boardOptions.Width; i++ {
		b[i] = make([]Tile, boardOptions.Height)
	}

	// Randomly set bombs
	assignedBombs := 0
	for {
		x := rand.Intn(boardOptions.Width)
		y := rand.Intn(boardOptions.Height)
		if b[x][y].Value != Bomb {
			b[x][y].Value = Bomb
			assignedBombs++
		}
		if assignedBombs == boardOptions.NumberOfBombs {
			break
		}
	}

	// Set numbers
	for x := 0; x < boardOptions.Width; x++ {
		for y := 0; y < boardOptions.Height; y++ {
			// Set location and current state properties
			b[x][y].CurrentState = Hidden
			b[x][y].XPos = x
			b[x][y].YPos = y

			if b[x][y].Value == Bomb {
				continue
			}

			// Check all adjacent tiles for bombs and set value accordingly
			bombCount := 0
			for x2 := x - 1; x2 <= x+1; x2++ {
				for y2 := y - 1; y2 <= y+1; y2++ {
					if x2 == x && y2 == y {
						continue
					}
					if x2 < 0 || x2 >= boardOptions.Width || y2 < 0 || y2 >= boardOptions.Height {
						continue
					}
					if b[x2][y2].Value == Bomb {
						bombCount++
					}
				}
			}
			b[x][y].Value = TileState(strconv.Itoa(bombCount))
		}
	}

	board := Board{
		Tiles: []Tile{},
	}
	for _, row := range b {
		for _, tile := range row {
			board.Tiles = append(board.Tiles, tile)
		}
	}

	return board
}

func NewGamesManager(
	storage GamesStorage,
) *GamesManager {
	return &GamesManager{
		GamesStorage: storage,
	}
}

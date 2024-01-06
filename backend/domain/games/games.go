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

	board := Board{
		Tiles:         []Tile{},
		Height:        boardOptions.Height,
		Width:         boardOptions.Width,
		NumberOfBombs: boardOptions.NumberOfBombs,
		RevealedTiles: 0,
	}

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

	if board.RevealedTiles == 0 {
		// First move of the game, populate the board
		board.PopulateBoard(action.XPos, action.YPos)
	}

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
		switch tile.CurrentState {
		case Hidden:
			tile.CurrentState = Flag
		case Flag:
			tile.CurrentState = Hidden
		}
	}

	b.SetTile(action.XPos, action.YPos, *tile)
	return nil
}

func (b *Board) RevealEmptyNeighbourTile(xPos int, yPos int) {
	tile, err := b.GetTile(xPos, yPos)
	if err != nil {
		return
	}

	if tile.CurrentState != Hidden {
		// Tile is already displayed so stop searching here
		return
	}

	if tile.Value != Empty && tile.CurrentState == Hidden {
		// Tile is not empty but is a neighbour to an empty tile so display it
		// but don't search any further
		tile.CurrentState = tile.Value
		b.SetTile(xPos, yPos, *tile)
		return
	}

	// Tile is empty so display it and search for any empty neighbours

	tile.CurrentState = Empty
	b.SetTile(xPos, yPos, *tile)

	for x := xPos - 1; x <= xPos+1; x++ {
		for y := yPos - 1; y <= yPos+1; y++ {
			if x == xPos && y == yPos {
				continue
			}
			if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
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
			if tile.CurrentState != Hidden {
				b.RevealedTiles++ // TODO This isn't right, it gets called from too many places
				fmt.Printf("Revealed: %v. Remaining: %v\n", b.RevealedTiles, (b.Width*b.Height)-b.RevealedTiles)
			}
			return
		}
	}
}

func (b *Board) PopulateBoard(startingXPos int, startingYPos int) error {
	tiles := make([][]Tile, b.Width)
	for i := 0; i < b.Width; i++ {
		tiles[i] = make([]Tile, b.Height)
	}

	// Randomly set bombs, avoiding the starting and adjacent tiles
	assignedBombs := 0
	for {
		x := rand.Intn(b.Width)
		y := rand.Intn(b.Height)

		if (x >= startingXPos-1 && x <= startingXPos+1) && (y >= startingYPos-1 && y <= startingYPos+1) {
			// x,y is equal to or adjacent to starting tile so try again
			continue
		}
		if tiles[x][y].Value != Bomb {
			tiles[x][y].Value = Bomb
			assignedBombs++
		}
		if assignedBombs == b.NumberOfBombs {
			break
		}
	}

	// Set numbers
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			// Set location and current state properties
			tiles[x][y].CurrentState = Hidden
			tiles[x][y].XPos = x
			tiles[x][y].YPos = y

			if tiles[x][y].Value == Bomb {
				continue
			}

			// Check all adjacent tiles for bombs and set value accordingly
			bombCount := 0
			for x2 := x - 1; x2 <= x+1; x2++ {
				for y2 := y - 1; y2 <= y+1; y2++ {
					if x2 == x && y2 == y {
						continue
					}
					if x2 < 0 || x2 >= b.Width || y2 < 0 || y2 >= b.Height {
						continue
					}
					if tiles[x2][y2].Value == Bomb {
						bombCount++
					}
				}
			}
			if bombCount == 0 {
				tiles[x][y].Value = Empty
				continue
			}
			tiles[x][y].Value = TileState(strconv.Itoa(bombCount))
		}
	}

	for _, row := range tiles {
		b.Tiles = append(b.Tiles, row...)
	}

	return nil
}

func NewGamesManager(
	storage GamesStorage,
) *GamesManager {
	return &GamesManager{
		GamesStorage: storage,
	}
}

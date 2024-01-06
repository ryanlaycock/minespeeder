package games

// game represents a single game of minespeeder
// contains a map of boards, identified by playerId
type Game struct {
	Boards map[string]Board
}

type Board struct {
	Tiles []Tile
	Height int
	Width int
	NumberOfBombs int
}

type BoardOptions struct {
	Width  int
	Height int
	NumberOfBombs int
}

type Tile struct {
	Value TileState
	CurrentState TileState
	XPos  int
	YPos  int
}

type TileState string

const (
	Hidden TileState = "hidden"
	Bomb   TileState = "bomb"
	Empty  TileState = "empty"
	Flag   TileState = "flag"
	N1     TileState = "1"
	N2     TileState = "2"
	N3     TileState = "3"
	N4     TileState = "4"
	N5     TileState = "5"
	N6     TileState = "6"
	N7     TileState = "7"
	N8     TileState = "8"
)

type Action struct {
	ActionType ActionType
	XPos       int
	YPos       int
}

type ActionType string

const (
	RevealAction ActionType = "reveal"
	FlagAction   ActionType = "flag"
)

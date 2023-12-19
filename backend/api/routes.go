package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanlaycock/minespeeder/domain/games"
)

// GetV1GamesGameIdBoardsBoardId implements ServerInterface.
func (m *MineSpeederServer) GetV1GamesGameIdBoardsBoardId(
	w http.ResponseWriter, 
	r *http.Request, 
	gameId string, 
	boardId string,
) {
	game, err := m.gamesManager.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	board, exists := game.Boards[boardId]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := DomainBoardToAPIBoard(board)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)	
}

// DomainBoardToAPIBoard converts a games.Board to a Board
func DomainBoardToAPIBoard(board games.Board) Board {
	apiBoard := Board{
		Tiles: []Tile{},
	}
	
	for _, tile := range board.Tiles {
		apiBoard.Tiles = append(apiBoard.Tiles, Tile{
			State: TileState(tile.State),
			XPos: tile.XPos,
			YPos: tile.YPos,
		})
	}
	return apiBoard
}
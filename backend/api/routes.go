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

func (m *MineSpeederServer) PostV1GamesGameIdBoardsBoardIdActions(
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

	_, exists := game.Boards[boardId]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var action Action
	err = json.NewDecoder(r.Body).Decode(&action)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	board, err := m.gamesManager.ApplyAction(gameId, boardId, APIActionToDomainAction(action))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := DomainBoardToAPIBoard(board)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// DomainBoardToAPIBoard converts a games.Board to a Board
func DomainBoardToAPIBoard(board games.Board) Board {
	apiBoard := Board{
		Tiles:                  []Tile{},
		Height:                 board.Height,
		Width:                  board.Width,
		NumberOfBombs:          board.NumberOfBombs,
		NumberOfRemainingBombs: board.RemainingBombs,
		NumberOfRemainingTiles: board.RemainingTiles,
		NumberOfTiles:          board.NumberOfTiles,
		State:                  BoardState(board.State),
	}

	for _, tile := range board.Tiles {
		apiBoard.Tiles = append(apiBoard.Tiles, Tile{
			State: TileState(tile.CurrentState),
			XPos:  tile.XPos,
			YPos:  tile.YPos,
		})
	}
	return apiBoard
}

func APIActionToDomainAction(action Action) games.Action {
	return games.Action{
		ActionType: games.ActionType(action.Type),
		XPos:       action.XPos,
		YPos:       action.YPos,
	}
}

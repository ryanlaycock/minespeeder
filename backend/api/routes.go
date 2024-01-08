package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryanlaycock/minespeeder/domain/games"
)

func (m *MineSpeederServer) GetV1GamesGameId(
	w http.ResponseWriter,
	r *http.Request,
	gameId string,
) {
	game, err := m.gamesManager.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Game{
		Id:     game.Id,
		State:  GameState(game.State),
		Boards: []Board{},
	}

	for _, board := range game.Boards {
		resp.Boards = append(resp.Boards, DomainBoardToAPIBoard(board))
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (m *MineSpeederServer) PostV1Games(
	w http.ResponseWriter,
	r *http.Request,
) {
	var gameReq CreateGameRequest
	err := json.NewDecoder(r.Body).Decode(&gameReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = m.gamesManager.GetGame(gameReq.Id)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	// TODO Add check that error is JUST not found, if not return 500

	m.gamesManager.CreateGame(gameReq.Id)
	for i := 0; i < gameReq.NumberOfBoards; i++ {
		m.gamesManager.CreateBoard(gameReq.Id, uuid.New().String(), games.BoardOptions{
			Width:         gameReq.BoardOptions.Width,
			Height:        gameReq.BoardOptions.Height,
			NumberOfBombs: gameReq.BoardOptions.NumberOfBombs,
		})
	}

	game, err := m.gamesManager.GetGame(gameReq.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBoards := []Board{}
	for _, board := range game.Boards {
		respBoards = append(respBoards, DomainBoardToAPIBoard(board))
	}

	resp := Game{
		Id:     game.Id,
		State:  GameState(Created),
		Boards: respBoards,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

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
		Id:                     board.Id,
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

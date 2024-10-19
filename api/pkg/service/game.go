package service

import (
	"errors"
	"movie_list/api/pkg/models"
	"sync"

	"github.com/google/uuid"
)

type GameInterface interface {
	AddPlayer(req models.AddPlayerRequest) models.AddPlayerResponse
	GetPlayers() (models.GetPlayersResponse, error)
	DeletePlayer(req models.DeletePlayerRequest) models.DeletePlayerResponse
}

type GameService struct {
	mu      sync.RWMutex
	Players map[uuid.UUID]models.Player
}

func NewGameService() *GameService {
	return &GameService{
		Players: make(map[uuid.UUID]models.Player),
	}
}

func (gs *GameService) AddPlayer(req models.AddPlayerRequest) models.AddPlayerResponse {
	if req.Name == "" {
		return models.AddPlayerResponse{
			ID:      uuid.Nil,
			Success: false,
			Message: "Name is required",
		}
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	playerID := uuid.New()
	player := models.Player{
		ID:   playerID,
		Name: req.Name,
	}

	gs.Players[playerID] = player

	return models.AddPlayerResponse{
		ID:      playerID,
		Success: true,
		Message: "Player added successfully",
	}
}

func (gs *GameService) GetPlayers() (models.GetPlayersResponse, error) {
	gs.mu.RLock()
	defer gs.mu.RUnlock()

	if gs.Players == nil {
		return models.GetPlayersResponse{
			Players: nil,
			Success: false,
			Message: "Players data is unavailable",
		}, errors.New("players data is unavailable")
	}

	players := make([]models.Player, 0, len(gs.Players))
	for _, player := range gs.Players {
		players = append(players, player)
	}

	return models.GetPlayersResponse{
		Players: players,
		Success: true,
		Message: "Players retrieved successfully",
	}, nil
}

func (gs *GameService) DeletePlayer(req models.DeletePlayerRequest) models.DeletePlayerResponse {
	if req.ID == uuid.Nil {
		return models.DeletePlayerResponse{
			Success: false,
			Message: "Invalid player ID",
		}
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if _, ok := gs.Players[req.ID]; !ok {
		return models.DeletePlayerResponse{
			Success: false,
			Message: "Player not found",
		}
	}

	delete(gs.Players, req.ID)

	return models.DeletePlayerResponse{
		Success: true,
		Message: "Player deleted successfully",
	}
}

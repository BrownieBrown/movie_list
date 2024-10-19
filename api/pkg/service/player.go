package service

import (
	"errors"
	"log"
	"movie_list/api/pkg/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type PlayerServiceInterface interface {
	AddPlayer(req models.AddPlayerRequest) models.AddPlayerResponse
	GetPlayers() (models.GetPlayersResponse, error)
	DeletePlayer(req models.DeletePlayerRequest) models.DeletePlayerResponse
}

type PlayerService struct {
	mu      sync.RWMutex
	Players map[uuid.UUID]models.Player
}

func NewPlayerService() *PlayerService {
	return &PlayerService{
		Players: make(map[uuid.UUID]models.Player),
	}
}

func (s *PlayerService) AddPlayer(req models.AddPlayerRequest) models.AddPlayerResponse {
	if req.Name == "" {
		log.Printf("[%s] AddPlayer: Name is required\n", time.Now().Format(time.RFC3339))
		return models.AddPlayerResponse{
			ID:      uuid.Nil,
			Success: false,
			Message: "Name is required",
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	playerID := uuid.New()
	player := models.Player{
		ID:   playerID,
		Name: req.Name,
	}

	s.Players[playerID] = player

	log.Printf("[%s] AddPlayer: Added player with ID %s\n", time.Now().Format(time.RFC3339), playerID)
	return models.AddPlayerResponse{
		ID:      playerID,
		Success: true,
		Message: "Player added successfully",
	}
}

func (s *PlayerService) GetPlayers() (models.GetPlayersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.Players == nil {
		log.Printf("[%s] GetPlayers: Players map is nil\n", time.Now().Format(time.RFC3339))
		return models.GetPlayersResponse{
			Players: nil,
			Success: false,
			Message: "Players data is unavailable",
		}, errors.New("players data is unavailable")
	}

	players := make([]models.Player, 0, len(s.Players))
	for _, player := range s.Players {
		players = append(players, player)
	}

	log.Printf("[%s] GetPlayers: Retrieved %d players\n", time.Now().Format(time.RFC3339), len(players))
	return models.GetPlayersResponse{
		Players: players,
		Success: true,
		Message: "Players retrieved successfully",
	}, nil
}

func (s *PlayerService) DeletePlayer(req models.DeletePlayerRequest) models.DeletePlayerResponse {
	if req.ID == uuid.Nil {
		log.Printf("[%s] DeletePlayer: Invalid UUID provided\n", time.Now().Format(time.RFC3339))
		return models.DeletePlayerResponse{
			Success: false,
			Message: "Invalid player ID",
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Players[req.ID]; !ok {
		log.Printf("[%s] DeletePlayer: Player with ID %s not found\n", time.Now().Format(time.RFC3339), req.ID)
		return models.DeletePlayerResponse{
			Success: false,
			Message: "Player not found",
		}
	}

	delete(s.Players, req.ID)

	log.Printf("[%s] DeletePlayer: Deleted player with ID %s\n", time.Now().Format(time.RFC3339), req.ID)
	return models.DeletePlayerResponse{
		Success: true,
		Message: "Player deleted successfully",
	}
}

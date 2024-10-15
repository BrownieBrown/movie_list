package server

import (
	"context"
	"log"
	"movie_list/pkg/protobuf"
	"sync"

	"github.com/google/uuid"
)

type Server struct {
	protobuf.UnimplementedPlayerServiceServer
	mu      sync.Mutex
	Players map[uuid.UUID]string
}

func (s *Server) AddPlayer(ctx context.Context, req *protobuf.AddPlayerRequest) (*protobuf.AddPlayerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New()
	s.Players[id] = req.GetName()
	return &protobuf.AddPlayerResponse{Id: id.String()}, nil
}

func (s *Server) RemovePlayer(ctx context.Context, req *protobuf.RemovePlayerRequest) (*protobuf.RemovePlayerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Println("error parsing id:", err)
		return nil, err
	}

	if _, exists := s.Players[id]; exists {
		delete(s.Players, id)
		log.Printf("Removed player: ID=%s", id)
		return &protobuf.RemovePlayerResponse{Success: true}, nil
	}

	log.Printf("Failed to remove player: ID=%s not found", id)
	return &protobuf.RemovePlayerResponse{Success: false}, nil
}

package models

import (
	"github.com/google/uuid"
)

type Player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type AddPlayerRequest struct {
	Name string `json:"name"`
}

type AddPlayerResponse struct {
	ID      uuid.UUID `json:"id"`
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
}

type DeletePlayerRequest struct {
	ID uuid.UUID `json:"id"`
}

type DeletePlayerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type GetPlayersResponse struct {
	Players []Player `json:"players"`
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
}

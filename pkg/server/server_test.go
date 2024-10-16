package server

import (
	"context"
	"movie_list/pkg/protobuf"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestAddPlayer(t *testing.T) {
	s := &Server{Players: make(map[uuid.UUID]string)}

	req := &protobuf.AddPlayerRequest{Name: "John Doe"}
	resp, err := s.AddPlayer(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetId())

	id, err := uuid.Parse(resp.GetId())
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", s.Players[id])
}

func TestListPlayers(t *testing.T) {
	s := &Server{Players: make(map[uuid.UUID]string)}

	req := &protobuf.ListPlayersRequest{}
	resp, err := s.ListPlayers(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Empty(t, resp.GetPlayers())

	s.Players[uuid.New()] = "John Doe"
	s.Players[uuid.New()] = "Jane Doe"

	resp, err = s.ListPlayers(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.GetPlayers(), 2)
	assert.Contains(t, resp.GetPlayers(), "John Doe")
	assert.Contains(t, resp.GetPlayers(), "Jane Doe")
}

func TestRemovePlayer(t *testing.T) {
	s := &Server{Players: make(map[uuid.UUID]string)}

	id := uuid.New()
	s.Players[id] = "John Doe"

	req := &protobuf.RemovePlayerRequest{Id: id.String()}
	resp, err := s.RemovePlayer(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.GetSuccess())
	assert.NotContains(t, s.Players, id)
}

func TestRemovePlayer_NotFound(t *testing.T) {
	s := &Server{Players: make(map[uuid.UUID]string)}

	req := &protobuf.RemovePlayerRequest{Id: uuid.New().String()}
	resp, err := s.RemovePlayer(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.False(t, resp.GetSuccess())
}

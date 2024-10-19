package controller

import (
	"bytes"
	"encoding/json"
	"movie_list/api/pkg/models"
	"movie_list/api/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()
	playerController := NewPlayerController(service.NewPlayerService())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	player := v1.Group("/player")
	player.Get("/", playerController.GetPlayers)
	player.Post("/", playerController.AddPlayer)
	player.Delete("/", playerController.DeletePlayer)

	return app
}

func TestAddPlayer(t *testing.T) {
	app := setupApp()

	reqBody, _ := json.Marshal(models.AddPlayerRequest{Name: "test"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/player", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var res models.AddPlayerResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal("Failed to decode response:", err)
		return
	}
	assert.NotEqual(t, res.ID, "")
	assert.True(t, res.Success)
}

func TestGetPlayers(t *testing.T) {
	app := setupApp()

	// Add a player first
	reqBody, _ := json.Marshal(models.AddPlayerRequest{Name: "test"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/player", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
		return
	}

	// Now get the players
	req = httptest.NewRequest(http.MethodGet, "/api/v1/player", nil)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var res models.GetPlayersResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
		return
	}
	assert.True(t, res.Success)
	assert.Len(t, res.Players, 1)
	assert.Equal(t, "test", res.Players[0].Name)
}

func TestDeletePlayer(t *testing.T) {
	app := setupApp()

	// Add a player first
	reqBody, _ := json.Marshal(models.AddPlayerRequest{Name: "test"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/player", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	var addRes models.AddPlayerResponse
	err := json.NewDecoder(resp.Body).Decode(&addRes)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
		return
	}

	// Now delete the player
	delReqBody, _ := json.Marshal(models.DeletePlayerRequest{ID: addRes.ID})
	req = httptest.NewRequest(http.MethodDelete, "/api/v1/player", bytes.NewBuffer(delReqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err = app.Test(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var delRes models.DeletePlayerResponse
	err = json.NewDecoder(resp.Body).Decode(&delRes)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
		return
	}
	assert.True(t, delRes.Success)
}

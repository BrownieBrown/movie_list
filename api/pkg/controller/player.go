package controller

import (
	"movie_list/api/pkg/models"
	"movie_list/api/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type PlayerControllerInterface interface {
	AddPlayer(c *fiber.Ctx) error
	GetPlayers(c *fiber.Ctx) error
	DeletePlayer(c *fiber.Ctx) error
}

type PlayerController struct {
	Service *service.PlayerService
}

func NewPlayerController(service *service.PlayerService) *PlayerController {
	return &PlayerController{
		Service: service,
	}
}

func (pc *PlayerController) AddPlayer(c *fiber.Ctx) error {
	var req models.AddPlayerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := pc.Service.AddPlayer(req)
	return c.JSON(res)
}

func (pc *PlayerController) GetPlayers(c *fiber.Ctx) error {
	res, err := pc.Service.GetPlayers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(res)
}

func (pc *PlayerController) DeletePlayer(c *fiber.Ctx) error {
	var req models.DeletePlayerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := pc.Service.DeletePlayer(req)
	return c.JSON(res)
}

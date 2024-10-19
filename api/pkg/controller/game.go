package controller

import (
	"movie_list/api/pkg/models"
	"movie_list/api/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type GameControllerInterface interface {
	AddPlayer(c *fiber.Ctx) error
	GetPlayers(c *fiber.Ctx) error
	DeletePlayer(c *fiber.Ctx) error
}

type GameController struct {
	Service *service.GameService
}

func NewGameController(service *service.GameService) *GameController {
	return &GameController{
		Service: service,
	}
}

func (gc *GameController) AddPlayer(c *fiber.Ctx) error {
	var req models.AddPlayerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := gc.Service.AddPlayer(req)
	return c.JSON(res)
}

func (gc *GameController) GetPlayers(c *fiber.Ctx) error {
	res, err := gc.Service.GetPlayers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(res)
}

func (gc *GameController) DeletePlayer(c *fiber.Ctx) error {
	var req models.DeletePlayerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := gc.Service.DeletePlayer(req)
	return c.JSON(res)
}

package main

import (
	"log"
	"movie_list/api/pkg/controller"
	"movie_list/api/pkg/service"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Update with your frontend's origin
		AllowMethods: "GET,POST,DELETE",
	}))

	playerController := controller.NewPlayerController(service.NewPlayerService())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	player := v1.Group("/player")
	player.Get("/", playerController.GetPlayers)
	player.Post("/", playerController.AddPlayer)
	player.Delete("/", playerController.DeletePlayer)

	port := ":3000"
	err := app.Listen(port)
	if err != nil {
		log.Fatal("Failed to start app:", err)
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kekubhai/zyntic/handlers"
	"github.com/kekubhai/zyntic/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Zyntic",
		ServerHeader: "Fiber",
	})
	eventRepository = repositories.NewEventRepository(nil)
	server := app.Group("/api")
	handlers.NewEventHandlers(server.Group("/event"), eventRepository)
	app.Listen(":3000")
}

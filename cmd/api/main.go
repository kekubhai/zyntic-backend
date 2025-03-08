package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kekubhai/zyntic/config"
	"github.com/kekubhai/zyntic/db"
	"github.com/kekubhai/zyntic/handlers"
	"github.com/kekubhai/zyntic/repositories"
)

func main() {
	envConfig:=config.NewEnvConfig()
	db:=db.Init(envConfig  		,db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "Zyntic",
		ServerHeader: "Fiber",
	})
	eventRepository := repositories.NewEventRepository(db)
	server := app.Group("/api")
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	app.Listen(fmt.Sprintf(":" +envConfig.ServerPort))
}

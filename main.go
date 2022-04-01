package main

import (
	_ "thailephan/flash-card-api/repository"
	"thailephan/flash-card-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	routes.InitRoutes(app)

	app.Listen("127.0.0.1:8080")
}
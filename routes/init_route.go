package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	initFlashcardRoutes(app)
}
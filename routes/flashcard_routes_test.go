package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestGetAllFlashCard(t *testing.T) {
	app := fiber.New()
	app.Get("/flashcards", getAllFlashCard)	
	
	req := httptest.NewRequest("GET", "/flashcards", nil)

	resp, _ := app.Test(req, 1)

	require.Equal(t, resp.StatusCode, 200)
	// tests := []struct {
	// 	description string	
	// }{
	// 	// TODO: Add test cases.
	// }

	// for _, tt := range tests {
	// }
}

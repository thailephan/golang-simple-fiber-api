package utils

import (
	"github.com/gofiber/fiber/v2"
)

func JSON[R Response](ctx *fiber.Ctx, response R) error {
	return ctx.JSON(response)
}

type Response interface {
	SuccessResponse | ErrorResponse
}

type Query struct {
	Offset int64 `json:"offset,omitempty"`
	Limit int64 `json:"limit,omitempty"`
	Filter string `json:"filter,omitempty"`
	Sort string `json:"sort,omitempty"`
}

type Data struct {
	Items interface{} `json:"items"`
	Count int `json:"Count,omitempty"`
	Totals int `json:"totals,omitempty"`
	Query Query `json:"query,omitempty"`
}

type Error struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

type SuccessResponse struct {
	Data Data `json:"data"`
	StatusCode int  `json:"statusCode"`
}

type ErrorResponse struct {
	StatusCode int  `json:"statusCode"`
	Error Error `json:"error"`
}
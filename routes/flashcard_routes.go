package routes

import (
	"thailephan/flash-card-api/models"
	"thailephan/flash-card-api/pkg/utils"
	"thailephan/flash-card-api/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
)

var usecase *usecases.FlashcardUsecase

func initFlashcardRoutes(app *fiber.App) {
	flashcardGroup := app.Group("/flashcards")

	usecase = usecases.NewUsecase()

	flashcardGroup.Get("/", getAllFlashCard)
	flashcardGroup.Get("/:id", getFlashCardById)
	flashcardGroup.Post("/", createFlashCard)
	flashcardGroup.Put("/:id", updateFlashCard)
	flashcardGroup.Delete("/:id", deleteFlashCard)
}

func getAllFlashCard(c *fiber.Ctx) error {
	var limit, offset int64
	var err error

	if limit, err = utils.ParseInt(c.Query("limit", "0"), 64); err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-B",
			},
		}) 
	}
	
	if offset, err = utils.ParseInt(c.Query("offset", "0"), 64); err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-C",
			},
		}) 
	}

	flashcards, err := usecase.GetAll(struct{}{}, utils.Options{
		Limit: limit,
		Offset: offset,
	})

	if err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-D",
			},
		}) 
	}
	return utils.JSON(c, utils.SuccessResponse{
		Data: utils.Data{
			Query: utils.Query{
				Offset: offset,
				Limit: limit,
			},
			Items: flashcards,
		},
		StatusCode: fiber.StatusOK,
	})
}

func getFlashCardById(c *fiber.Ctx) error {
	id := c.Params("id")

	flashcard, err := usecase.GetByID(id)
	if err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-C",
			},
		}) 
	}

	return utils.JSON(c, utils.SuccessResponse{
		Data: utils.Data{
			Items: flashcard,
		},
		StatusCode: fiber.StatusOK,
	})
}

func createFlashCard(c *fiber.Ctx) error {
	var flashcard *models.FlashCard

	if err := c.BodyParser(&flashcard); err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-C",
			},
		}) 
	}

	flashcard.ID = utils.NewObjectID()
	flashcard.CreatedAt = time.Now()
	flashcard.UpdatedAt = time.Now()

	err := usecase.Store(flashcard)

	if (err != nil) {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-C",
			},
		}) 
	}
	return utils.JSON(c, utils.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data: utils.Data{
			Items: flashcard,
		},
	})
}

func updateFlashCard(c *fiber.Ctx) error {
	var updated	*models.FlashCard
	id := c.Params("id")

	idPrimitive, err := utils.ObjectIDFromHex(id)
	if err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-A",
			},
		}) 
	}

	if err := c.BodyParser(&updated); err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-B",
			},
		}) 
	}

	updated.UpdatedAt = time.Now()
	out, err :=utils.BsonMapFromEntity(updated)
	filter := utils.Map{"_id": idPrimitive}
	modifiedCount, err := usecase.Update(filter, utils.Map{"$set": out})
	
	if err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-D",
			},
		}) 
	}

	if modifiedCount ==  0 {
		return utils.JSON(c, utils.SuccessResponse{
			Data: utils.Data{
				Items: fiber.Map{},
			},
			StatusCode: fiber.StatusOK,
		})
	}

	flashcard, err := usecase.GetByID(id)
	if err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-C",
			},
		}) 
	}

	return utils.JSON(c, utils.SuccessResponse{
		Data: utils.Data{
			Items: flashcard,
		},
		StatusCode: fiber.StatusOK,
	})
}

func deleteFlashCard(c *fiber.Ctx) error {
	var primitiveId interface{}
	var err error

	id := c.Params("id")

	if primitiveId, err = utils.ObjectIDFromHex(id); err != nil {
		return c.JSON(fiber.Map{
			"results": nil,
			"statusCode": fiber.StatusBadRequest,
			"error": err.Error(),
		})
	}
	// Interface for implement bson
	count, err := usecase.Delete(utils.Map{"_id": primitiveId})

	if err != nil {
		return utils.JSON(c, utils.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Error: utils.Error{
				Message: err.Error(),
				Code: "A-C",
			},
		}) 
	}

	return utils.JSON(c, utils.SuccessResponse{
		Data: utils.Data{
			Items: count,
		},
		StatusCode: fiber.StatusOK,
	})
}
package usecases

import (
	"thailephan/flash-card-api/models"
	"thailephan/flash-card-api/pkg/utils"
	"thailephan/flash-card-api/repository"
)

type FlashcardUsecase struct {
	Repository *repository.FlashCardRepository
}

func NewUsecase() (usecase *FlashcardUsecase){
	return &FlashcardUsecase{
		Repository: repository.NewFlashCard(),
	}
} 

func (usecase *FlashcardUsecase)GetAll(filter interface{}, opts utils.Options) (flashcards interface{}, err error){
	return usecase.Repository.GetAll(filter, opts)
}

func (usecase *FlashcardUsecase) GetByID(id string) (interface{}, error) {
	return usecase.Repository.GetByID(id)
}

func (usecase *FlashcardUsecase)  Store(model *models.FlashCard) error {
	return usecase.Repository.Store(model)
}

func (usecase *FlashcardUsecase) Update(filter interface{}, update interface{}) (int64, error) {
	return usecase.Repository.Update(filter, update)
}

func (usecase *FlashcardUsecase) Delete(filter interface{}) (int64, error) {
	result, err := usecase.Repository.Delete(filter)
	return result, err
}
package repository

import (
	"math"
	"thailephan/flash-card-api/models"
	"thailephan/flash-card-api/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FlashCardRepository struct {
	Collection *mongo.Collection
}

func NewFlashCard() *FlashCardRepository {
	return &FlashCardRepository{
		Collection: FlashCardsCollection,
	}
}

func (r *FlashCardRepository) GetAll(filter interface{}, opts utils.Options) (interface{}, error) {
	var results = make([]bson.M, 0)

	o := options.Find().SetLimit(int64(math.Min(float64(opts.Limit), 10.0))).SetSkip(opts.Offset)
	cursor, err := r.Collection.Find(Ctx, filter, o)
	if (err != nil) {
		return nil, err
	}
	if err = cursor.All(Ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r*FlashCardRepository) GetByID(id string) (interface{}, error) {
	var result bson.M = make(primitive.M)

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	
	if err = r.Collection.FindOne(Ctx, bson.M{"_id": _id}).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *FlashCardRepository) Store(model *models.FlashCard) error {
	_, err := r.Collection.InsertOne(Ctx, model)
	return err
}

func (r *FlashCardRepository) Update(filter interface{}, update interface{}) (modifiedCount int64, err error) {
	result, err := r.Collection.UpdateOne(Ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r*FlashCardRepository) Delete(filter interface{}) (deletedCount int64, err error) {
	result, err := r.Collection.DeleteOne(Ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
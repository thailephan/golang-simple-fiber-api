package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromHex(id string) (objectId interface{}, err error){
	primitiveId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return primitiveId, nil
}

func NewObjectID() (objectId primitive.ObjectID) {
	return primitive.NewObjectID()
}

func BsonMapFromEntity(entity interface{}) (out Map, err error){
	data, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}

	_ = bson.Unmarshal(data, &out)
	return out, err
} 

type Map bson.M
type Document bson.D
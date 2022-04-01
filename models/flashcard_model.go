package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meaning struct {
	Type string `json:"type" bson:"type"` // nil, verb, noun
	Value string `json:"value" bson:"value"`
}

type FlashCard struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Word string `json:"word" bson:"word"`
	Meanings []Meaning `json:"meanings" bson:"meanings"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

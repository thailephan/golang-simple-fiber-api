package mocks

import "thailephan/flash-card-api/models"


var Flashcards = []models.FlashCard{
	{
		Word: "Tortoise",
		Meanings: []models.Meaning{
			{
				Type: "noun",
				Value: "Breast animal",
			},
			{
				Type: "noun",
				Value: "Four legs, live in water, swim well",
			},
		},
	},
	{
		Word: "Human",
		Meanings: []models.Meaning{
			{
				Type: "noun",
				Value: "Breast animal",
			},
			{
				Type: "noun",
				Value: "Four legs, live in water, swim well",
			},
		},
	},
}
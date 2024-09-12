package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Brainee struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Text   string             `json:"text,omitempty" validate:"required"`
	Author string             `json:"author,omitempty" validate:"required"`
	Brand  string             `json:"brand,omitempty" validate:"required"`
}

type CreateBrainee struct {
	Text   string `json:"text,omitempty" validate:"required"`
	Author string `json:"author,omitempty" validate:"required"`
	Brand  string `json:"brand,omitempty" validate:"required"`
}

type UpdateBrainee struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Text   *string `json:"text,omitempty"`
	Author *string `json:"author,omitempty"`
	Brand  *string `json:"brand,omitempty"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name" validate:"required"`
	ImageURL string             `bson:"image_url" json:"image_url"`
	Email    string             `bson:"email" validate:"required,email"`
	Password string             `bson:"password" validate:"required" json:"password"`
}

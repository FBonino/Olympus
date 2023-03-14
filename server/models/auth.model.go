package models

import "time"

type LoginInput struct {
	Identifier string `json:"identifier" bson:"identifier" binding:"required"`
	Password   string `json:"password" bson:"password" binding:"required"`
}

type SignupInput struct {
	Username        string    `json:"username" bson:"username" binding:"required"`
	Email           string    `json:"email" bson:"email" binding:"required"`
	Password        string    `json:"password" bson:"password" binding:"required"`
	PasswordConfirm string    `json:"passwordConfirm" bson:"passwordConfirm" binding:"required"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
}

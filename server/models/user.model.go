package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id" bson:"_id"`
	Username     string    `json:"username" bson:"username"`
	Email        string    `json:"email" bson:"email"`
	Avatar       string    `json:"avatar" bson:"avatar"`
	Password     string    `json:"password" bson:"password"`
	Status       uint8     `json:"status" bson:"status"`
	CustomStatus string    `json:"customStatus" bson:"customStatus"`
	Friends      []Friend  `json:"friends" bson:"friends"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Friend struct {
	ID       string `json:"id" bson:"_id"`
	Relation uint8  `json:"relation" bson:"relation"`
}

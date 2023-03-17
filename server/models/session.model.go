package models

import "time"

type Session struct {
	ID        string    `json:"id" bson:"_id"`
	UserID    string    `json:"userId" bson:"userId"`
	Token     string    `json:"token" bson:"token"`
	ExpireAt  time.Time `json:"expireAt" bson:"expireAt"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

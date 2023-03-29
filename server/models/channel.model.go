package models

import "time"

type Channel struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Type      string    `json:"type" bson:"type"`
	Topic     string    `json:"topic" bson:"topic"`
	Messages  []string  `json:"messages" bson:"messages"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type CreateChannelInput struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

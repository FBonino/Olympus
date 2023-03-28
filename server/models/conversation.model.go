package models

import "time"

type Conversation struct {
	ID        string    `json:"id" bson:"_id"`
	Avatar    string    `json:"avatar" bson:"avatar"`
	Users     []string  `json:"users" bson:"users"`
	Messages  []string  `json:"messages" bson:"messages"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

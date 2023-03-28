package models

import "time"

type Message struct {
	ID        string    `json:"id" bson:"_id"`
	Author    string    `json:"author" bson:"author"`
	Content   string    `json:"content" bson:"content"`
	IsDeleted bool      `json:"isDeleted" bson:"isDeleted"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

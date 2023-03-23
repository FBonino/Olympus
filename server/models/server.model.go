package models

import "time"

type Server struct {
	ID             string       `json:"id" bson:"_id"`
	Name           string       `json:"name" bson:"name"`
	Avatar         string       `json:"avatar" bson:"avatar"`
	Roles          []ServerRole `json:"roles" bson:"roles"`
	Users          []ServerUser `json:"users" bson:"users"`
	Channels       []Channel    `json:"channels" bson:"channels"`
	DefaultChannel string       `json:"defaultChannel" bson:"defaultChannel"`
	CreatedAt      time.Time    `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt" bson:"updatedAt"`
}

type ServerRole struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Color string `json:"color" bson:"color"`
}

type ServerUser struct {
	ID    string       `json:"id" bson:"_id"`
	Roles []ServerRole `json:"roles" bson:"roles"`
}

type Channel struct {
	ID       string    `json:"id" bson:"_id"`
	Name     string    `json:"name" bson:"name"`
	Type     string    `json:"type" bson:"type"`
	Topic    string    `json:"topic" bson:"topic"`
	Messages []Message `json:"messages" bson:"messages"`
}

type Message struct {
	ID        string `json:"id" bson:"_id"`
	UserID    string `json:"userId" bson:"userId"`
	Text      string `json:"text" bson:"text"`
	IsDeleted bool   `json:"isDeleted" bson:"isDeleted"`
}

type CreateServerInput struct {
	Name   string `json:"name" bson:"name" binding:"required"`
	Avatar string `json:"avatar" bson:"avatar"`
}

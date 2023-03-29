package models

import "time"

type Server struct {
	ID             string       `json:"id" bson:"_id"`
	Name           string       `json:"name" bson:"name"`
	Avatar         string       `json:"avatar" bson:"avatar"`
	Roles          []ServerRole `json:"roles" bson:"roles"`
	Users          []ServerUser `json:"users" bson:"users"`
	Channels       []string     `json:"channels" bson:"channels"`
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
	ID    string   `json:"id" bson:"_id"`
	Roles []string `json:"roles" bson:"roles"`
}

type CreateServerInput struct {
	Name     string   `json:"name" bson:"name" binding:"required"`
	Avatar   string   `json:"avatar" bson:"avatar"`
	Channels []string `json:"channels" bson:"channels"`
}

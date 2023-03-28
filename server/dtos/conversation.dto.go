package dtos

import (
	"server/models"
	"time"
)

type ConversationBasicDTO struct {
	ID     string    `json:"id" bson:"_id"`
	Avatar string    `json:"avatar" bson:"avatar"`
	Users  []UserDTO `json:"users" bson:"users"`
}

type ConversationDTO struct {
	ID        string       `json:"id" bson:"_id"`
	Avatar    string       `json:"avatar" bson:"avatar"`
	Users     []UserDTO    `json:"users" bson:"users"`
	Messages  []MessageDTO `json:"messages" bson:"messages"`
	CreatedAt time.Time    `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt" bson:"updatedAt"`
}

func MapConversationBasicDTO(conversation models.Conversation, users []*models.User) ConversationBasicDTO {
	return ConversationBasicDTO{
		ID:     conversation.ID,
		Avatar: conversation.Avatar,
		Users:  MapUsersDTO(users),
	}
}

func MapConversationDTO(conversation models.Conversation, users []*models.User, messages []*models.Message) ConversationDTO {
	return ConversationDTO{
		ID:        conversation.ID,
		Avatar:    conversation.Avatar,
		Users:     MapUsersDTO(users),
		Messages:  MapMessagesDTO(messages),
		CreatedAt: conversation.CreatedAt,
		UpdatedAt: conversation.UpdatedAt,
	}
}

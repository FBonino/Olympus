package dtos

import (
	"server/models"
	"time"
)

type MessageDTO struct {
	ID        string    `json:"id" bson:"_id"`
	Author    string    `json:"author" bson:"author"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func MapMessageDTO(message *models.Message) MessageDTO {
	return MessageDTO{
		ID:        message.ID,
		Author:    message.Author,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}

func MapMessagesDTO(messages []*models.Message) []MessageDTO {
	messagesDTO := []MessageDTO{}

	for _, message := range messages {
		messageDTO := MapMessageDTO(message)
		messagesDTO = append(messagesDTO, messageDTO)
	}

	return messagesDTO
}

package services

import "server/models"

type ConversationService interface {
	FindUserConversations(string) ([]*models.Conversation, []string, error)
	FindOrCreate(*models.CreateConversationInput) (*models.Conversation, error)
	FindByID(string) (*models.Conversation, error)
	AddMessage(string, string) error
}

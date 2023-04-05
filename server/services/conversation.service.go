package services

import "server/models"

type ConversationService interface {
	FindUserConversations(string) ([]*models.Conversation, []string, error)
	Create(*models.CreateConversationInput) (*models.Conversation, error)
	FindByID(string) (*models.Conversation, error)
	AddMessage(string, string) error
}

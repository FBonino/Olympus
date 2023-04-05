package services

import "server/models"

type MessageService interface {
	Create(string, *models.CreateMessageInput) (*models.Message, error)
	FindMessages([]string, int64) ([]*models.Message, error)
}

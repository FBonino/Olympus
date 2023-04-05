package services

import (
	"server/models"
)

type UserService interface {
	FindByID(string) (*models.User, error)
	FindByIdentifier(string) (*models.User, error)
	FindManyByID([]string) ([]*models.User, error)
}

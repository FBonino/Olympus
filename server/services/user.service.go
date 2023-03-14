package services

import "server/models"

type UserService interface {
	FindUserByID(string) (*models.User, error)
	FindUserByIdentifier(string) (*models.User, error)
}

package services

import "server/models"

type AuthService interface {
	Login(*models.LoginInput) (*models.User, error)
	Signup(*models.SignupInput) (*models.User, error)
}

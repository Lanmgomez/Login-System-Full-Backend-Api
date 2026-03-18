package auth

import "login_system_complete_api/internal/models"

type ServiceInterface interface {
	Login(input models.LoginInput) (*models.LoginResponse, error)
}
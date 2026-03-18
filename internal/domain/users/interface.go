package users

import "login_system_complete_api/internal/models"

type RepositoryInterface interface {
	FindAllPaginated(limit, offset int) ([]models.User, int, error)
}

type ServiceInterface interface {
	GetUsers(page, limit int) ([]models.User, int, error)
}
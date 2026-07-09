package postgre

import (
	"context"
	usecasemodels "go-skeleton/usecases/v1/models"
)

// UserRepository defines the contract for User database operations in PostgreSQL.
type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*usecasemodels.UserDTO, error)
	Create(ctx context.Context, user *usecasemodels.UserDTO) (*usecasemodels.UserDTO, error)
}

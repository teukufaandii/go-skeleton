package database

import (
	"context"
	dbmodels "go-skeleton/database/postgre/models"
)

// UserRepository defines the contract for User database operations.
type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*dbmodels.UserDB, error)
	Create(ctx context.Context, user *dbmodels.UserDB) (*dbmodels.UserDB, error)
}

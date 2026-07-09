package usecase

import (
	"context"
	"go-skeleton/database/postgre"
	usecasemodels "go-skeleton/usecases/v1/models"
)

// UserUsecase defines the business logic contract for User entity.
type UserUsecase interface {
	RegisterUser(ctx context.Context, user *usecasemodels.UserDTO) (*usecasemodels.UserDTO, error)
	GetProfile(ctx context.Context, id uint) (*usecasemodels.UserDTO, error)
}

// userUsecase is the concrete implementation of UserUsecase interface.
type userUsecase struct {
	repo postgre.UserRepository
}

// NewUserUsecase creates a new instance of UserUsecase.
func NewUserUsecase(repo postgre.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

// RegisterUser handles the business logic for registering a new user.
func (u *userUsecase) RegisterUser(ctx context.Context, user *usecasemodels.UserDTO) (*usecasemodels.UserDTO, error) {
	// Call repository to save the user, mapping happens at repo implementation level
	return u.repo.Create(ctx, user)
}

// GetProfile retrieves a user's profile by ID.
func (u *userUsecase) GetProfile(ctx context.Context, id uint) (*usecasemodels.UserDTO, error) {
	// Call repository, mapping happens at repo implementation level
	return u.repo.FindByID(ctx, id)
}

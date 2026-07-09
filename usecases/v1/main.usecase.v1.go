package usecase

import (
	"context"
	"go-skeleton/database/postgre"
	dbmodels "go-skeleton/database/postgre/models"
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
	// 1. Convert DTO to DB model using the previously created mapper
	userDB := dbmodels.NewUserDB(user)

	// 2. Call repository to save the user
	createdDB, err := u.repo.Create(ctx, userDB)
	if err != nil {
		return nil, err
	}

	// 3. Map back to domain model
	return createdDB.ToDomain(), nil
}

// GetProfile retrieves a user's profile by ID.
func (u *userUsecase) GetProfile(ctx context.Context, id uint) (*usecasemodels.UserDTO, error) {
	// 1. Call repository
	userDB, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. Map to domain DTO using .ToDomain() method
	return userDB.ToDomain(), nil
}

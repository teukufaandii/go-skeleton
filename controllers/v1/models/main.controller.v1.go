package models

import (
	usecasemodels "go-skeleton/usecases/v1/models"
)

// UserRequest represents the JSON payload for creating or updating a user.
type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserResponse represents the JSON response returned to the client.
type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

// ToDomain converts the HTTP request model to the Domain DTO.
func (req *UserRequest) ToDomain() *usecasemodels.UserDTO {
	return &usecasemodels.UserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

// NewUserResponse converts the Domain DTO to the HTTP response model.
func NewUserResponse(dto *usecasemodels.UserDTO) *UserResponse {
	if dto == nil {
		return nil
	}
	return &UserResponse{
		ID:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
	}
}

package models

import (
	"time"

	usecasemodels "go-skeleton/usecases/v1/models"
	"gorm.io/gorm"
)

// UserDB represents the schema for the `users` table in PostgreSQL.
type UserDB struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the default table name.
func (UserDB) TableName() string {
	return "users"
}

// ToDomain converts the DB model to the Domain DTO.
func (db *UserDB) ToDomain() *usecasemodels.UserDTO {
	return &usecasemodels.UserDTO{
		ID:        db.ID,
		Name:      db.Name,
		Email:     db.Email,
		Password:  db.Password,
		CreatedAt: db.CreatedAt.Unix(),
		UpdatedAt: db.UpdatedAt.Unix(),
	}
}

// NewUserDB creates a DB model from a Domain DTO.
func NewUserDB(dto *usecasemodels.UserDTO) *UserDB {
	if dto == nil {
		return nil
	}

	var createdAt, updatedAt time.Time
	if dto.CreatedAt != 0 {
		createdAt = time.Unix(dto.CreatedAt, 0)
	}
	if dto.UpdatedAt != 0 {
		updatedAt = time.Unix(dto.UpdatedAt, 0)
	}

	return &UserDB{
		ID:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

package models

// UserDTO represents the core domain entity for a user.
// It is pure and has no framework or database specific tags.
type UserDTO struct {
	ID        uint
	Name      string
	Email     string
	Password  string // Hashed password
	CreatedAt int64
	UpdatedAt int64
}

package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users repository
type Users struct {
	db *sql.DB
}

// NewRepositoryUsers create new repository of the users
func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Create new user into DB
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}

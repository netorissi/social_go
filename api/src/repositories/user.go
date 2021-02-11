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
	query := `INSERT INTO users (name, nick, email, password) values (?, ?, ?, ?);`

	resp, err := u.db.Prepare(query)
	if err != nil {
		return 0, nil
	}

	defer resp.Close()

	result, err := resp.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

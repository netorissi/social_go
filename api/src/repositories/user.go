package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Find users
func (u Users) Find(param string) ([]models.User, error) {
	param = fmt.Sprintf("%%%s%%", param)

	query := `
		SELECT id, name, nick, email, create_at 
		FROM users
		WHERE name LIKE ?
		OR nick LIKE ?;
	`

	resp, err := u.db.Query(query, param, param)
	if err != nil {
		return nil, nil
	}

	defer resp.Close()

	var users []models.User

	for resp.Next() {
		var user models.User

		if err = resp.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

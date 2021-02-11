package models

import (
	"errors"
	"strings"
	"time"
)

// User -
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

// BeforeSave validate fields before save into DB
func (u *User) BeforeSave() (err error) {
	u.format()
	err = u.validate()
	return err
}

func (u *User) validate() (err error) {
	if len(u.Name) == 0 {
		err = errors.New("field name is required")
	} else if len(u.Nick) == 0 {
		err = errors.New("field nick is required")
	} else if len(u.Email) == 0 {
		err = errors.New("field email is required")
	} else if len(u.Password) == 0 {
		err = errors.New("field password is required")
	}
	return err
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}

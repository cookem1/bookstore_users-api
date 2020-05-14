package users

import (
	"strings"

	"github.com/cookem1/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `"-"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Password = strings.TrimSpace(strings.ToLower(user.Password))

	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}

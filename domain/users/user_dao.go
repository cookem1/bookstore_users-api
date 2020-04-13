package users

import (
	"fmt"

	"github.com/cookem1/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.DateCreated = result.DateCreated
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email address %s is already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("Users %d already exists", user.Id))
	}

	userDB[user.Id] = user
	return nil
}

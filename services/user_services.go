package services

import (
	"github.com/cookem1/bookstore_users-api/domain/users"
	"github.com/cookem1/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	//fmt.Println("Searching for user:", userID)
	result := &users.User{Id: userID}
	err := result.Get()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SearchUser() {}

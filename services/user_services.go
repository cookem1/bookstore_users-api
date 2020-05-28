package services

import (
	"github.com/cookem1/bookstore_users-api/domain/users"
	"github.com/cookem1/bookstore_users-api/utils/crypto_utils"
	"github.com/cookem1/bookstore_users-api/utils/date_utils"
	"github.com/cookem1/bookstore_users-api/utils/errors"
)

const (
	salt = "unclematsabitsalty"
)

var (
	UsersService usersServiceInterface = &userService{}
)

type userService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	SearchUser(string) (users.Users, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(userID int64) *errors.RestErr
}

func (u *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password, salt)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userID}
	err := result.Get()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (u *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err = user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.Email = user.Email
		currentUser.LastName = user.LastName
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (u *userService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{Id: userID}
	return user.Delete()
}

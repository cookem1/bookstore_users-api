package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cookem1/bookstore_users-api/domain/users"
	"github.com/cookem1/bookstore_users-api/services"
	"github.com/cookem1/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("Invalid user ID (not a number): %v", userIDParam))
	}
	return userId, nil
}

func Get(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
	var user users.User
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func Delete(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	delErr := services.UsersService.DeleteUser(userId)
	if delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

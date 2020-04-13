package app

import (
	"github.com/cookem1/bookstore_users-api/controllers/ping"
	"github.com/cookem1/bookstore_users-api/controllers/users"
)

func mapURL() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/find", controllers.SearchUser)
	router.POST("/users/", users.CreateUser)

}

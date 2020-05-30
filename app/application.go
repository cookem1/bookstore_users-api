package app

import (
	"github.com/cookem1/bookstore_users-api/domain/users/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURL()
	logger.Info("About to start the application")
	router.Run(":8082")

}


package router

import (
	"social_media/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/users", userHandler.CreateUser)

	return router
}

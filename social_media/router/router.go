package router

import (
	"social_media/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/users", userHandler.CreateUser)

	userGroup := router.Group("/users/:username")
	{
		userGroup.POST("/tweets", userHandler.PostTweet)
		userGroup.GET("/tweets", userHandler.DisplayTweets)
		userGroup.GET("/tweets/search", userHandler.SearchTweets)

		userGroup.POST("/follow", userHandler.FollowUser)
		userGroup.POST("/approve-follow", userHandler.ApproveFollowRequest)
	}

	return router
}

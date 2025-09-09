package handlers

import (
	"net/http"
	"social_media/domain/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Privacy  string `json:"privacy" binding:"required,oneof=public private"`
}

type PostTweetRequest struct {
	Message string `json:"message" binding:"required"`
}

type FollowUserRequest struct {
	Followee string `json:"followee" binding:"required"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUseCase.CreateUser(req.Username, req.Privacy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) PostTweet(c *gin.Context) {
	username := c.Param("username")
	var req PostTweetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUseCase.PostTweet(username, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tweet posted successfully"})

}

func (h *UserHandler) FollowUser(c *gin.Context) {
	follower := c.Param("username")
	var req FollowUserRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUseCase.FollowUser(follower, req.Followee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow request sent successfully"})
}

func (h *UserHandler) ApproveFollowRequest(c *gin.Context) {
	followee := c.Param("username")
	follower := c.Query("follower")

	err := h.userUseCase.ApproveFollowRequest(follower, followee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow request approved successfully"})
}

func (h *UserHandler) DisplayTweets(c *gin.Context) {
	username := c.Param("username")

	tweets, err := h.userUseCase.DisplayTweets(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

func (h *UserHandler) SearchTweets(c *gin.Context) {
	username := c.Param("username")
	searchWord := c.Query("q")

	tweets, err := h.userUseCase.SearchTweets(username, searchWord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

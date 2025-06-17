package usecase

import (
	"social_media/domain/entity"
	"social_media/domain/service"
)

type UserUseCase interface {
	CreateUser(username, privacy string) (*entity.User, error)
	PostTweet(username, message string) error
}

type userUseCase struct {
	userService  service.UserService
	tweetService service.TweetService
}

func NewUserUseCase(userService service.UserService, tweetService service.TweetService) UserUseCase {
	return &userUseCase{
		userService:  userService,
		tweetService: tweetService,
	}
}

func (u *userUseCase) CreateUser(username, privacy string) (*entity.User, error) {
	return u.userService.CreateUser(username, privacy)
}

func (u *userUseCase) PostTweet(username, message string) error {
	return u.tweetService.CreateTweet(username, message)
}

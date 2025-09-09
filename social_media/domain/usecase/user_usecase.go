package usecase

import (
	"errors"
	"social_media/domain/entity"
	"social_media/domain/service"
)

type UserUseCase interface {
	CreateUser(username, privacy string) (*entity.User, error)
	PostTweet(username, message string) error
	FollowUser(follower, followee string) error
	ApproveFollowRequest(follower, followee string) error
	DisplayTweets(username string) ([]entity.Tweet, error)
	SearchTweets(username, searchWord string) ([]entity.Tweet, error)
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

func (u *userUseCase) FollowUser(follower, followee string) error {
	followeeUser, err := u.userService.GetUser(followee)
	if err != nil {
		return err
	}

	followerUser, err := u.userService.GetUser(follower)
	if err != nil {
		return err
	}

	if followeeUser.IsPublic() {
		followeeUser.AddFollower(follower)
		followerUser.AddFollowing(followee)
		return u.userService.UpdateUser(followeeUser)
	}

	followerUser.AddFollowing(followee)
	return u.userService.UpdateUser(followerUser)
}

func (u *userUseCase) ApproveFollowRequest(follower, followee string) error {
	followeeUser, err := u.userService.GetUser(followee)
	if err != nil {
		return err
	}

	followerUser, err := u.userService.GetUser(follower)
	if err != nil {
		return err
	}

	if followeeUser.IsPrivate() && !followeeUser.HasFollower(follower) {
		followeeUser.AddFollower(follower)
		followerUser.AddFollowing(followee)
		return u.userService.UpdateUser(followeeUser)
	}

	return errors.New("no follow request found")
}

func (u *userUseCase) DisplayTweets(username string) ([]entity.Tweet, error) {
	return u.tweetService.GetUserTweets(username)
}

func (u *userUseCase) SearchTweets(username, searchWord string) ([]entity.Tweet, error) {
	return u.tweetService.SearchTweets(username, searchWord)
}

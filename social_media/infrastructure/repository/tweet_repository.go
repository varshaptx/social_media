package repository

import (
	"errors"
	"social_media/domain/entity"
)

type tweetRepository struct {
	users map[string]*entity.User
}

func NewTweetRepository(users map[string]*entity.User) *tweetRepository {
	return &tweetRepository{
		users: users,
	}
}

func (r *tweetRepository) CreateTweet(username string, tweet *entity.Tweet) error {
	user, exists := r.users[username]
	if !exists {
		return errors.New("user not found")
	}
	user.AddTweet(*tweet)
	return nil
}

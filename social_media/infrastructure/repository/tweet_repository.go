package repository

import (
	"errors"
	"social_media/domain/entity"
	"strings"
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

func (r *tweetRepository) GetUserTweets(username string) ([]entity.Tweet, error) {
	user, exists := r.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user.Tweets, nil
}

func (r *tweetRepository) SearchTweets(username, searchWord string) ([]entity.Tweet, error) {
	user, exists := r.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}

	var matchingTweets []entity.Tweet
	for _, tweet := range user.Tweets {
		if len(matchingTweets) >= 10 {
			break
		}
		if strings.Contains(tweet.Message, searchWord) {
			matchingTweets = append(matchingTweets, tweet)
		}
	}

	return matchingTweets, nil
}

func (r *tweetRepository) DeleteTweet(username string, tweetIndex int) error {
	user, exists := r.users[username]
	if !exists {
		return errors.New("user not found")
	}

	if tweetIndex < 0 || tweetIndex >= len(user.Tweets) {
		return errors.New("invalid tweet index")
	}

	user.Tweets = append(user.Tweets[:tweetIndex], user.Tweets[tweetIndex+1:]...)
	return nil
}

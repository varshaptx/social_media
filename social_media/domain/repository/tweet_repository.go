package repository

import "social_media/domain/entity"

type TweetRepository interface {
	CreateTweet(username string, tweet *entity.Tweet) error
	GetUserTweets(username string) ([]entity.Tweet, error)
	SearchTweets(username, searchWord string) ([]entity.Tweet, error)
	DeleteTweet(username string, tweetIndex int) error
}

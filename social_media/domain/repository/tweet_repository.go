package repository

import "social_media/domain/entity"

type TweetRepository interface {
	CreateTweet(username string, tweet *entity.Tweet) error
}

package service

import (
	"errors"
	"social_media/domain/entity"
	"social_media/domain/repository"
)

type TweetService interface {
	CreateTweet(username string, message string) error
	GetUserTweets(username string) ([]entity.Tweet, error)
	SearchTweets(username, searchWord string) ([]entity.Tweet, error)
	DeleteTweet(username string, tweetIndex int) error
	ValidateTweet(tweet *entity.Tweet) error
}

type tweetService struct {
	tweetRepo   repository.TweetRepository
	userService UserService
}

func NewTweetService(tweetRepo repository.TweetRepository, userService UserService) TweetService {
	return &tweetService{
		tweetRepo:   tweetRepo,
		userService: userService,
	}
}

func (s *tweetService) CreateTweet(username string, message string) error {
	_, err := s.userService.GetUser(username)
	if err != nil {
		return err
	}

	tweet := entity.NewTweet(username, message)
	if err := s.ValidateTweet(tweet); err != nil {
		return err
	}

	return s.tweetRepo.CreateTweet(username, tweet)
}

func (s *tweetService) GetUserTweets(username string) ([]entity.Tweet, error) {
	_, err := s.userService.GetUser(username)
	if err != nil {
		return nil, err
	}

	return s.tweetRepo.GetUserTweets(username)
}

func (s *tweetService) SearchTweets(username, searchWord string) ([]entity.Tweet, error) {
	_, err := s.userService.GetUser(username)
	if err != nil {
		return nil, err
	}

	return s.tweetRepo.SearchTweets(username, searchWord)
}

func (s *tweetService) DeleteTweet(username string, tweetIndex int) error {
	_, err := s.userService.GetUser(username)
	if err != nil {
		return err
	}

	return s.tweetRepo.DeleteTweet(username, tweetIndex)
}

func (s *tweetService) ValidateTweet(tweet *entity.Tweet) error {
	if tweet == nil {
		return errors.New("Tweet cannot be nil.")
	}
	if tweet.User == "" {
		return errors.New("Tweet Username cannot be empty.")
	}
	if tweet.Message == "" {
		return errors.New("Tweet Message cannot be nil.")
	}
	if !tweet.IsValid() {
		return errors.New("Tweet Message is more than 280 characters.")
	}
	return nil
}

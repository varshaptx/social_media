package entity

type User struct {
	Username  string
	Privacy   string
	Tweets    []Tweet
	Followers map[string]bool
	Following map[string]bool
}

func NewUser(username, privacy string) *User {
	return &User{
		Username:  username,
		Privacy:   privacy,
		Tweets:    []Tweet{},
		Followers: make(map[string]bool),
		Following: make(map[string]bool),
	}
}

func (u *User) IsPublic() bool {
	return u.Privacy == "public"
}

func (u *User) IsPrivate() bool {
	return u.Privacy == "private"
}

func (u *User) AddTweet(tweet Tweet) {
	u.Tweets = append(u.Tweets, tweet)
}

func (u *User) AddFollower(follower string) {
	u.Followers[follower] = true
}

func (u *User) AddFollowing(following string) {
	u.Followers[following] = true
}

func (u *User) HasFollower(follower string) bool {
	return u.Followers[follower]
}

func (u *User) HasFollowing(following string) bool {
	return u.Following[following]
}

package entity

type User struct {
	Username string
	Privacy  string
	Tweets   []Tweet
}

func NewUser(username, privacy string) *User {
	return &User{
		Username: username,
		Privacy:  privacy,
		Tweets:   []Tweet{},
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

package entity

type Tweet struct {
	User    string
	Message string
}

func NewTweet(user, message string) *Tweet {
	return &Tweet{
		User:    user,
		Message: message,
	}
}

func (t *Tweet) IsValid() bool {
	return len(t.Message) <= 280
}

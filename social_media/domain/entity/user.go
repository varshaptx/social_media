package entity

type User struct {
	Username string
	Privacy  string
}

func NewUser(username, privacy string) *User {
	return &User{
		Username: username,
		Privacy:  privacy,
	}
}

func (u *User) IsPublic() bool {
	return u.Privacy == "public"
}

func (u *User) IsPrivate() bool {
	return u.Privacy == "private"
}

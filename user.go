package host

import (
	"github.com/Permaweb/host/internal/github"
	"github.com/oklog/ulid"
)

// User is a collection of oAuth tokens.
type User struct {
	ulid.ULID
	GitHub *GitHubUser
}

// GitHubUser is a GitHub user.
type GitHubUser struct {
	github.User
	github.AccessToken
}

// UserService is a service that manages users.
type UserService interface {
	User(uuid string) (user *User, err error)
	Users() (users []*User, err error)
	Create(user *User) (err error)
	Delete(user *User) (err error)
	Edit(user *User) (err error)
	Repos(user *User) (repos []*Repo, err error)
}

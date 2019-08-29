package host

import (
	"github.com/Permaweb/host/internal/github"
	"github.com/google/uuid"
)

// User is a collection of oAuth tokens.
type User struct {
	uuid.UUID
	GitHub *GitHubUser
}

// GitHubUser is a GitHub user.
type GitHubUser struct {
	github.User
	github.AccessToken
}

// UserService is a service that manages users.
type UserService interface {
	User(uuid uuid.UUID) (user *User, err error)
	Users() (users []*User, err error)
	Create(user *User) (err error)
	Delete(user *User) (err error)
	Edit(user *User) (err error)
	Repos(user *User) (repos []*Repo, err error)
}

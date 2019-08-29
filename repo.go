package host

import "github.com/google/uuid"

// Repo is a user-submitted Git repository.
type Repo struct {
	uuid.UUID        // 1. Generate a UUID
	URL       string // 2. Download the repo
	IPFS      string // 3. Add it to IPFS
	Key       string // 4. Generate a key
	IPNS      string // 5. Add it to IPNS

	Owner   uuid.UUID // User
	Service string
}

// RepoService is a service that helps retreiving repos.
type RepoService interface {
	Repo(uuid uuid.UUID) (repo *Repo, err error)
	Repos(user *User) (repos []*Repo, err error)
	Owner(repo *Repo) (user *User, err error)
	Create(repo *Repo) (err error)
	Delete(repo *Repo) (err error)
	Update(repo *Repo) (err error)
}

package host

// Repo is a user-submitted Git repository.
type Repo struct {
	UUID string // 1. Generate a UUID
	URL  string // 2. Download the repo
	IPFS string // 3. Add it to IPFS
	Key  string // 4. Generate a key
	IPNS string // 5. Add it to IPNS

	Owner string // User
}

// RepoService is a service that helps retreiving repos.
type RepoService interface {
	Repo(url string) (repo *Repo, err error)
	Repos() (repos []*Repo, err error)
	CreateRepo(repo *Repo) (err error)
	DeleteRepo(repo *Repo) (err error)
}

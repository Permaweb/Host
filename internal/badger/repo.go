package badger

import (
	"github.com/Permaweb/host"
	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
)

/*
	Table

	repo:uuid           | Repo
	user:uuid:repo:uuid | Repo
*/

// RepoService is a BadgerDS implementation of host.RepoService.
type RepoService struct {
	*badger.DB
	host.UserService
}

const (
	repoPrefix = "repo:"
	repoSuffix = ":repo:"
)

// Repo gets a single repo.
func (s *RepoService) Repo(uuid uuid.UUID) (repo *host.Repo, err error) {
	err = s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(repoPrefix + uuid.String()))
		if err != nil {
			return err
		}
		return item.Value(func(bytes []byte) (err error) {
			*repo, err = decodeRepo(bytes)
			return
		})
	})

	return
}

// Repos gets all repos from a user.
func (s *RepoService) Repos(user *host.User) (repos []*host.Repo, err error) {
	s.DB.View(func(txn *badger.Txn) (err error) {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(userPrefix + user.UUID.String() + repoSuffix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			// k := string(item.Key())
			err := item.Value(func(v []byte) (err error) {

				// UUID
				uuid, err := decodeUUID(v)
				if err != nil {
					return err
				}

				// Repo
				repo, err := s.Repo(uuid)
				if err != nil {
					return err
				}

				// Repos
				repos = append(repos, repo)
				return err
			})
			if err != nil {
				return err
			}
		}
		return err
	})
	return repos, err
}

// Owner gets a repo's owner.
func (s *RepoService) Owner(repo *host.Repo) (user *host.User, err error) {
	return s.UserService.User(repo.UUID.String())
}

// Create a repo.
func (s *RepoService) Create(repo *host.Repo) (err error) {

	// UUID
	if repo.UUID == "" {
		buuid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		repo.UUID = buuid.String()
	}

	return s.Update(repo)
}

// Delete a repo.
func (s *RepoService) Delete(repo *host.Repo) (err error) {
	return
}

// Update a repo.
func (s *RepoService) Update(repo *host.Repo) (err error) {

	// Repo
	b, err := encodeRepo(*repo)
	if err != nil {
		return
	}

	// Update
	return s.DB.Update(func(txn *badger.Txn) (err error) {
		err = txn.Set([]byte(repoPrefix+repo.UUID), b)
		if err != nil {
			return err
		}

		// User
		return txn.Set([]byte(userPrefix+repo.Owner+repoSuffix+repo.UUID), b)
	})
}

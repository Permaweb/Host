package badger

import (
	"github.com/Permaweb/host"
	"github.com/dgraph-io/badger"
)

// RepoService is a BadgerDS implementation of host.RepoService.
type RepoService struct {
	DB *badger.DB
}

// PrefixRepo is a prefix used internally.
const repoPrefix = "repo:"

// Repo gets a single repo from its URL.
func (s *RepoService) Repo(url string) (repo *host.Repo, err error) {
	err = s.DB.View(func(txn *badger.Txn) error {

		// Get
		item, err := txn.Get([]byte(repoPrefix + url))
		if err != nil {
			return err
		}

		// Value
		return item.Value(func(bytes []byte) (err error) {
			*repo, err = decodeRepo(bytes)
			return
		})
	})
	return
}

package badger

import (
	"strings"

	"github.com/Permaweb/host"
	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
)

/*
	User Table

	user:uuid | User
*/

// UserService is a BadgerDS implementation of host.UserService.
type UserService struct {
	*badger.DB
}

const userPrefix = "user:"

// User gets a user.
func (s *UserService) User(uuid string) (user *host.User, err error) {
	err = s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(userPrefix + uuid))
		if err != nil {
			return err
		}
		return item.Value(func(bytes []byte) (err error) {
			*user, err = decodeUser(bytes)
			return
		})
	})
	return
}

// Users gets all users.
func (s *UserService) Users() (users []*host.User, err error) {
	s.DB.View(func(txn *badger.Txn) (err error) {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(userPrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()

			// Skip repos.
			if strings.Contains(string(item.Key()), repoSuffix) {
				continue
			}

			err = item.Value(func(v []byte) (err error) {

				// User
				user, err := decodeUser(v)
				if err != nil {
					return err
				}

				// Users
				users = append(users, &user)
				return err
			})
			if err != nil {
				return err
			}
		}
		return err
	})
	return users, err
}

// Create a user.
func (s *UserService) Create(user *host.User) (err error) {

	// UUID
	if user.UUID == "" {
		buuid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		user.UUID = buuid.String()
	}

	// User
	b, err := encodeUser(*user)
	if err != nil {
		return
	}

	// Update
	return s.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(userPrefix+user.UUID), b)
	})
}

// Delete a user.
func (s *UserService) Delete(user *host.User) (err error) {
	return s.DB.Update(func(txn *badger.Txn) (err error) {
		return txn.Delete([]byte(userPrefix + user.UUID))
	})
}

// Edit a user.
func (s *UserService) Edit(user *host.User) (err error) {

	b, err := encodeUser(*user)
	if err != nil {
		return
	}

	return s.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(userPrefix+user.ULID.String()), b)
	})
}

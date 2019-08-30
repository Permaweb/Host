package badger

import (
	"time"

	"github.com/Permaweb/host"
	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
)

// SessionService is a BadgerDS implementation of host.SessionService.
type SessionService struct {
	*badger.DB
	host.UserService
}

const (
	sessionPrefix = "session:"
	sessionTTL    = time.Hour * 8
)

// User gets a user using the Session Service
func (s *SessionService) User(key uuid.UUID) (user *host.User, err error) {
	err = s.DB.View(func(txn *badger.Txn) (err error) {
		item, err := txn.Get([]byte(sessionPrefix + key.String()))
		if err != nil {
			return err
		}
		return item.Value(func(bytes []byte) (err error) {
			uuid, err := decodeUUID(bytes)
			if err != nil {
				return err
			}

			// User
			user, err = s.UserService.User(uuid)
			return
		})
	})
	return
}

// Create a session.
func (s *SessionService) Create(user *host.User) (session *host.Session, err error) {

	// UUID
	key, err := uuid.NewRandom()
	if err != nil {
		return
	}

	// User
	value, err := encodeUUID(user.UUID)
	if err != nil {
		return
	}

	// Update
	err = s.DB.Update(func(txn *badger.Txn) (err error) {
		return txn.SetEntry(badger.NewEntry([]byte(sessionPrefix+key.String()), value).WithTTL(sessionTTL))
	})
	if err != nil {
		return
	}

	// Session
	session = &host.Session{
		Key:  key,
		User: user.UUID,
	}

	return
}

// Refresh a session.
func (s *SessionService) Refresh(old *host.Session) (session *host.Session, err error) {

	// User
	user, err := s.User(old.Key)
	if err != nil {
		return
	}

	// Delete
	err = s.Delete(old)
	if err != nil {
		return
	}

	// New
	return s.Create(user)
}

// Delete a session.
func (s *SessionService) Delete(session *host.Session) (err error) {
	return s.DB.Update(func(txn *badger.Txn) (err error) {
		return txn.Delete([]byte(sessionPrefix + session.Key.String()))
	})
}

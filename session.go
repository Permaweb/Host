package host

import "github.com/google/uuid"

// Session makes the link between a session and a user.
type Session struct {
	Key  uuid.UUID
	User uuid.UUID
}

// SessionService manages sessions.
type SessionService interface {
	User(key uuid.UUID) (user *User, err error)
	New(user User) (session *Session, err error)
	Refresh(old *Session) (session *Session, err error)
	Delete(session *Session) (err error)
}

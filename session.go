package host

// Session makes the link between a session and a user.
type Session struct {
	Key  string
	User string
}

// SessionService manages sessions.
type SessionService interface {
	User(key string) (user *User, err error)
	New(user User) (session *Session, err error)
	Refresh(old *Session) (session *Session, err error)
	Delete(session *Session) (err error)
}

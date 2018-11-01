package domain

import "net/http"

// Session data
type Session struct {
}

// SessionDriver is interface
type SessionDriver interface {
	New(cookies []*http.Cookie) *Session
	Write(session *Session, object interface{}) error
	Read(session *Session, object interface{}) error
}

// Auth interface
type Auth interface {
	Login(username string, password string) *Session
	Logout(session *Session) error
}

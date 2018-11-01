package domain

import "net/http"

// Session interface
type Session interface {
	Set(cookies []*http.Cookie)
	Write(object interface{}) error
	Read(object interface{}) error
}

// Auth interface
type Auth interface {
	Login(username string, password string) error
	Logout() error
}

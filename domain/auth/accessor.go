package auth

// Accessor interface
type Accessor interface {
	Login(user string, password string) (*Session, error)
	Logout(s *Session) error
}

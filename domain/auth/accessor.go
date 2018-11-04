package auth

// Accessor interface
type Accessor interface {
	Login(user string, password string)
	Logout(s *Session)
}

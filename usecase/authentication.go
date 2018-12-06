package usecase

// Authentication :
type Authentication interface {
	Login(username string, password string)
	Logout()
}

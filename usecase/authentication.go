package usecase

// Authentication :
type Authentication interface {
	Login(param Parameter)
	Logout()
}

package usecase

import (
	"github.com/rucm/cptools/usecase/request"
)

// AuthInteractor : auth input port
type AuthInteractor struct {
	Handler usecase.RequestHandler
	Repo    usecase.SessionRepository
}

// Login : login to site
func (interactor *AuthInteractor) Login(username string, password string) {

}

// Logout : logout to site
func (interactor *AuthInteractor) Logout() {

}

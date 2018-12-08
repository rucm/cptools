package atcoder

import (
	"github.com/rucm/cptools/usecase"
)

// AuthenticationInteractor :
type AuthenticationInteractor struct {
	Repo usecase.SessionRepository
}

// Login : AuthenticationInteractor
func (interactor *AuthenticationInteractor) Login(username string, password string) {

}

// Logout : AuthenticationInteractor
func (interactor *AuthenticationInteractor) Logout() {

}

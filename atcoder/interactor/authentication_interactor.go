package atcoder

import (
	"github.com/rucm/cptools/usecase"
)

// AuthenticationInteractor :
type AuthenticationInteractor struct {
	Handler usecase.RequestHandler
	Repo    usecase.SessionRepository
}

// Login : AuthenticationInteractor
func (interactor *AuthenticationInteractor) Login(param usecase.Parameter) {

	response := interactor.Handler.Execute(param)
	interactor.Repo.Write(response.Session)
}

// Logout : AuthenticationInteractor
func (interactor *AuthenticationInteractor) Logout() {

	interactor.Repo.Remove()
}

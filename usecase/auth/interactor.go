package usecase

import (
	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase/request"
)

// AuthInteractor : auth input port
type AuthInteractor struct {
	Handler usecase.RequestHandler
	Repo    usecase.SessionRepository
}

// Login : login to site
func (interactor *AuthInteractor) Login() {

	res := interactor.Handler.Execute()

	session := &domain.Session{}
	res.Bind(session)

	interactor.Repo.Write(session)
}

// Logout : logout to site
func (interactor *AuthInteractor) Logout() {

	interactor.Repo.Remove()
}

// Session : if state in login, return session
func (interactor *AuthInteractor) Session() *domain.Session {

	return interactor.Repo.Read()
}

package usecase

import (
	"github.com/rucm/cptools/domain"
)

// AuthInteractor : auth input port
type AuthInteractor struct {
	repo SessionRepository
}

// Set : set session
func (interactor *AuthInteractor) Set(session *domain.Session) {

	interactor.repo.Write(session)
}

// Get : get session
func (interactor *AuthInteractor) Get() *domain.Session {

	return interactor.repo.Read()
}

// Del : remove session
func (interactor *AuthInteractor) Del() {

	interactor.repo.Remove()
}

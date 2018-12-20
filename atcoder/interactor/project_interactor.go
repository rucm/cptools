package atcoder

import (
	"github.com/rucm/cptools/usecase"
)

// ProjectInteractor :
type ProjectInteractor struct {
	Handler     usecase.RequestHandler
	ContestRepo usecase.ContestRepository
	SessionRepo usecase.SessionRepository
}

// Load :
func (interactor *ProjectInteractor) Load(contestID string) {

}

// Expand :
func (interactor *ProjectInteractor) Expand(problemID string) {

}

// ExpandAll :
func (interactor *ProjectInteractor) ExpandAll() {

}

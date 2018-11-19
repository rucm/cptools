package auth

import (
	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase/auth"
)

// FileRepository struct
type FileRepository struct {
	filaname string
}

// NewFileRepository : repository
func NewFileRepository(filename string) usecase.AuthRepository {

	repo := new(FileRepository)
	repo.filaname = filename

	return repo
}

func (repo *FileRepository) Write(session *domain.Session) {

}

func (repo *FileRepository) Read() *domain.Session {

	return &domain.Session{}
}

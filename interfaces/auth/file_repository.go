package auth

import (
	"encoding/gob"
	"log"
	"os"

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
	file, err := os.Create(repo.filaname)

	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(session)
}

func (repo *FileRepository) Read() *domain.Session {
	file, err := os.Open(repo.filaname)

	if err != nil {
		log.Fatal(err)
	}

	session := &domain.Session{}
	decoder := gob.NewDecoder(file)
	decoder.Decode(session)

	return session
}

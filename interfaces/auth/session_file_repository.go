package interfaces

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase/auth"
)

// SessionFileRepository struct
type SessionFileRepository struct {
	filaname string
}

// NewSessionFileRepository : repository
func NewSessionFileRepository(filename string) usecase.SessionRepository {

	repo := new(SessionFileRepository)
	repo.filaname = filename

	return repo
}

func (repo *SessionFileRepository) Write(session *domain.Session) {
	file, err := os.Create(repo.filaname)

	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(session)
}

func (repo *SessionFileRepository) Read() *domain.Session {
	file, err := os.Open(repo.filaname)

	if err != nil {
		log.Fatal(err)
	}

	session := &domain.Session{}
	decoder := gob.NewDecoder(file)
	decoder.Decode(session)

	return session
}

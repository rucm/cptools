package atcoder

import (
	"github.com/rucm/cptools/domain"
	"github.com/rucm/cptools/usecase"
)

// ContestFileRepository :
type ContestFileRepository struct {
	Filename string
}

// NewContestFileRepository :
func NewContestFileRepository(filename string) usecase.ContestRepository {

	repo := &ContestFileRepository{}
	repo.Filename = filename

	return repo
}

// Write :
func (repo *ContestFileRepository) Write(contest *domain.Contest) {

}

// Read :
func (repo *ContestFileRepository) Read() *domain.Contest {

	return nil
}

// Remove :
func (repo *ContestFileRepository) Remove() {

}

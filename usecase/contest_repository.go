package usecase

import (
	"github.com/rucm/cptools/domain"
)

// ContestRepository :
type ContestRepository interface {
	Write(contest *domain.Contest)
	Read() *domain.Contest
	Remove()
}

package usecase

import (
	"github.com/rucm/cptools/domain"
)

// Query :
type Query interface {
	Download(contestID string) *domain.Contest
}

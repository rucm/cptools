package usecase

import "github.com/rucm/cptools/domain"

// Project :
type Project interface {
	GetContest(contestID string) *domain.Contest
	GetProblem(problemID string) *domain.Problem
}

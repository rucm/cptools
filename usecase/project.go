package usecase

// Project :
type Project interface {
	Load(contestID string)
	// GetContest(contestID string) *domain.Contest
	// GetProblem(problemID string) *domain.Problem
}

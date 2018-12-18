package usecase

// Project :
type Project interface {
	Load(contestID string)
	Expand(problemID string)
	ExpandAll()
}

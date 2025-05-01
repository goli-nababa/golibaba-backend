package domain

type RepositoryFilter struct {
	Field    string
	Operator string
	Value    string
}

type RepositorySort struct {
	Field    string
	SortType string
}

type RepositoryRequest struct {
	Filters  []*RepositoryFilter
	Sorts    []*RepositorySort
	Preloads []string
	Limit    uint
	Offset   uint
}

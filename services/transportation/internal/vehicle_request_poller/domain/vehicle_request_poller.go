package domain

type PollerRequest struct {
	TotalRecords   int
	BatchSize      int
	ConcurrentJobs int
}

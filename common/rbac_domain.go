package common

type TypeStatementAction uint8
type StatementID uint64

const (
	StatementActionAllow TypeStatementAction = iota + 1
	StatementActionDeny
)
package ice

import query "github.com/ice-judge/query-parser"

// OrderDirection describes the direction of oderring
type OrderDirection bool

const (
	// OrderDirectionUp up
	OrderDirectionUp = iota
	// OrderDirectionDown down
	OrderDirectionDown
)

// QueryField descirbe a field that would be the standard for ordering
type QueryField struct {
	Name   string
	DBName string
	Type   string
}

type queriable interface {
	PossibleOrderBy() []string
	PossibleFindBy() []string
	QueryFieldByName(name string) *QueryField
}

// Pagination describes how the page will restrict the data
type Pagination struct {
	Page           uint
	MaxEntities    uint
	OrderDirection OrderDirection
	OrderBy        *QueryField
	Condition      *query.Condition
}

// IDEA would use map to achieve more performance. but seems trivial
func validateOrderBy(orderby string, possible []string) bool {
	for _, item := range possible {
		if item == orderby {
			return true
		}
	}
	return false
}

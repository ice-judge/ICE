package ice

// SolveStatistics is statistics about the users' attempts to the problem
type SolveStatistics struct {
	Solved    uint `json:"solved"`
	Total     uint `json:"total"`
	SolveRate uint `json:"rate"`
}

// ProblemHeaders is a set of ProblemHeader
type ProblemHeaders []ProblemHeader

// PosssibleOrderBy queriable
func (phs ProblemHeaders) PosssibleOrderBy() []string {
	return []string{"id", "name"}
}

// PossibleFindBy queriable
func (phs ProblemHeaders) PossibleFindBy() []string {
	return []string{}
}

// QueryFieldByName queriable
func (phs ProblemHeaders) QueryFieldByName(name string) *QueryField {
	switch name {
	case "id":
		return &QueryField{
			Name: "id",
		}
	case "name":
		return &QueryField{
			Name: "name",
		}
	default:
		return nil
	}
}

// ProblemHeader is a brief info of Problem
type ProblemHeader struct {
	Title           string          `json:"title"`
	ID              uint            `json:"id"`
	SolveStatistics SolveStatistics `json:"statistics"`
	Difficulty      string          `json:"difficulty"`
	Tags            []string        `json:"tags"`
}

// GetProblemHeaders fetchs ProblemHeaders that quilifies the condition
// -10 ~ 10
func (ice *ICE) GetProblemHeaders(pagination Pagination) (int, ProblemHeaders, error) {
	return 0, []ProblemHeader{
			ProblemHeader{
				Title: "ㅁㅇㄹㄴㅁㄹ",
				ID:    1234,
			},
		},
		nil
}

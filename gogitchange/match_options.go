package gogitchange

type MatchOptions struct {
	matchType string // ".go" / ".txt". match the file name extension
	matchPath func(string) bool
}

func NewMatchOptions() *MatchOptions {
	return &MatchOptions{}
}

func (m *MatchOptions) MatchType(fileExtension string) *MatchOptions {
	m.matchType = fileExtension
	return m
}

func (m *MatchOptions) MatchPath(matchPath func(path string) bool) *MatchOptions {
	m.matchPath = matchPath
	return m
}

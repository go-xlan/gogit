package gogitchange

type Options struct {
	matchType string // ".go" / ".txt". match the file name extension
	matchPath func(string) bool
}

func NewOptions() *Options {
	return &Options{}
}

func (m *Options) MatchType(fileExtension string) *Options {
	m.matchType = fileExtension
	return m
}

func (m *Options) MatchPath(matchPath func(path string) bool) *Options {
	m.matchPath = matchPath
	return m
}

package gogitchange

// MatchOptions configures file matching criteria for changed file processing
// Provides flexible filtering by file extension and custom path matching functions
// Supports fluent configuration pattern for convenient setup
//
// MatchOptions 配置用于变更文件处理的文件匹配条件
// 提供通过文件扩展名和自定义路径匹配函数的灵活过滤
// 支持流畅配置模式以便于设置
type MatchOptions struct {
	matchType string            // File extension screen like ".go", ".txt" // 文件扩展名过滤器，如 ".go", ".txt"
	matchPath func(string) bool // Custom path matching function // 自定义路径匹配函数
}

// NewMatchOptions creates a new instance with default blank matching criteria
// Returns MatchOptions prepared for fluent configuration chaining
// Use with MatchType and MatchPath for complete setup
//
// NewMatchOptions 创建带默认空匹配条件的新实例
// 返回准备进行流畅配置链式调用的 MatchOptions
// 与 MatchType 和 MatchPath 一起使用以完成设置
func NewMatchOptions() *MatchOptions {
	return &MatchOptions{}
}

// MatchType sets file extension screen and returns updated MatchOptions
// Enables filtering files by extension like ".go", ".txt", ".md"
// Supports fluent configuration pattern for method chaining
//
// MatchType 设置文件扩展名过滤器并返回更新的 MatchOptions
// 支持按扩展名过滤文件，如 ".go", ".txt", ".md"
// 支持流畅配置模式进行方法链式调用
func (m *MatchOptions) MatchType(fileExtension string) *MatchOptions {
	m.matchType = fileExtension
	return m
}

// MatchPath sets custom path matching function and returns updated MatchOptions
// Enables advanced path filtering using custom logic and criteria
// Function receives absolute path and should return true for matching files
//
// MatchPath 设置自定义路径匹配函数并返回更新的 MatchOptions
// 支持使用自定义逻辑和条件进行高级路径过滤
// 函数接收绝对路径并对匹配文件返回 true
func (m *MatchOptions) MatchPath(matchPath func(path string) bool) *MatchOptions {
	m.matchPath = matchPath
	return m
}

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-xlan/gogit/release.yml?branch=main&label=BUILD)](https://github.com/go-xlan/gogit/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/gogit)](https://pkg.go.dev/github.com/go-xlan/gogit)
[![Coverage Status](https://img.shields.io/coveralls/github/go-xlan/gogit/main.svg)](https://coveralls.io/github/go-xlan/gogit?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-xlan/gogit.svg)](https://github.com/go-xlan/gogit/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/gogit)](https://goreportcard.com/report/github.com/go-xlan/gogit)

# gogit

Enhanced Git operations toolkit providing streamlined repo management with comprehensive commit and remote sync capabilities.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Core Features

🎯 **Streamlined Git Operations**: Intelligent staging, committing, and status checking with comprehensive API
⚡ **Smart Commit Management**: Auto staging with commit and amend support, prevents unsafe operations
🔄 **Remote Push Detection**: Automatic checking of commit push status across multiple remotes
🌍 **Cross-Platform Support**: Pure Go implementation without CLI dependencies using go-git foundation
📋 **Fluent API Design**: Builder pattern for convenient configuration and method chaining

## Related Projects

- **[gitgo](https://github.com/go-xlan/gitgo)** - Streamlined Git command execution engine with fluent chaining interface, using os/exec to run Git CLI commands
- **[gogit](https://github.com/go-xlan/gogit)** (this project) - Enhanced Git operations toolkit with go-git foundation, providing pure Go implementation

## Installation

```bash
go get github.com/go-xlan/gogit
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/go-xlan/gogit"
)

func main() {
    // Initialize Git client
    client, err := gogit.New("/path/to/your/repo")
    if err != nil {
        log.Fatal(err)
    }

    // Stage all changes
    err = client.AddAll()
    if err != nil {
        log.Fatal(err)
    }

    // Create commit info with fluent API
    commitInfo := gogit.NewCommitInfo("Initial commit").
        WithName("Your Name").
        WithMailbox("your.mailbox@example.com")

    // Commit changes
    hash, err := client.CommitAll(commitInfo)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Commit created: %s\n", hash)
}
```

### Advanced Features

```go
// Check repository status
status, err := client.Status()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Repository status: %+v\n", status)

// Amend last commit (with safety check)
amendConfig := &gogit.AmendConfig{
    CommitInfo: gogit.NewCommitInfo("Updated commit message").
        WithName("Updated Name").
        WithMailbox("updated.mailbox@example.com"),
    ForceAmend: false, // Prevents amending pushed commits
}

hash, err := client.AmendCommit(amendConfig)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Amended commit: %s\n", hash)

// Check if latest commit was pushed to remote
pushed, err := client.IsLatestCommitPushed()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Latest commit pushed: %t\n", pushed)
```

## API Reference

### Core Methods

- **`gogit.New(root string) (*Client, error)`**
  Creates a new Git client for the specified repo path with ignore file support

- **`client.AddAll() error`**
  Stages all changes including new files, modifications, and deletions

- **`client.Status() (git.Status, error)`**
  Returns current worktree status with comprehensive file change info

- **`client.CommitAll(info *CommitInfo) (string, error)`**
  Commits all staged changes with provided creator signature and message

- **`client.AmendCommit(cfg *AmendConfig) (string, error)`**
  Amends the last commit with safety checks for pushed commits

- **`client.IsLatestCommitPushed() (bool, error)`**
  Checks if current branch has been pushed to any configured remote

- **`client.IsLatestCommitPushedToRemote(name string) (bool, error)`**
  Checks push status against a specific remote repo

- **`client.GetCurrentBranch() (string, error)`**
  Returns the name of the current branch

- **`client.GetLatestCommit() (*object.Commit, error)`**
  Returns the latest commit object with message and author info

- **`client.HasChanges() (bool, error)`**
  Checks if the repo has uncommitted changes

- **`client.GetRemoteURL(name string) (string, error)`**
  Returns the URL for the specified remote

### Configuration Types

```go
// CommitInfo - Fluent commit configuration
type CommitInfo struct {
    Name    string // Creator name for Git commits
    Mailbox string // Creator mailbox for Git commits
    Message string // Commit message content
}

// AmendConfig - Amend operation configuration
type AmendConfig struct {
    CommitInfo *CommitInfo // New commit info for amend operation
    ForceAmend bool        // Allow amend even if commit was pushed
}
```

### Fluent API Examples

```go
// Create commit info with method chaining
commitInfo := gogit.NewCommitInfo("Feature implementation").
    WithName("Developer Name").
    WithMailbox("dev@company.com")

// Use default message generation if no message provided
commitInfo := gogit.NewCommitInfo("").
    WithName("Auto Account").
    WithMailbox("auto@example.com")
// Generates timestamp-based message: "[gogit](github.com/go-xlan/gogit) 2024-01-15 14:30:45"
```

## Safety Features

- **Push Detection**: Prevents amending commits that have been pushed to remote repos
- **Ignore File Support**: Respects .gitignore patterns during operations
- **Empty Commit Handling**: Returns empty string for no-change commits
- **Error Context**: Comprehensive error wrapping with context info
- **Hash Verification**: Validates commit integrity after operations

## Best Practices

```go
// Always check for errors
client, err := gogit.New("/path/to/repo")
if err != nil {
    return fmt.Errorf("failed to create client: %w", err)
}

// Use fluent API for clean configuration
info := gogit.NewCommitInfo("Fix critical bug").
    WithName("Bug Fixer").
    WithMailbox("fixer@company.com")

// Check push status before amending
if pushed, _ := client.IsLatestCommitPushed(); pushed {
    log.Println("Warning: Cannot amend pushed commit")
} else {
    // Safe to amend
    hash, err := client.AmendCommit(&gogit.AmendConfig{
        CommitInfo: info,
        ForceAmend: false,
    })
}
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/go-xlan/gogit.svg?variant=adaptive)](https://starchart.cc/go-xlan/gogit)
# gogitv5git
use `git add` `git commit` `git push` with "github.com/go-git/go-git/v5".

## README
[中文说明](README.zh.md)

## Installation

```bash
go get github.com/go-xlan/gogitv5git
```

## Usage

### Initializing the Git Client

First, create a new Git client instance by calling the `New` function with the repository's root directory:

```go
package main

import (
	"fmt"
	"log"

	"github.com/go-xlan/gogitv5git"
	"github.com/yyle88/done"
)

func main() {
	client := done.VPE(gogitv5git.New("/path/to/your/repository")).Nice()
	fmt.Println("Git client initialized!")
}
```

### Adding All Changes

To add all changes (including deletions) to the Git index (staging area), use the `AddAll` method:

```go
err := client.AddAll()
done.Done(err)
```

### Viewing Git Status

To view the current status of the working tree, use the `Status` method:

```go
status, err := client.Status()
done.Done(err)

fmt.Println("Git Status: ", status)
```

### Committing Changes

To commit all changes, use the `CmtAll` method. You need to provide a `CommitMessage` struct, which defines the commit message and signature.

```go
commitMessage := gogitv5git.CommitMessage{
	Name:    "Your Name",
	Emails:  "youremail@example.com",
	Message: "Your commit message",
}

commitHash, err := client.CmtAll(commitMessage)
done.Done(err)

fmt.Println("Commit successful! Commit hash: ", commitHash)
```

### Amending the Latest Commit

To amend the latest commit (e.g., to modify the commit message or add more changes), use the `CAmend` method:

```go
commitMessage := gogitv5git.CommitMessage{
	Message: "Amended commit message",
}

commitHash, err := client.CAmend(commitMessage)
done.Done(err)

fmt.Println("Amend successful! Commit hash: ", commitHash)
```

### Other Features

`gogitv5git` provides additional functionality such as retrieving commit hashes and logs. Feel free to explore the source code for more advanced features and extensions.

## Function Overview

- **`New(root string) (*Client, error)`**  
  Initializes and returns a new `Client` instance for interacting with the Git repository located at the specified path.

- **`AddAll() error`**  
  Adds all changes (including deletions) to the Git index (staging area).

- **`Status() (git.Status, error)`**  
  Returns the current status of the working tree.

- **`CmtAll(options CommitMessage) (string, error)`**  
  Commits all changes with the provided `CommitMessage` for the commit's author and message.

- **`CAmend(options CommitMessage) (string, error)`**  
  Amends the latest commit with the provided commit message or adds new changes. The commit is amended using the `--amend` flag.

## Contributing

Contributions are welcome! If you'd like to help improve this project, please feel free to:

- Open an issue for bug reports or feature requests
- Submit a pull request with your improvements

## License

MIT License - See the `LICENSE` file for more details.

## Thank you

Give me stars. Thank you!!!

[![see stars](https://starchart.cc/go-xlan/gogitv5git.svg?variant=adaptive)](https://starchart.cc/go-xlan/gogitv5git)

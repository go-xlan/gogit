# gogit
use `git add` `git commit` `git push` with "github.com/go-git/go-git/v5".

## README
[中文说明](README.zh.md)

## Installation

```bash
go get github.com/go-xlan/gogit
```

## Usage

### Initializing the Git Client

First, create a new Git client instance by calling the `New` function with the repository's root directory:

```go
package main

import (
	"fmt"
	"log"

	"github.com/go-xlan/gogit"
)

func main() {
	client, _ := gogit.New("/path/to/your/repository")
	fmt.Println("OK!")
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

To commit all changes, use the `CommitAll` method. You need to provide a `CommitInfo` struct, which defines the commit message and signature.

```go
commitInfo := gogit.CommitInfo{
	Name:    "Your Name",
	Eddress:  "youremail@example.com",
	Message: "Your commit message",
}

commitHash, err := client.CommitAll(commitInfo)
done.Done(err)

fmt.Println("Commit successful! Commit hash: ", commitHash)
```

### Amending the Latest Commit

To amend the latest commit (e.g., to modify the commit message or add more changes), use the `AmendCommit` method:

```go
amendConfig := gogit.AmendConfig{
	//message
}

commitHash, err := client.AmendCommit(amendConfig)
done.Done(err)

fmt.Println("Amend successful! Commit hash: ", commitHash)
```

### Other Features

`gogit` provides additional functionality such as retrieving commit hashes and logs. Feel free to explore the source code for more advanced features and extensions.

## Function Overview

- **`New(root string) (*Client, error)`**  
  Initializes and returns a new `Client` instance for interacting with the Git repository located at the specified path.

- **`AddAll() error`**  
  Adds all changes (including deletions) to the Git index (staging area).

- **`Status() (git.Status, error)`**  
  Returns the current status of the working tree.

- **`CommitAll(options CommitInfo) (string, error)`**  
  Commits all changes with the provided `CommitInfo` for the commit's author and message.

- **`AmendCommit(options AmendConfig) (string, error)`**  
  Amends the latest commit with the provided commit message or adds new changes. The commit is amended using the `--amend` flag.

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/go-xlan/gogit.svg?variant=adaptive)](https://starchart.cc/go-xlan/gogit)

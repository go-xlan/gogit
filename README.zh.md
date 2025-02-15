# gogit

`gogit` 是一个 Go 语言库，用于操作 Git 仓库。该库提供了一些常用的 Git 操作，如添加文件、查看状态、提交更改等。基于 `go-git` 库实现，提供了易于使用的 API 来简化 Git 操作。

## 说明
[ENGLISH-README](README.md)

## 安装

```bash
go get github.com/go-xlan/gogit
```

## 使用

### 初始化 Git 客户端

首先，您需要初始化一个 Git 仓库的客户端实例。你可以通过 `New` 函数创建一个新的客户端对象。

```go
package main

import (
	"fmt"
	"log"

	"github.com/go-xlan/gogit"
)

func main() {
	client := gogit.MustNew("/path/to/your/repository")
	fmt.Println("Git client initialized!")
}
```

### 添加所有更改

要将所有更改添加到 Git 索引（即暂存区），可以使用 `AddAll` 方法：

```go
err := client.AddAll()
done.Done(err)
```

### 查看 Git 状态

要查看当前工作区的状态，可以使用 `Status` 方法：

```go
status, err := client.Status()
done.Done(err)

fmt.Println("Git Status: ", status)
```

### 提交更改

要提交所有的更改，可以使用 `CommitAll` 方法，您需要传入一个 `CommitInfo` 结构体，该结构体用于定义提交信息和签名。

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

### 修改最新提交（Amend）

如果您想修改最新的提交信息，可以使用 `AmendCommit` 方法：

```go
amendConfig := gogit.AmendConfig{
	//message
}

commitHash, err := client.AmendCommit(amendConfig)
done.Done(err)

fmt.Println("Amend successful! Commit hash: ", commitHash)
```

### 其他功能

`gogit` 还提供了一些其他功能，例如获取提交哈希和日志等。你可以参考源码进行扩展或修改。

## 函数说明

- **`New(root string) (*Client, error)`**  
  初始化并返回一个新的 `Client` 实例，用于操作指定路径下的 Git 仓库。

- **`AddAll() error`**  
  添加所有更改（包括删除文件）到 Git 索引（暂存区）。

- **`Status() (git.Status, error)`**  
  获取当前工作区的状态。

- **`CommitAll(options CommitInfo) (string, error)`**  
  提交所有更改，并使用提供的 `CommitInfo` 生成提交信息。

- **`AmendCommit(options AmendConfig) (string, error)`**  
  修改最近的一次提交（使用 `--amend` 标志），并且支持为空的提交信息从最近的提交中获取。

## 贡献

欢迎贡献代码和提出问题，帮助这个项目变得更好！您可以通过以下方式参与：

- 提交 Issue 来报告问题
- 提交 Pull Request 进行代码改进

## 许可证

MIT License - 参阅 `LICENSE` 文件获取更多信息。

## 谢谢

帮我点个星星。谢谢!!!

[![starring](https://starchart.cc/go-xlan/gogit.svg?variant=adaptive)](https://starchart.cc/go-xlan/gogit)

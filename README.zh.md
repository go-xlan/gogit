[![GoDoc](https://pkg.go.dev/badge/github.com/go-xlan/gogit)](https://pkg.go.dev/github.com/go-xlan/gogit)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xlan/gogit)](https://goreportcard.com/report/github.com/go-xlan/gogit)

# gogit

增强的 Git 操作工具包，提供简化的仓库管理，具备全面的提交和远程同步功能。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **简化的 Git 操作**: 智能暂存、提交和状态检查，具备全面的 API
⚡ **智能提交管理**: 自动暂存与提交和修正支持，防止不安全操作
🔄 **远程推送检测**: 自动检查提交在多个远程仓库的推送状态
🌍 **跨平台支持**: 纯 Go 实现，无需 CLI 依赖，基于 go-git 基础
📋 **流畅的 API 设计**: 构建器模式，便于配置和方法链式调用

## 安装

```bash
go get github.com/go-xlan/gogit
```

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "log"

    "github.com/go-xlan/gogit"
)

func main() {
    // 初始化 Git 客户端
    client, err := gogit.New("/path/to/your/repo")
    if err != nil {
        log.Fatal(err)
    }

    // 暂存所有更改
    err = client.AddAll()
    if err != nil {
        log.Fatal(err)
    }

    // 使用流畅 API 创建提交信息
    commitInfo := gogit.NewCommitInfo("初始提交").
        WithName("您的姓名").
        WithMailbox("your.email@example.com")

    // 提交更改
    hash, err := client.CommitAll(commitInfo)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("提交创建成功: %s\n", hash)
}
```

### 高级功能

```go
// 检查仓库状态
status, err := client.Status()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("仓库状态: %+v\n", status)

// 修正最后一次提交（带安全检查）
amendConfig := &gogit.AmendConfig{
    CommitInfo: gogit.NewCommitInfo("更新的提交信息").
        WithName("更新的姓名").
        WithMailbox("updated.email@example.com"),
    ForceAmend: false, // 防止修正已推送的提交
}

hash, err := client.AmendCommit(amendConfig)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("修正提交成功: %s\n", hash)

// 检查最新提交是否已推送到远程
pushed, err := client.IsLatestCommitPushed()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("最新提交已推送: %t\n", pushed)
```

## API 参考

### 核心方法

- **`gogit.New(root string) (*Client, error)`**
  为指定的仓库路径创建新的 Git 客户端，支持忽略文件

- **`client.AddAll() error`**
  暂存所有更改，包括新文件、修改和删除

- **`client.Status() (git.Status, error)`**
  返回当前工作树状态，包含全面的文件更改信息

- **`client.CommitAll(info *CommitInfo) (string, error)`**
  使用提供的创建者签名和消息提交所有已暂存的更改

- **`client.AmendCommit(cfg *AmendConfig) (string, error)`**
  修正最后一次提交，对已推送的提交进行安全检查

- **`client.IsLatestCommitPushed() (bool, error)`**
  检查当前分支是否已推送到任何配置的远程仓库

- **`client.IsLatestCommitPushedToRemote(name string) (bool, error)`**
  检查针对特定远程仓库的推送状态

### 配置类型

```go
// CommitInfo - 流畅的提交配置
type CommitInfo struct {
    Name    string // 用于 Git 提交的创建者姓名
    Mailbox string // 用于 Git 提交的创建者邮箱
    Message string // 提交消息内容
}

// AmendConfig - 修正操作配置
type AmendConfig struct {
    CommitInfo *CommitInfo // 修正操作的新提交信息
    ForceAmend bool        // 即使提交已推送也允许修正
}
```

### 流畅 API 示例

```go
// 使用方法链式调用创建提交信息
commitInfo := gogit.NewCommitInfo("功能实现").
    WithName("开发者姓名").
    WithMailbox("dev@company.com")

// 如果没有提供消息，则使用默认消息生成
commitInfo := gogit.NewCommitInfo("").
    WithName("自动用户").
    WithMailbox("auto@example.com")
// 生成基于时间戳的消息: "[gogit](github.com/go-xlan/gogit) 2024-01-15 14:30:45"
```

## 安全特性

- **推送检测**: 防止修正已推送到远程仓库的提交
- **忽略文件支持**: 在操作期间遵守 .gitignore 模式
- **空提交处理**: 对于无更改的提交返回空字符串
- **错误上下文**: 全面的错误包装，包含上下文信息
- **哈希验证**: 在操作后验证提交完整性

## 最佳实践

```go
// 始终检查错误
client, err := gogit.New("/path/to/repo")
if err != nil {
    return fmt.Errorf("创建客户端失败: %w", err)
}

// 使用流畅 API 进行清晰配置
info := gogit.NewCommitInfo("修复严重错误").
    WithName("错误修复者").
    WithMailbox("fixer@company.com")

// 修正前检查推送状态
if pushed, _ := client.IsLatestCommitPushed(); pushed {
    log.Println("警告: 无法修正已推送的提交")
} else {
    // 安全修正
    hash, err := client.AmendCommit(&gogit.AmendConfig{
        CommitInfo: info,
        ForceAmend: false,
    })
}
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/go-xlan/gogit.svg?variant=adaptive)](https://starchart.cc/go-xlan/gogit)
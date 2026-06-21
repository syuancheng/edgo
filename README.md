# edgo：Go 学习笔记与可运行示例

`edgo` 是一个以可运行代码为核心的 Go 学习仓库。内容按“语言基础 → 并发与标准库 → 工程实践 → 算法与代码风格”组织；每个示例都应能独立运行，仓库整体应通过测试和静态检查。

## 快速开始

环境要求：Go 1.19 或更高版本。

```bash
make test
go run ./language/maps
go run ./language/concurrency/channel
go test -v ./algorithms/leetcode
```

## 目录结构

| 目录 | 内容 |
| --- | --- |
| [`language/`](./language/) | 变量、字符串、函数、方法、值语义、map、接口、错误、泛型、并发 |
| [`standard-library/`](./standard-library/) | `context`、JSON、字符串、时间、HTTP 等标准库示例 |
| [`examples/`](./examples/) | 异步任务、限流、panic 恢复、gzip、反射、接口嵌入等工程示例 |
| [`book-notes/`](./book-notes/) | 按章节保留的读书笔记和配套代码 |
| [`algorithms/`](./algorithms/) | 算法实现及表驱动测试 |
| [`style-guide/`](./style-guide/) | 可执行的 Go 代码风格示例 |
| [`docs/`](./docs/) | 学习路线、工具链和仓库约定 |

建议先阅读[学习路线](./docs/learning-roadmap.md)，再按主题运行示例。示例中的“错误写法”只放在注释或文档代码块中，确保默认构建始终可用。

## 常用命令

```bash
make fmt      # 格式化 Go 代码
make vet      # 静态检查
make test     # 运行所有测试
make check    # fmt 检查 + vet + test
```

## 维护原则

- 一个目录只包含一个可构建的包；独立程序使用独立目录。
- 文件和目录使用准确的英文命名，示例入口统一为 `main.go`。
- 知识结论必须配套最小示例或测试；危险示例不能让默认测试死锁或 panic。
- 新增代码后至少运行 `make check`；并发代码额外运行 `go test -race ./...`。

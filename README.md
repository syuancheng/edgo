# edgo — Go 学习空间

个人 Go 学习仓库：语法基础 → 标准库 → 工程模式 → 实战项目。

## 目录结构

| 目录 | 内容 | 运行方式 |
|------|------|---------|
| `grammar/` | 语法基础：函数、类型、引用、并发（goroutine / channel）、cgo | `go run ./grammar/<主题> [示例名]` |
| `book/` | 《Go 语言学习笔记》书章练习（ch5 数组/切片/结构体，ch7 接口） | `go run ./book/ch5 [示例名]` |
| `stdlib/` | 标准库练习：time、context、json、strings、encoding、gzip 解析 | 见各目录 |
| `patterns/` | **工程模式（工作重点）**：错误处理、并发模式、限流、异步包装、优雅退出、测试、Uber 风格指南 | 见下文 |
| `projects/` | 综合小项目：BFF 异步 Tab 处理 | `go run ./projects/bff_dd_tab_case` |
| `leetcode/` | LeetCode 题解（n2 / n4 / n11） | `go run ./leetcode n2` |
| `playground/` | 临时试验场 | `go run ./playground` |

## 统一入口分发（2026-10 整理引入）

同一主题的多个示例合并成一个可编译的包后，每个示例的 `main` 改成了与文件名同名的函数，
由各目录统一的 `main.go` 分发：

```bash
cd grammar/concurrency/channel
go run .            # 列出该目录全部示例
go run . test11_1   # 运行指定示例
```

参数即示例名（文件名去掉 `.go`）。

## patterns/ —— 工作中最重要的模式

| 模块 | 内容 | 推荐指数 |
|------|------|---------|
| [errorhandling](./patterns/errorhandling) | %w 包装、errors.Is/As、哨兵错误、自定义错误类型、panic/recover 边界 | ⭐⭐⭐⭐⭐ |
| [concurrencypattern](./patterns/concurrencypattern) | errgroup、worker pool、pipeline（fan-out/fan-in） | ⭐⭐⭐⭐⭐ |
| [gracefulshutdown](./patterns/gracefulshutdown) | 信号监听 + http.Server.Shutdown 优雅退出 | ⭐⭐⭐⭐⭐ |
| [testing](./patterns/testing) | 表驱动测试、子测试、benchmark、-race | ⭐⭐⭐⭐⭐ |
| [ratelimit](./patterns/ratelimit) | 令牌桶限流（x/time/rate、uber ratelimit）、channel 并发限制 | ⭐⭐⭐⭐ |
| [async](./patterns/async) | 基于 context 的异步函数包装器（手写实现） | ⭐⭐⭐ |
| [uber_style](./patterns/uber_style) | Uber Go 风格指南逐条示例（接口、mutex、切片拷贝、嵌入等） | ⭐⭐⭐⭐ |

每个模块都有独立的 README 讲解要点。

## 测试

```bash
go test ./...                                            # 全部测试
go test ./patterns/testing/counter/ -bench=. -benchmem   # 基准测试
go test ./grammar/data_type/map/benchmark/ -bench=.      # map 预分配性能对比
go test -race ./...                                      # 竞态检测
```

## 建议学习路线

1. `grammar/`（变量 → 函数 → 引用/值语义 → goroutine → channel）
2. `book/`（把切片、结构体、接口的底层概念补扎实）
3. `stdlib/`（context 是重点，贯穿后面所有模式）
4. `patterns/`（errorhandling → concurrencypattern → gracefulshutdown → testing）
5. `leetcode/` + `projects/`（边刷题边把模式用起来）

## 仓库约定

- 不再出现同目录多个 `main`：新示例一律写成具名函数，注册到该目录 `demos` map
- 所有代码可编译：`go build ./...` 必须通过；提交前跑 `gofmt -l .` 和 `go vet ./...`
- 修复历史代码时，用注释保留原写法并说明原因，**不直接删代码**
- 依赖统一 vendor（`go mod vendor`），网络受限时用 `GOPROXY=https://goproxy.cn,direct`

## 历史整理说明（2026-10-14）

- 全部旧文件经 `git mv` 迁移（保留 git 历史，可用 `git log --follow` 追溯），**未删除任何文件**
- `grammer/` → `grammar/`（修正拼写）；`ch5`、`ch7` → `book/`；`project/` 按内容拆入 `stdlib/`、`patterns/`、`projects/`；`src/` → `patterns/`
- 修复了历史遗留的编译错误（多 main 冲突、未使用变量/导入、nil channel、无限递归、整数除法丢精度等），关键修复处均有 `// 修复` 注释

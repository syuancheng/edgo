# edgo

`edgo` 是一个以可运行代码为核心的 Go 学习仓库，用于记录语言知识、标准库用法、工程模式和算法练习。

这里的内容不是零散代码备忘录。每个主题都尽量包含简明笔记、最小示例或测试，并保证整个仓库可以统一格式化、检查和构建。

## 项目内容

- Go 基础：变量、字符串、函数、方法、值语义、map、接口和泛型。
- 并发编程：goroutine、channel、同步、取消传播和数据竞争。
- 标准库：`context`、`encoding/json`、`net/http`、`strings` 和 `time`。
- 工程实践：异步任务、限流、错误处理、panic 恢复、反射和 gzip。
- 算法练习：带表驱动测试的 LeetCode 实现。
- 代码风格：以可运行示例解释常见 Go 工程约定。

## 目录结构

```text
edgo/
├── fundamentals/    # Go 语言核心知识
├── stdlib/          # Go 标准库专题
├── examples/        # 可独立运行的工程示例
├── algorithms/      # 算法实现与测试
├── chapter-notes/   # 按章节整理的读书笔记
├── style-guides/    # Go 代码风格示例
└── docs/            # 学习路线与工具链说明
```

| 目录 | 说明 |
| --- | --- |
| [`fundamentals/`](./fundamentals/) | 从变量、函数和方法逐步学习接口、泛型、测试与并发 |
| [`stdlib/`](./stdlib/) | 通过小程序理解常用标准库 API 和资源管理 |
| [`examples/`](./examples/) | 面向实际问题的完整示例，每个程序拥有独立入口 |
| [`algorithms/`](./algorithms/) | 使用普通 Go 包和测试记录算法解法及边界条件 |
| [`chapter-notes/`](./chapter-notes/) | 保留章节顺序的读书笔记与配套实验 |
| [`style-guides/`](./style-guides/) | 所有权边界、方法集、互斥锁和类型嵌入等风格示例 |
| [`docs/`](./docs/) | 推荐学习顺序和 Go 工具链速查 |

## 快速开始

环境要求：Go 1.19 或更高版本。cgo 示例还需要本机 C 编译器。

```bash
git clone https://github.com/syuancheng/edgo.git
cd edgo
make check
```

选择一个主题运行：

```bash
go run ./fundamentals/maps
go run ./fundamentals/concurrency/channel
go run ./stdlib/http
go run ./examples/rate-limiting/token-bucket
go test -v ./algorithms/leetcode
```

## 推荐学习顺序

1. 从 [`fundamentals/`](./fundamentals/) 掌握类型、值语义、函数和接口。
2. 学习 goroutine、channel，再进入 [`stdlib/context`](./stdlib/context/) 理解取消传播。
3. 运行 [`examples/`](./examples/) 中的工程示例，观察不同组件如何组合。
4. 用 [`algorithms/`](./algorithms/) 和测试练习边界条件与复杂度分析。
5. 阅读 [`style-guides/`](./style-guides/) 建立可维护代码的基本约定。

更详细的安排参见[学习路线](./docs/learning-roadmap.md)。

## 常用命令

```bash
make fmt      # 格式化全部 Go 代码
make vet      # 运行静态检查
make test     # 运行全部测试
make race     # 使用 race detector 运行测试
make check    # 检查格式，并执行 vet 和 test
```

## 内容约定

- 一个目录只包含一个可构建的包，独立程序使用独立目录。
- 示例入口统一命名为 `main.go`，测试文件使用 `_test.go`。
- 知识结论应配套代码、测试或可重复的运行结果。
- 错误写法只放在注释或 Markdown 代码块中，不能破坏默认构建。
- 并发示例不使用 `time.Sleep` 代替正确的同步机制。
- 新增内容后至少运行 `make check`；修改并发代码后运行 `make race`。

## 文档入口

- [学习路线](./docs/learning-roadmap.md)
- [Go 工具链速查](./docs/tooling.md)
- [章节笔记](./chapter-notes/)
- [工程示例索引](./examples/)

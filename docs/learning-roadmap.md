# Go 学习路线

本仓库采用“先理解值和类型，再学习并发与工程实践”的顺序。

## 1. 语言基础

1. [变量与零值](../fundamentals/variables/)
2. [字符串](../fundamentals/strings/)
3. [数组、切片和结构体](../chapter-notes/chapter-05-data/)
4. [map](../fundamentals/maps/)
5. [函数与闭包](../fundamentals/functions/)
6. [方法和方法集](../fundamentals/methods/)
7. [值语义与引用共享](../fundamentals/value-semantics/)
8. [接口](../chapter-notes/chapter-07-interfaces/)
9. [错误处理](../fundamentals/errors/)
10. [泛型](../fundamentals/generics/)

学习重点不是记住语法，而是能回答：赋值时复制了什么、哪些数据仍然共享、方法集如何决定接口实现、错误由谁补充上下文并处理。

## 2. 并发

依次学习 goroutine、channel、`select`、取消传播和同步原语：

- [goroutine](../fundamentals/concurrency/goroutine/)
- [channel](../fundamentals/concurrency/channel/)
- [`context`](../stdlib/context/)

不要用 `time.Sleep` 作为正确性同步手段。优先使用 `WaitGroup`、channel 或 `context`，并用 `go test -race ./...` 检查数据竞争。

## 3. 标准库与工程能力

- 数据格式：`encoding/json`、`compress/gzip`
- I/O：`io.Reader`、`io.Writer`、资源关闭
- 网络：`net/http`、超时、状态码和响应体处理
- 质量：表驱动测试、基准测试、`gofmt`、`go vet`、race detector

## 4. 推荐练习方式

每学习一个主题，完成以下闭环：

1. 运行现有示例并预测输出。
2. 修改输入，解释值、地址或 goroutine 生命周期的变化。
3. 把结论写成一个可重复的测试。
4. 运行 `make check`，并为并发代码运行 `make race`。

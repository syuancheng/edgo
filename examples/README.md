# 工程示例

这些示例聚焦可复用的工程模式，而不是单一语法点。每个程序都可通过 `go run ./examples/<name>` 或其子目录运行。

- [`async`](./async/)：支持取消并缓存结果的异步任务
- [`rate-limit`](./rate-limit/)：并发限制和速率限制
- [`gzip`](./gzip/)：流式解压并正确关闭资源
- [`panic-recovery`](./panic-recovery/)：在发生 panic 的 goroutine 内恢复
- [`reflection`](./reflection/)：带参数校验的反射字段复制
- [`interface-embedding`](./interface-embedding/)：接口与结构体嵌入
- [`http-requests`](./http-requests/)：编辑器可执行的 HTTP 请求草稿

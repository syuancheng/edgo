# 限流

- [`concurrency-limit`](./concurrency-limit/) 用有界 semaphore 限制同时执行的任务数量。
- [`token-bucket`](./token-bucket/) 使用 `golang.org/x/time/rate` 控制平均速率和突发量。
- [`leaky-bucket`](./leaky-bucket/) 使用 `go.uber.org/ratelimit` 平滑调用间隔。

并发限制和速率限制解决不同问题，实际服务常同时使用。示例中的非阻塞 semaphore 会拒绝超额任务；需要排队时可改为阻塞获取，并结合 context 控制等待期限。

# Context

`context.Context` 在调用链中传播截止时间、取消信号和请求级数据。

- `context.Background()` 通常用于进程或请求的根。
- `WithCancel`、`WithTimeout` 和 `WithDeadline` 返回派生 context 与取消函数。
- 获得取消函数后应尽快 `defer cancel()`，释放定时器和关联资源。
- 长任务应周期性检查 `ctx.Done()`，并返回 `ctx.Err()` 或带上下文的包装错误。
- Context 应作为函数第一个参数传递，不存入结构体，也不要传 `nil`。
- `WithValue` 只用于跨 API 边界的请求级元数据，不用于可选参数。

取消是协作式的：它不会强制杀死 goroutine。任务代码必须主动观察取消信号并退出。

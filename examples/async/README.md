# 异步任务

`Task[T]` 立即在 goroutine 中执行函数，并通过关闭 `done` channel 发布结果。多个等待者可安全读取同一个完成结果；每次等待还可使用独立的 `context` 提前退出。

任务函数本身必须观察创建时传入的 context，取消等待者不会强制终止正在执行的函数。

```bash
go run ./examples/async/cmd
go test ./examples/async
```

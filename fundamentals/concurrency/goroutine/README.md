# Goroutine

`go f(args...)` 会先在当前 goroutine 中计算函数值和参数，再启动新的 goroutine 执行调用。调度顺序没有保证，`main` 返回时进程会直接结束，不会自动等待其他 goroutine。

## 生命周期与同步

- 使用 `sync.WaitGroup` 等待一组任务。
- `Add` 必须在启动 goroutine 前完成，避免 `Wait` 提前返回。
- 使用 channel 传递所有权或结果，使用 `context` 传播取消。
- 不要用 `time.Sleep` 猜测任务是否已经完成。

## 数据竞争

多个 goroutine 并发访问同一变量，且至少一个访问是写操作时，需要 channel、互斥锁或原子操作建立同步。程序“看起来运行正常”不代表没有竞争。

```bash
go run ./fundamentals/concurrency/goroutine
go test -race ./...
```

# 并发工程模式（Concurrency Patterns）

面试和实际项目里最高频的三个并发模式：**errgroup、worker pool、pipeline（fan-out/fan-in）**。

## 运行

```bash
go run ./patterns/concurrencypattern errgroupDemo
go run ./patterns/concurrencypattern workerPoolDemo
go run ./patterns/concurrencypattern pipelineDemo
```

## 三个模式分别解决什么问题

### 1. errgroup —— "一群任务，一个失败就整体失败"

并行扇出多个下游调用（RPC / DB / 缓存）时，手工管理 sync.WaitGroup + 错误收集很啰嗦。
errgroup 一行 `g.Go(fn)` 搞定，`g.Wait()` 返回第一个触发的错误。

- `g.SetLimit(n)`：限制最大并发，本质是信号量，防打爆下游
- 需要 context 联动取消（一路失败其余全部取消）用 `errgroup.WithContext(ctx)`，
  失败时 ctx 会被 cancel，其他子任务应在 select 里监听 `ctx.Done()` 提前退出

### 2. worker pool —— "并发度钉死，任务排队"

任务量不可控时，"一个任务一个 goroutine"会失控（连接数、内存）。
固定 N 个 worker 消费任务 channel，并发度恒为 N。

关键细节：
- worker 用 `for t := range tasks` —— 通道 close 后自动退出，不需要退出标志位
- 先 `close(tasks)` → `wg.Wait()` → 再 `close(results)`，关闭顺序错了会死锁或丢数据

### 3. pipeline —— "流水线阶段化，数据流动起来"

每个阶段一个（或多个）goroutine，阶段间只靠 channel 通信：
- 每个阶段 defer close 自己的输出通道 —— 结束信号沿着流水线自然传播
- fan-out：一个输入通道接多个处理阶段，自动分流
- fan-in：merge 用额外 goroutine 在 wg.Wait() 后关闭输出 —— 收尾关闭是模板代码，值得背下来

## 配套阅读

- `patterns/async` —— 你手写的异步包装器，本质是"单任务 + context 超时"模式
- `patterns/ratelimit` —— 限流：控制"速率"，worker pool 控制"并发度"，两者互补

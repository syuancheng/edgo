# 并发

Go 的并发模型由 goroutine、channel 和同步原语共同组成。Channel 适合通信和所有权转移；`sync.Mutex` 适合保护共享状态；`context.Context` 适合跨调用边界传播取消和截止时间。

同步操作不仅决定执行顺序，也建立内存可见性关系。普通变量上的并发读写不能依靠调度概率或 `Sleep` 保证安全。

- [Goroutine](./goroutine/)
- [Channel](./channel/)
- [`context` 示例](../../stdlib/context/)

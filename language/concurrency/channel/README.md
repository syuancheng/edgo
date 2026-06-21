# Channel

Channel 用于在 goroutine 之间传递值并建立同步关系。发送和接收都会复制 channel 元素；如果元素是指针、slice 或 map，复制后仍可能共享底层数据。

## 无缓冲与有缓冲

- 无缓冲 channel：一次发送必须与一次接收会合。
- 有缓冲 channel：缓冲区未满时发送可继续，非空时接收可继续。
- 缓冲容量不是 channel 类型的一部分。

`len(ch)` 和 `cap(ch)` 只适合观察或调试，不能用于并发正确性判断，因为返回后状态可能立即变化。

## 关闭规则

- 通常由发送方关闭 channel，表示“不会再发送”。
- 关闭后仍可读取缓冲数据；随后接收会立即得到元素零值和 `ok == false`。
- 向已关闭 channel 发送或重复关闭会 panic。
- 从 `nil` channel 收发会永久阻塞，关闭 `nil` channel 会 panic。

接收方不需要为了释放资源而关闭 channel。垃圾回收器会回收不再引用的 channel。

## `select`

`select` 在多个已就绪分支中选择一个执行。`default` 会让操作变为非阻塞，但在循环里使用可能造成忙等。已关闭 channel 的接收分支永远就绪；处理多个输入时，可把已关闭的 channel 变量设为 `nil` 来禁用对应分支。

运行示例：

```bash
go run ./language/concurrency/channel
```

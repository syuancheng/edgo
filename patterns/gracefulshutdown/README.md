# 优雅退出（Graceful Shutdown）

服务发布/重启时不停机丢请求的标准做法。

## 运行

```bash
go run ./patterns/gracefulshutdown
# 另开终端
curl localhost:8080/slow    # 3 秒慢接口
# 回到服务终端按 Ctrl+C，观察：收到信号 -> slow 请求仍正常返回 -> 进程才退出
```

## 核心要点

1. **信号**：`SIGINT`（Ctrl+C）/ `SIGTERM`（kill、k8s 滚动发布）。`signal.Notify` 的 channel 缓冲给 1。
2. **`srv.Shutdown(ctx)`**：停止接收新连接，等存量请求处理完才返回。`ListenAndServe` 随之返回 `http.ErrServerClosed` —— 这是正常退出，不要当成错误打日志报警。
3. **超时兜底**：用 `context.WithTimeout` 限制最长等待时间，防止某个挂死的请求卡住整个发布。
4. **不止 HTTP**：数据库连接池、Kafka 消费者、后台任务，都应在退出流程里依次 Close。
5. **k8s 环境**：配合 `terminationGracePeriodSeconds`，让它略大于应用内的等待上限。

## 常见错误写法

- 直接 `os.Exit(0)`：存量请求瞬间断掉，用户看到 connection reset
- 在信号处理里只 `srv.Close()`：Close 会立刻中断所有活跃连接，和优雅退出是两回事

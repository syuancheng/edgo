# 错误处理（Error Handling）

工作中 Go 错误处理的核心三板斧：**包装（%w）、判断（errors.Is）、提取（errors.As）**。

## 运行

```bash
go run ./patterns/errorhandling panicDemo    # goroutine panic/recover
go run ./patterns/errorhandling sentinelDemo # 哨兵错误
go run ./patterns/errorhandling wrapDemo     # 包装 + Is/As
```

## 核心要点

### 1. 永远不要用字符串比较错误

```go
// ❌ 错误
if err.Error() == "record not found" { ... }
// ✅ 正确
if errors.Is(err, ErrNotFound) { ... }
```

### 2. %w vs %v（最常见的坑）

- `fmt.Errorf("...: %w", err)` —— 包装错误，保留链条，`errors.Is/As` 可以穿透查找
- `fmt.Errorf("...: %v", err)` —— 只拼接文本，链条丢失，上层再也判断不了错误类型

### 3. 三种错误定义方式的选择

| 方式 | 适用场景 | 判断方法 |
|------|---------|---------|
| 哨兵错误 `var ErrXxx = errors.New(...)` | 错误种类少、只需判断"是哪一种" | `errors.Is` |
| 自定义类型 `type XxxError struct` | 需要携带结构化信息（字段、code） | `errors.As` |
| 错误码枚举 | 对外 API 需要稳定错误码 | 配合 switch |

### 4. 包装时补充上下文

每一层返回错误时都加上当前层的信息（函数名、参数），
最终日志里的错误链会类似：`handler: validateAge: validation failed: field=age reason=不能为负数`，
能直接定位问题发生的位置和路径。

### 5. panic 的边界

- panic 只用于**不可恢复**的程序性错误（数组越界、断言失败、初始化失败）
- 业务错误一律用 error 返回值，不要用 panic 代替
- goroutine 里的 panic 不会传播到父 goroutine，必须在**同一个 goroutine 内** defer recover（见 panicDemo）
- HTTP 框架的 recovery 中间件本质上就是每个请求的 goroutine 里包了一层 recover

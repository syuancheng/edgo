# 测试（Testing）

table-driven test 和 benchmark 是 Go 工程化的两块基石。

## 运行

```bash
# 跑测试（-v 显示每个子用例）
go test ./patterns/testing/counter/ -v

# 跑基准测试（-benchmem 显示每次操作内存分配）
go test ./patterns/testing/counter/ -bench=. -benchmem

# 已有示例：map 预分配 vs 动态扩容的 benchmark 对比
go test ./grammar/data_type/map/benchmark/ -bench=. -benchmem

# 覆盖率
go test ./patterns/testing/counter/ -cover
```

## 核心要点

### 1. 表驱动测试（table-driven test）

用例集中在一张 struct 切片里，**加用例只加一行**。
每个用例有独立名字，失败输出直接显示 `TestInc/负步长也合法`。

常用断言姿势：

```go
if got := c.Total(); got != tt.want {
    t.Errorf("Total() = %d, want %d", got, tt.want)  // 带 got/want
}
```

- 失败信息一定带上 got / want，别只写一句话
- 需要判断"应该报错"的用例：`wantErr bool` 字段 + `err != nil` 检查
- 工作中一般再配合断言库（testify 的 assert/require）减少样板代码

### 2. 基准测试（benchmark）

- 函数签名 `func BenchmarkXxx(b *testing.B)`，核心循环跑 `b.N` 次
- **一定要消费结果**，否则编译器会把无副作用的循环优化掉，测了个寂寞
- `-benchmem` 看每次操作分配了几次堆内存 —— 性能优化先看分配，再看耗时

### 3. 其他高频工具

| 工具 | 用途 | 命令 |
|------|------|------|
| 子测试 | 用例分组、单独运行 | `go test -run TestInc/用例名` |
| Example 测试 | 文档 + 可执行示例 | `func ExampleXxx()` 输出注释校验 |
| `-race` | 竞态检测（CI 必开） | `go test -race ./...` |
| `-cover` | 覆盖率 | `go test -cover ./...` |
| `t.Parallel()` | 用例并行执行 | 放在用例函数第一行 |

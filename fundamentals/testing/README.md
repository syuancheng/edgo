# 测试

Go 测试文件以 `_test.go` 结尾，测试函数使用 `TestXxx` 命名。表驱动测试适合用统一逻辑覆盖多个输入、边界和错误场景。

```bash
go test -v ./fundamentals/testing
go test -cover ./fundamentals/testing
```

测试应验证公开行为，而不是内部实现。失败信息要包含输入、实际值和期望值。可重复的测试不能依赖 map 遍历顺序、真实网络或固定 `Sleep`。

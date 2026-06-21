# Go 工具链速查

## 模块与依赖

```bash
go mod tidy        # 清理并补齐依赖声明
go mod download    # 下载依赖
go list -m all     # 查看依赖图
```

库代码应使用完整模块导入路径。应用入口放在独立目录中，避免同一包出现多个 `main` 函数。

## 构建质量

```bash
gofmt -w .
go vet ./...
go test ./...
go test -race ./...
go test -cover ./...
go test -bench=. ./language/maps/benchmark
```

`gofmt` 负责统一格式，`go vet` 检查常见可疑代码，测试验证行为，race detector 检查运行时数据竞争。四者职责不同，不能互相替代。

## 调试与分析

```bash
go test -run TestName -v ./path/to/package
go test -bench=. -benchmem ./path/to/package
go build -gcflags='-m=2' ./path/to/package
go tool pprof cpu.out
```

性能优化前先建立基准。逃逸分析结果用于理解分配，不应机械地追求“全部在栈上”。

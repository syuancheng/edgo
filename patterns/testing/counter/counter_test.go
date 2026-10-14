package counter

import "testing"

// TestInc table-driven test（表驱动测试）：Go 社区标准测试写法。
// 好处：加用例只加一行；每个用例独立命名，失败时输出一眼定位；
// t.Run 让你可以用 go test -run TestInc/负步长 单独跑某个用例。
func TestInc(t *testing.T) {
	tests := []struct {
		name     string // 用例名：失败时显示 TestInc/xxx
		step     int    // 输入：步长
		incTimes int    // 输入：自增次数
		want     int    // 期望结果
	}{
		{"步长1自增3次", 1, 3, 3},
		{"步长5自增2次", 5, 2, 10},
		{"自增0次", 1, 0, 0},
		{"负步长也合法", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.step)
			for i := 0; i < tt.incTimes; i++ {
				c.Inc()
			}
			if got := c.Total(); got != tt.want {
				// 失败信息带上 got/want，不要只写一句话
				t.Errorf("Total() = %d, want %d", got, tt.want)
			}
		})
	}
}

// BenchmarkInc 基准测试：go test -bench=. -benchmem
// 输出形如 BenchmarkInc-8   1000000000   0.5 ns/op   0 B/op   0 allocs/op
// -benchmem 可以看每次操作的堆分配，性能优化的第一手数据。
func BenchmarkInc(b *testing.B) {
	c := New(1)
	// b.N 由框架控制，会自动调整到运行时间足够长为止
	for i := 0; i < b.N; i++ {
		c.Inc()
	}
	// 一定要消费结果，防止编译器把没有副作用的循环优化掉
	_ = c.Total()
}

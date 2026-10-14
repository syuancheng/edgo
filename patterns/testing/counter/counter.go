// Package counter 是 testing 模块的被测对象：
// 一个带步长的计数器，用来演示 table-driven test 和 benchmark 的写法。
package counter

// Counter 带固定步长的计数器
type Counter struct {
	total int
	step  int
}

// New 创建计数器，step 为每次自增的步长
func New(step int) *Counter {
	return &Counter{step: step}
}

// Inc 自增一次
func (c *Counter) Inc() {
	c.total += c.step
}

// Total 返回当前累计值
func (c *Counter) Total() int {
	return c.total
}

## Goroutine

#### basic info
创建一个并发任务单元（保存函数指针，调用参数，分配栈内存空间， 与defer一样，立即**计算并复制执行参数**），实际上会将其加入系统队列。
- 当前流程不会阻塞
- 不会等待该任务启动
- 不保证并发任务的执行次序

#### Test1
与defer一样，立即**计算并复制执行参数**

#### Test2
chan的初步使用
一个main任务，等待一个go并发任务

#### Test3
WaitGroup
一个main任务，等待一组go并发任务

#### Test4
错误用例， 直接exit
add未执行，wait已经退出


#### Test6
当main 方法退出之后， goroutine也会终止执行

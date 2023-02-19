### Channel
#### basic
同步模式（无buffer）： 发送和接受双方配对，如果配对失败，置入等待队列，直到另一方出现后才被唤醒。

异步模式（有buffer）： 抢夺buffer，发送方要有空buffer可供写入，接收方则要有buffer data可读，不符时同样加入等待队列，直到有data或位置时被唤醒。

及时用close func关闭通道引发结束通知。

#### 阻塞：
- no buffer chan
- buffer full or empty
- 读写nil chan 永久阻塞，不是写入或者消费了数据就能解除

#### panic：
- send data to closed chan(receive data from closed chan, return buffer data or zero value)
- close nil chan
- close closed chan

#### close chan
- ok-idom, ok=false
- for-range exit for-loop
- can receive data from closed chan

#### Test
ok-idom, 只有chan closed，ok=false,value=zero value
没有收到消息只会阻塞， ok不为任何值。

#### Test1
no buffer chan会阻塞
close(done) 会发送消息


#### Test2
buffer chan basic use

#### Test3
buffer size是内部属性，不属于类型组成部分。
通道变量本身就是指针

#### Test4
内置函数len和cap
cap返回buffer size，len返回当前buffer数量
对于no buffer chan，len和cap都是0。
可以据此来判断是同步还是异步通道。

#### Test5
ok-idom and close
收发test

#### Test6
for-range 接收
for-range exit for-loop, if close chan

#### Test7
WaitGroup close
收发测试

#### Test8
ok-idom 只有在close的时候会返回false  测试在buffer chan

#### Test9
单向chan，WaitGroup，close example

#### Test10
单向channel不可以逆向操作
close不能用于接收channel
不能把单向通道重新转回去


#### Test11
select 语句随机选择一个可用通道做收发操作。
即使channel closed， select依旧可以收到零值和ok=false

#### Test12
nil channel, select 会阻塞，这样对应的case就不会被执行了

#### Test13
select对同一个通道， 也会随机选择执行的
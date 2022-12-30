### 方法
#### 基本概念
- 对象实例绑定
- 面向对象
- 有receiver（可只保留类型）
    - 值： copy对象
    - 指针： 不copy，引用
    - 多级pointer不支持
- 不支持重载
- 实例value或pointer可以调用全部的方法，编译器会自动转换

#### T or *T
- 要改实例状态：*T
- 大对象： *T
— include Mutex： *T
- 无需更改状态的小对象： T
- 引用类型，字符串，函数等指针包装对象： T
- 无法确定： *T

#### 方法集: 
- 类型 T 方法集，包含全部 receiver T 方法。
- 类型 *T 方法集，包含全部 receiver T + *T 方法。
- 匿名嵌入S，T方法集包含receiver S 方法。
- 匿名嵌入*S，T方法集包含receiver S + *S 方法。
- 匿名嵌入S或*S，*T方法集包含receiver S + *S 方法。

```go
type S struct {}

type T struct {
    S
}
```


Ref： https://mp.weixin.qq.com/s/ESw8td1ipqgqvfJ0c7A4cg
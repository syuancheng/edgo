# 函数
### basic info
- 参数：函数的局部变量
- 签名=参数+返回值列表

### Test1
golang中参数是拷贝传递的。

### Test2
函数中的变参(一个或者多个同类型数据，且放在尾部)，就是一个切片。

### Test3
- 第一类对象，可用作函数参数或返回值。
- 从函数中返回局部变量指针是安全的，逃逸分析会决定是否在堆上分配内存。

### Test4
defer func() 注册，参数值被复制并缓存起来。
### basic info
key 必须是支持相等运算符的数据类型， 如： 数字，指针，string， 数组，结构体

#### test1:
map的初始化
make 可以提前预留足够的空间， 减少扩容。
匿名初始化
var : map 是nil，可读不可写

#### test2
如果不存在ok=false
并且删除不存在的key也不会panic

#### test3
对map进行遍历， 每次的顺序都是不一样的
并且遍历期间可以对其进行删除或者新增

#### test4
map中的value不可以直接修改 如: m2[0].x ++. 正确做法是返回整个value，待修改后再设置map键值
或者使用指针类型

#### test5
不是线程安全的

#### Benchmark_test
make cap 

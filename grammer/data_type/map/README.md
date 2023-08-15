# MAP
## 基本信息
### key
- key 必须是支持相等运算符的数据类型，如: 数字，指针，string， 数组，结构体

### 初始化
- make可以提前预留足够的空间， 减少扩容. `m := make(map[string]int, 5)`
- 匿名初始化
```
m2 := map[int]struct {
    x int
} {
    1: {x: 100},
    2: {x: 200},
}
```
- var初始化map 是nil，可读不可写 `var m3 map[int]string`  m3是nil

#### test2
如果不存在ok=false
并且删除不存在的key也不会panic

```go
var m5 map[int]string

fmt.Println(m5 == nil) //m5 is nil
//m5[1] = "" //panic
delete(m5, 1) //
val, ok := m5[1]
fmt.Println(val, ok) //"" false
```

#### test3
对map进行遍历， 每次的顺序都是不一样的
并且遍历期间可以对其进行删除或者新增

#### test4
map中的value不可以直接修改 如: m2[0].x ++. 正确做法是返回整个value，待修改后再设置map键值
或者使用指针类型
见： uber_style/interface/style2.go

#### test5
不是线程安全的

#### Benchmark_test
make cap 

# 第 5 章：数据

## 数组

数组长度是类型的一部分，`[2]int` 与 `[3]int` 是不同类型。数组是值类型，赋值和传参会复制全部元素；元素类型可比较时，数组也可比较。

- [声明与初始化](./arrays/basics/)
- [数组指针和指针数组](./arrays/pointers/)
- [复制语义](./arrays/copy/)

## Slice

Slice 是包含底层数组指针、长度和容量的描述符。复制 slice 只复制描述符，多个 slice 可能共享底层数组。

```go
var nilSlice []int       // nil，len=0，cap=0
empty := []int{}         // 非 nil，len=0，cap=0
values := make([]int, 2, 8)
```

- 切片表达式范围左闭右开。
- 普通切片 `a[low:high]` 要求索引不超过可用容量；三索引形式 `a[low:high:max]` 可限制新 slice 的容量。
- `append` 必须接收返回值。容量不足时会分配新数组，但扩容倍率是实现细节，不能假设恒为 2 倍。
- Slice 只能与 `nil` 比较；比较内容可使用循环或标准库中的相应函数。

示例：[声明](./slices/declaration/)、[nil slice](./slices/nil-slice/)、[指针](./slices/pointers/)、[`append`](./slices/append/)、[三索引表达式](./slices/full-expression/)。

## 结构体

结构体把不同类型的字段组合成一个值。推荐使用命名字段初始化，避免字段顺序变化影响调用方。结构体复制时每个字段按自身类型复制，因此 slice、map、指针等字段可能继续共享底层数据。

- [复制](./structs/copy/)
- [比较](./structs/comparison/)
- [指针](./structs/pointers/)
- [空结构体](./structs/empty/)
- [嵌入字段](./structs/embedding/)
- [结构体标签](./structs/tags/)

`struct{}` 的大小为 0，常用于只表达信号的 channel 或 set 的 map 值。嵌入会提升字段和方法，但不是继承；发生同名冲突时必须显式选择路径。

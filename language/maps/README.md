# Map

Map 是引用底层哈希表的描述符。赋值和传参会复制描述符，但两个 map 值通常仍操作同一组数据。

## 关键规则

- key 必须可比较，例如布尔值、数字、字符串、指针、channel、数组，以及字段都可比较的结构体。
- slice、map 和 function 不可作为 key。
- `var m map[K]V` 的零值是 `nil`：可以读取、查询和删除，写入会 panic。
- `make(map[K]V, n)` 的 `n` 是容量提示，不是固定容量，`len` 仍从 0 开始。
- map 元素不可寻址；修改结构体元素时要“取出—修改—写回”，或存储结构体指针。
- 遍历顺序未定义。需要稳定输出时先收集并排序 key。
- 普通 map 不支持并发读写。使用互斥锁、单一所有者 goroutine，或在合适场景使用 `sync.Map`。

运行：

```bash
go run ./language/maps
go test -bench=. -benchmem ./language/maps/benchmark
```

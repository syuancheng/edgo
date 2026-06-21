# Go 风格示例

本目录参考 Uber Go Style Guide，并将规则整理为独立可运行示例。

- [`interface-values`](./interface-values/)：接口按值传递，动态值可为指针。
- [`method-sets`](./method-sets/)：值与指针接收者的方法集。
- [`mutex`](./mutex/)：互斥锁使用零值并作为非导出字段保存。
- [`copying`](./copying/)：在所有权边界复制 slice 和 map。
- [`time-and-enums`](./time-and-enums/)：使用 `time.Time`、`time.Duration` 和明确的枚举零值。
- [`embedding`](./embedding/)：只在外层 API 应公开内嵌方法时嵌入类型。
- [`argument-names`](./argument-names/)：布尔参数在调用处缺乏语义时应改为配置类型或添加参数注释。

互斥锁保护的是不变量，而不只是某个字段。不要复制已经使用过的锁，也不要把内部 map 或 slice 直接返回给调用方后误以为仍受锁保护。

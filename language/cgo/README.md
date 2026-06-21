# cgo

cgo 允许 Go 调用 C 代码，但会增加构建、交叉编译和运行时边界成本。仅在必须复用 C 库或系统接口时使用。

`C.CString` 在 C 堆上分配内存，Go 垃圾回收器不会释放它，必须调用 `C.free`。跨边界传递 Go 指针还受 cgo 指针规则限制；不要让 C 长期保存指向 Go 内存的指针。

运行需要本机 C 编译器和 `CGO_ENABLED=1`：

```bash
go run ./language/cgo
```

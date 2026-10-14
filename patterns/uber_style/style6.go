package main

import (
	"bytes"
	"io"
	"sync"
)

// not recommend
type A struct {
	sync.Mutex
}

type Book struct {
	io.ReadWriter //零值是nil
}

// recommend
type countingWriteCloser struct {
	io.WriteCloser
	count int
}

func (wc *countingWriteCloser) Write(p []byte) (int, error) {
	wc.count += len(p)
	// 修复：原代码 return wc.Write(p) 是无限递归（Write 调用自己），
	// 正确写法是委托给嵌入的 WriteCloser：
	return wc.WriteCloser.Write(p)
}

type BookPro struct {
	bytes.Buffer
}

// demoNilInterfacePanic 演示 Book 零值调用嵌入接口方法的后果：panic。
// 用 recover 捕获，让 style6 可以完整跑完。
func demoNilInterfacePanic() {
	defer func() {
		if r := recover(); r != nil {
			println("Book 零值调用 nil 接口方法 -> panic:", r)
		}
	}()

	var b Book
	b.Read([]byte{}) // will panic：Book 零值中嵌入的 io.ReadWriter 是 nil
	// b.String()     // 原代码此行无法编译：Book 与 io.ReadWriter 都没有 String 方法
	b.Write([]byte{}) // 同样 will panic（执行不到这里）
}

func style6() {

	demoNilInterfacePanic()

	var bp BookPro
	bp.Read([]byte{})
	_ = bp.String()    // 修复（vet）：调用结果需消费
	bp.Write([]byte{}) // BookPro 嵌入的是 bytes.Buffer（零值可用），不会 panic
	// fmt.Println(test == nil)

}

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
	// return wc.WriteCloser.Write(p)
	return wc.Write(p)
}

type BookPro struct {
	bytes.Buffer
}

func main() {

	var b Book
	b.Read([]byte{})// will panic
	b.String()
	b.Write([]byte{})

	var bp BookPro
	bp.Read([]byte{})
	bp.String()
	bp.Write([]byte{})// will
	// fmt.Println(test == nil)

}

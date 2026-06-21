package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriteCloser struct {
	writer io.WriteCloser
	count  int
}

func (w *countingWriteCloser) Write(data []byte) (int, error) {
	written, err := w.writer.Write(data)
	w.count += written
	return written, err
}

func (w *countingWriteCloser) Close() error {
	return w.writer.Close()
}

type bufferCloser struct {
	bytes.Buffer
}

func (*bufferCloser) Close() error { return nil }

func main() {
	underlying := &bufferCloser{}
	writer := &countingWriteCloser{writer: underlying}
	if _, err := writer.Write([]byte("Go")); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
	fmt.Println(underlying.String(), writer.count)
}

package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"io"
)

//go:embed sg.json.gz
var compressedData []byte

func main() {
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

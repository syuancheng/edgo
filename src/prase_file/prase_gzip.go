package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	filePath = "/Users/siyuan.cheng/workspace/private/edgo/src/prase_file/sg.json.gz"
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	rawData, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(rawData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("value:%v", string(bytes))
}

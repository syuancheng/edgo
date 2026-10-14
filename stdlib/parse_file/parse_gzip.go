package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

// 修复：原来是写死的绝对路径（/Users/siyuan.cheng/...），换机器就失效。
// 现在优先取命令行参数，默认尝试两个相对路径（仓库根目录 / 本目录）。
// 用法: go run ./stdlib/parse_file [gz文件路径]
func resolveFilePath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	candidates := []string{
		"stdlib/parse_file/sg.json.gz", // 在仓库根目录运行时
		"sg.json.gz",                   // 在本目录运行时
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	log.Fatal("找不到 gz 文件，请把文件路径作为第一个参数传入")
	return ""
}

func main() {
	filePath := resolveFilePath()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer rawData.Close()

	// ioutil.ReadAll 已废弃（Go 1.16+），改用 io.ReadAll
	bytes, err := io.ReadAll(rawData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("value:%v", string(bytes))
}

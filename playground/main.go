package main

import "fmt"

// MixerFallbackConfig ...
type MixerFallbackConfig struct {
	Bundles map[string]BundleFallbackConfig `json:"bundles"`
}

// BundleFallbackConfig ...
type BundleFallbackConfig struct {
	EnableFallback bool `json:"enable_fallback"`
}

func main() {
	defaultConfig := MixerFallbackConfig{}

	config, ok := defaultConfig.Bundles["test"]
	// 修复：ok 是 bool，格式动词应为 %t（原代码用 %s 会输出 %!s(bool=false)）
	fmt.Printf("result:%v+%t", config, ok)
}

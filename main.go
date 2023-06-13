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
	fmt.Printf("result:%v+%s", config, ok)
}

package main

import "testing"

var result map[int]int

func buildMap(size int, capacity int) map[int]int {
	m := make(map[int]int, capacity)
	for i := 0; i < size; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkMapWithoutCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = buildMap(1_000, 0)
	}
}

func BenchmarkMapWithCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = buildMap(1_000, 1_000)
	}
}

package main

import (
	"strings"
	"testing"
)

func BenchmarkBoyerMoore(b *testing.B) {
	text := strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1000)
	pattern := "MNOPQRSTUVWXYZABCDEFGHIJKL"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMooreSearch(text, pattern)
	}
}

func BenchmarkStringsIndex(b *testing.B) {
	text := strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1000)
	pattern := "MNOPQRSTUVWXYZABCDEFGHIJKL"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.Index(text, pattern)
	}
}

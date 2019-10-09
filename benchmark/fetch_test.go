package main

import (
	"testing"
)

func BenchmarkFetch(b *testing.B) {
	urls := []string{
		"https://google.com",
		"https://github.com",
	}
	for n := 0; n < b.N; n++ {
		fetch(urls)
	}
}

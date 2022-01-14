package main_test

import (
	"encoding/json"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	b.SetBytes(123)
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal("Hello, world!")
	}
}

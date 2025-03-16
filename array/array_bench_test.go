package array

import (
	"testing"
)

func BenchmarkGet(b *testing.B) {
	a := NewArray()
	for i := 0; i < 1000; i++ {
		a.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = a.Get(500) // O(1) random access
	}
}

func BenchmarkRemove(b *testing.B) {
	a := NewArray()
	for i := 0; i < 1000; i++ {
		a.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Remove(500) // O(n) shifting
	}
}

func BenchmarkAppend(b *testing.B) {
	a := NewArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Append(i) // O(1) amortized
	}
}

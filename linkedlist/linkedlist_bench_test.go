package linkedlist

import (
	"testing"
)

func BenchmarkGet(b *testing.B) {
	l := NewLinkedList()
	for i := 0; i < 1000; i++ {
		l.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = l.Get(500) // O(n) traversal
	}
}

func BenchmarkRemove(b *testing.B) {
	l := NewLinkedList()
	for i := 0; i < 1000; i++ {
		l.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l.Remove(500) // O(n) traversal
	}
}

func BenchmarkAppend(b *testing.B) {
	l := NewLinkedList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Append(i) // O(n) traversal
	}
}

func BenchmarkRemoveAtNode(b *testing.B) {
	l := NewLinkedList()
	nodes := make([]*Node, 0)
	for i := 0; i < 1000; i++ {
		l.Append(i)
		current := l.Head
		for j := 0; j < i; j++ {
			current = current.Next
		}
		nodes = append(nodes, current)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l.RemoveAtNode(nodes[500]) // O(1) removal at known node
	}
}

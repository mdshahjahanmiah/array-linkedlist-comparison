package linkedlist

import "fmt"

// Node represents a node in the linked list
type Node struct {
	value int
	Next  *Node
}

// LinkedList represents the linked list structure
type LinkedList struct {
	Head *Node // Pointer to the first node
}

// NewLinkedList creates a new LinkedList instance
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds an element to the end (O(n))
func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Get accesses an element by index (O(n))
func (l *LinkedList) Get(index int) (int, error) {
	current := l.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return 0, fmt.Errorf("index out of bounds")
		}
		current = current.Next
	}
	if current == nil {
		return 0, fmt.Errorf("index out of bounds")
	}
	return current.value, nil
}

// Remove removes an element at a specific index (O(n) due to traversal)
func (l *LinkedList) Remove(index int) error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}
	if index == 0 {
		l.Head = l.Head.Next
		return nil
	}
	current := l.Head
	for i := 0; i < index-1; i++ {
		if current == nil || current.Next == nil {
			return fmt.Errorf("index out of bounds")
		}
		current = current.Next
	}
	if current.Next == nil {
		return fmt.Errorf("index out of bounds")
	}
	current.Next = current.Next.Next
	return nil
}

// RemoveAtNode removes a node with O(1) time if the node reference is provided
func (l *LinkedList) RemoveAtNode(node *Node) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}
	if l.Head == node {
		l.Head = node.Next
		return nil
	}
	current := l.Head
	for current != nil && current.Next != node {
		current = current.Next
	}
	if current == nil {
		return fmt.Errorf("node not found in list")
	}
	current.Next = node.Next
	return nil
}

// Print prints all elements (O(n))
func (l *LinkedList) Print() string {
	current := l.Head
	result := ""
	for current != nil {
		result += fmt.Sprintf("%d -> ", current.value)
		current = current.Next
	}
	result += "nil"
	return fmt.Sprintf("`%s`", result)
}

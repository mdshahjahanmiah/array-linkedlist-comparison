package array

import "fmt"

// ArrayOperations represents operations on a slice (dynamic array)
type ArrayOperations struct {
	data []int // Using a slice for dynamic sizing
}

// NewArray creates a new ArrayOperations instance
func NewArray() *ArrayOperations {
	return &ArrayOperations{
		data: make([]int, 0), // Start with empty slice
	}
}

// Append adds an element to the end (O(1) amortized)
func (a *ArrayOperations) Append(value int) {
	a.data = append(a.data, value)
}

// Get accesses an element by index (O(1))
func (a *ArrayOperations) Get(index int) (int, error) {
	if index < 0 || index >= len(a.data) {
		return 0, fmt.Errorf("index out of bounds")
	}
	return a.data[index], nil
}

// Remove removes an element at a specific index (O(n) due to shifting)
func (a *ArrayOperations) Remove(index int) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index out of bounds")
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	return nil
}

// Print prints all elements (O(n))
func (a *ArrayOperations) Print() string {
	return fmt.Sprintf("`%v`", a.data)
}

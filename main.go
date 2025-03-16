// main.go
package main

import (
	"fmt"

	"github.com/mdshahjahanmiah/array-linkedlist-comparison/array"
	"github.com/mdshahjahanmiah/array-linkedlist-comparison/linkedlist"
)

func main() {
	// Initialize both data structures
	arrayInst := array.NewArray()
	listInst := linkedlist.NewLinkedList()

	// Perform operations and collect results for the table
	var tableRows []string

	// Step 1: Append 5 elements
	elements := []int{10, 20, 30, 40, 50}
	for _, val := range elements {
		arrayInst.Append(val)
		listInst.Append(val)
	}
	arrayResult := arrayInst.Print()
	listResult := listInst.Print()
	tableRows = append(tableRows, fmt.Sprintf("| Append [10, 20, 30, 40, 50]  | %s | %s |", arrayResult, listResult))

	// Step 2: Access element at index 2
	arrayVal, _ := arrayInst.Get(2)
	listVal, _ := listInst.Get(2)
	tableRows = append(tableRows, fmt.Sprintf("| Get (Index 2)                | `%d` | `%d` |", arrayVal, listVal))

	// Step 3: Remove element at index 1
	arrayInst.Remove(1)
	current := listInst.Head
	for i := 0; i < 1; i++ { // Traverse to find node at index 1
		if current == nil {
			fmt.Println("Linked List Error: Index out of bounds")
			break
		}
		current = current.Next
	}
	if current != nil {
		_ = listInst.RemoveAtNode(current) // O(1) if we have the node
	}
	arrayResult = arrayInst.Print()
	listResult = listInst.Print()
	tableRows = append(tableRows, fmt.Sprintf("| Remove (Index 1)             | %s | %s |", arrayResult, listResult))

	// Print the Markdown table and summary
	fmt.Println("# Arrays vs. Linked Lists Comparison")
	fmt.Println()
	fmt.Println("| Step/Operation               | Array Result                  | Linked List Result                     |")
	fmt.Println("|------------------------------|-------------------------------|----------------------------------------|")
	for _, row := range tableRows {
		fmt.Println(row)
	}
	// Print the summary
	fmt.Println()
	fmt.Println("=== Summary ===")
	fmt.Println("- Arrays shine with fast random access (O(1) for Get) and cache efficiency, but struggle with dynamic size changes and middle insertions/deletions (O(n) for Remove).")
	fmt.Println("- Linked Lists shine with dynamic size and O(1) insertions/deletions at known positions, but struggle with random access (O(n) for Get) and cache inefficiency.")
}

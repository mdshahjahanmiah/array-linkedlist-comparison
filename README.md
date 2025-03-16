# Arrays vs. Linked Lists
This repository compares arrays and linked lists in Go, highlighting their performance differences. It implements both data structures and tests them with common operations to confirm their theoretical strengths and trade-offs in practice.
## Setup Instructions

### Clone the Repository
```bash
git clone https://github.com/mdshahjahanmiah/array-linkedlist-comparison
```
After successfully cloning the repository, open the terminal and navigate to the project directory using the command:
```bash
cd array-linkedlist-comparison
```

# Run the Application

Once you're inside the `array-linkedlist-comparison` directory, run the following command in the terminal:
```bash 
go run main.go
```
After successful run, the output will look like this:
#### Arrays vs. Linked Lists Comparison
| Step/Operation               | Array Result        | Linked List Result               |
|------------------------------|---------------------|----------------------------------|
| Append [10, 20, 30, 40, 50]  | `[10 20 30 40 50]` | `10 -> 20 -> 30 -> 40 -> 50 -> nil` |
| Get (Index 2)                | `30`               | `30`                             |
| Remove (Index 1)             | `[10 30 40 50]`    | `10 -> 30 -> 40 -> 50 -> nil`   |

#### Summary

- **Arrays** shine with fast random access (**O(1) for Get**) and cache efficiency but struggle with dynamic size changes and middle insertions/deletions (**O(n) for Remove**).
- **Linked Lists** shine with dynamic size and **O(1) insertions/deletions at known positions**, but struggle with random access (**O(n) for Get**) and cache inefficiency.

## Running the Tests

```bash
# Run array benchmarks
go test ./array -bench=.

# Run linked list benchmarks
go test ./linkedlist -bench=.
```

### Key Findings
Based on benchmark results on an Apple M3 CPU:

| Operation     | Array Performance | Linked List Performance | Ratio (LL/Array) |
|--------------|------------------|------------------------|------------------|
| Get (Index)  | 0.2529 ns/op     | 463.5 ns/op           | ~1,833x slower  |
| Remove (Index) | 43.35 ns/op     | 535.0 ns/op           | ~12x slower     |
| Append       | 7.106 ns/op      | 115,165 ns/op         | ~16,206x slower |

### Interpreting These Results

- **Array Get**: Extremely fast (0.2529 ns/op) due to O(1) direct memory access.
- **Linked List Get**: Much slower (463.5 ns/op) due to O(n) traversal needed.
- **Array Remove**: Fast (43.35 ns/op) despite being O(n) theoretically due to efficient memory operations.
- **Linked List Append**: The high time for standard append (115,165 ns/op) indicates the benchmark likely didn't use the tail pointer optimization.  

### Time Complexity Comparison

| Operation                | Array  | Linked List | Linked List (Optimized) |
|--------------------------|--------|-------------|--------------------------|
| Access (Get)            | O(1)   | O(n)        | O(n)                     |
| Insert/Remove (Beginning) | O(n)   | O(1)        | O(1)                     |
| Insert/Remove (Middle)  | O(n)   | O(n)        | O(n)                     |
| Insert/Remove (End)     | O(1)*  | O(n)        | O(1) with tail pointer   |
| Search                  | O(n)   | O(n)        | O(n)                     |


## Conclusion

The benchmark results confirm the theoretical strengths and weaknesses of each data structure:

### Arrays Excel When:
- Random access is frequent
- Collection size is known or changes infrequently
- Cache efficiency matters
- Memory usage needs to be minimized

### Linked Lists Excel When:
- Insertions/deletions at known positions are frequent
- Collection size changes frequently
- Memory overhead of pointers is acceptable

## Real-world Considerations

While theoretical time complexity is important, other factors often influence actual performance:

- **Cache Locality**: Arrays typically have better cache performance due to contiguous memory storage.
- **Implementation Details**: Optimizations like tail pointers significantly improve linked list performance.
- **Hardware Factors**: Modern CPU architecture favors predictable memory access patterns (arrays).
- **Collection Size**: For small collections, overhead factors often outweigh asymptotic complexity.  



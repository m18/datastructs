# Data structures
Data structures for the Go programming language

### Binary tree
Example:
```
package main
 
import (
    "fmt"
    "github.com/m18/datastructs"
)

func main() {
    bt, err := datastructs.NewBinaryTree(1, 2, 3, nil, 4, 5, 6, nil, nil, 7, 8)
    if err != nil {
        // handle the error here
        return
    }
    fmt.Println(bt.Value, bt.Left.Value, bt.Right.Value)
    fmt.Println(bt)
}
```
Prints:
```
1 2 3
           1
      /¯¯¯¯¯¯¯¯¯\
     2           3
      ¯\       /¯¯¯\
        4     5     6
       /¯\            
      7   8             
```

### Stack
Example:
```
package main
 
import (
    "fmt"
    "github.com/m18/datastructs"
)

func main() {
	s := datastructs.NewStack(1, 2)
	s.Push(3)

	for {
		v, ok := s.Pop()
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}
}
```
Prints:
```
3 true
2 true
1 true
0 false
```

### Heap
Example:
```
package main
 
import (
    "fmt"
    "github.com/m18/datastructs"
)

func main() {
	h := datastructs.NewMinHeap(5, 2, 4)
	h.Push(3)

	for {
		v, ok := h.Pop()
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}
}
```
Prints:
```
2 true
3 true
4 true
5 true
0 false
```
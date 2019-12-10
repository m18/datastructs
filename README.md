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
    fmt.Println(bt)
}
```
Prints:
```
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
    st := datastructs.NewStack(1, 2)
    st.Push(3)

    fmt.Println(st.Peek())
    fmt.Println(st)
    fmt.Println(st.Pop())
    fmt.Println(st)
}
```
Prints:
```
3
[1 2 3]
3
[1 2]
```
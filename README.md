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
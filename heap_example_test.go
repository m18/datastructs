package datastructs_test

import (
	"fmt"
	"github.com/m18/datastructs"
)

func ExampleHeap() {
	h := datastructs.NewMinHeap(5, 2, 4)
	h.Push(3)

	for {
		v, ok := h.Pop()
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}

	// Output:
	// 2 true
	// 3 true
	// 4 true
	// 5 true
	// 0 false
}

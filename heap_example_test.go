package datastructs_test

import (
	"fmt"

	"github.com/m18/datastructs"
)

func ExampleHeap() {
	h := datastructs.NewMinHeap(3, 2, 5, 4)
	h.Push(3)

	for {
		if h.Size() == 0 {
			break
		}
		fmt.Println(h.Pop())
	}

	// Output:
	// 2
	// 3
	// 4
}

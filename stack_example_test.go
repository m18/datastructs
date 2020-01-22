package datastructs_test

import (
	"fmt"
	"github.com/m18/datastructs"
)

func ExampleStack() {
	s := datastructs.NewStack(1, 2)
	s.Push(3)

	for {
		v, ok := s.Pop()
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}

	// Output:
	// 3 true
	// 2 true
	// 1 true
	// 0 false
}

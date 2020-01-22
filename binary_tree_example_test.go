package datastructs_test

import (
	"fmt"
	"github.com/m18/datastructs"
)

func ExampleBinaryTreeNode() {
	bt, err := datastructs.NewBinaryTree(1, 2, 3, nil, 4, 5, 6, nil, nil, 7, 8)
	if err != nil {
		// handle the error here
		return
	}
	fmt.Println(bt.Value, bt.Left.Value, bt.Right.Value)
	fmt.Println(bt)

	// TODO: Output:
	// 1 2 3
	//            1
	//       /¯¯¯¯¯¯¯¯¯\
	//      2           3
	//       ¯\       /¯¯¯\
	//         4     5     6
	//        /¯\
	//       7   8

}

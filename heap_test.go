package datastructs

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {

}

func TestHeap(t *testing.T) {
	tests := []struct {
		id                     int
		capacity               int
		items                  []int
		push                   []int
		popTimes               int
		expectMinPeekAfterPush int
		expectMinPeekAfterPop  int
		expectMaxPeekAfterPush int
		expectMaxPeekAfterPop  int
	}{
		{
			id:                     1,
			capacity:               0,
			items:                  []int{5, 3, 10, 15},
			push:                   []int{1, 2},
			popTimes:               3,
			expectMinPeekAfterPush: 1,
			expectMinPeekAfterPop:  5,
			expectMaxPeekAfterPush: 15,
			expectMaxPeekAfterPop:  3,
		},
		{
			id:                     2,
			capacity:               0,
			items:                  []int{1, 1, 1},
			push:                   []int{1, 0, 2},
			popTimes:               2,
			expectMinPeekAfterPush: 0,
			expectMinPeekAfterPop:  1,
			expectMaxPeekAfterPush: 2,
			expectMaxPeekAfterPop:  1,
		},
		{
			id:                     3,
			capacity:               0,
			items:                  []int{0, 0, 0},
			push:                   []int{0, 0},
			popTimes:               4,
			expectMinPeekAfterPush: 0,
			expectMinPeekAfterPop:  0,
			expectMaxPeekAfterPush: 0,
			expectMaxPeekAfterPop:  0,
		},
		{
			id:                     4,
			capacity:               3,
			items:                  []int{2, 5, 4},
			push:                   []int{3},
			popTimes:               1,
			expectMinPeekAfterPush: 2,
			expectMinPeekAfterPop:  3,
			expectMaxPeekAfterPush: 5,
			expectMaxPeekAfterPop:  4,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("heap test #%d", test.id), func(t *testing.T) {
			t.Parallel()
			for _, isMin := range []bool{true, false} {
				var (
					h                                       *Heap
					expectPeekAfterPush, expectPeekAfterPop int
				)
				if isMin {
					h = NewMinHeap(test.capacity, test.items...)
					expectPeekAfterPush = test.expectMinPeekAfterPush
					expectPeekAfterPop = test.expectMinPeekAfterPop
				} else {
					h = NewMaxHeap(test.capacity, test.items...)
					expectPeekAfterPush = test.expectMaxPeekAfterPush
					expectPeekAfterPop = test.expectMaxPeekAfterPop
				}
				for _, v := range test.push {
					h.Push(v)
				}
				if h.Peek() != expectPeekAfterPush {
					t.Errorf("expected %d after push but got %d", expectPeekAfterPush, h.Peek())
				}
				for i := 0; i < test.popTimes; i++ {
					h.Pop()
				}
				if h.Peek() != expectPeekAfterPop {
					t.Errorf("expected %d after pop but got %d", expectPeekAfterPop, h.Peek())
				}
			}
		})
	}
}

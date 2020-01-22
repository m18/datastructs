package datastructs

import (
	"fmt"
	"testing"
)

func TestParentIndex(t *testing.T) {
	tests := []struct {
		idx    int
		expect int
	}{
		{idx: 4, expect: 1},
		{idx: 5, expect: 2},
		{idx: 6, expect: 2},
		{idx: 7, expect: 3},
		{idx: 8, expect: 3},
		{idx: 9, expect: 4},
		{idx: 0, expect: -1},
		{idx: -1, expect: -1},
	}

	heap := NewMinHeap()

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("parent index of %d", test.idx), func(t *testing.T) {
			t.Parallel()
			if res := heap.parentIndex(test.idx); res != test.expect {
				t.Errorf("expected %d but got %d", test.expect, res)
			}
		})
	}
}

func TestChildIndices(t *testing.T) {
	tests := []struct {
		idx     int
		expectL int
		expectR int
	}{
		{idx: 0, expectL: 1, expectR: 2},
		{idx: 1, expectL: 3, expectR: 4},
		{idx: 2, expectL: 5, expectR: 6},
		{idx: 3, expectL: 7, expectR: 8},
		{idx: 4, expectL: 9, expectR: 10},
		{idx: -1, expectL: -1, expectR: -1},
	}

	heap := NewMinHeap()

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("child indices of %d", test.idx), func(t *testing.T) {
			t.Parallel()
			if l, r := heap.childIndices(test.idx); l != test.expectL || r != test.expectR {
				t.Errorf("expected l=%d and r=%d but got l=%d and r=%d", test.expectL, test.expectR, l, r)
			}
		})
	}
}

func TestHeap(t *testing.T) {
	tests := []struct {
		id                     int
		vals                   []int
		push                   []int
		popTimes               int
		expectMinPeekAfterPush int
		expectMinPeekAfterPop  int
		expectMaxPeekAfterPush int
		expectMaxPeekAfterPop  int
		expectNoOKAfterPush    bool
		expectNoOKAfterPop     bool
	}{
		{
			id:                     1,
			vals:                   []int{5, 3, 10, 15},
			push:                   []int{1, 2},
			popTimes:               3,
			expectMinPeekAfterPush: 1,
			expectMinPeekAfterPop:  5,
			expectMaxPeekAfterPush: 15,
			expectMaxPeekAfterPop:  3,
		},
		{
			id:                     2,
			vals:                   []int{1, 1, 1},
			push:                   []int{1, 0, 2},
			popTimes:               2,
			expectMinPeekAfterPush: 0,
			expectMinPeekAfterPop:  1,
			expectMaxPeekAfterPush: 2,
			expectMaxPeekAfterPop:  1,
		},
		{
			id:                     3,
			vals:                   []int{0, 0, 0},
			push:                   []int{0, 0},
			popTimes:               4,
			expectMinPeekAfterPush: 0,
			expectMinPeekAfterPop:  0,
			expectMaxPeekAfterPush: 0,
			expectMaxPeekAfterPop:  0,
		},
		{
			id:                     4,
			vals:                   []int{0},
			push:                   []int{0},
			popTimes:               3,
			expectMinPeekAfterPush: 0,
			expectMaxPeekAfterPush: 0,
			expectNoOKAfterPop:     true,
		},
		{
			id:                  5,
			vals:                []int{},
			push:                []int{},
			popTimes:            1,
			expectNoOKAfterPush: true,
			expectNoOKAfterPop:  true,
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
					h = NewMinHeap(test.vals...)
					expectPeekAfterPush = test.expectMinPeekAfterPush
					expectPeekAfterPop = test.expectMinPeekAfterPop
				} else {
					h = NewMaxHeap(test.vals...)
					expectPeekAfterPush = test.expectMaxPeekAfterPush
					expectPeekAfterPop = test.expectMaxPeekAfterPop
				}
				for _, v := range test.push {
					h.Push(v)
				}
				peek, ok := h.Peek()
				if test.expectNoOKAfterPush && ok {
					t.Errorf("got an unexpected OK after push")
				}
				if !test.expectNoOKAfterPush && !ok {
					t.Errorf("did not get an expected OK after push")
				}
				if peek != expectPeekAfterPush {
					t.Errorf("expected %d after push but got %d", expectPeekAfterPush, peek)
				}
				for i := 0; i < test.popTimes; i++ {
					h.Pop()
				}
				peek, ok = h.Peek()
				if test.expectNoOKAfterPop && ok {
					t.Errorf("got an unexpected OK after pop")
				}
				if !test.expectNoOKAfterPop && !ok {
					t.Errorf("did not get an expected ok after pop")
				}
				if peek != expectPeekAfterPop {
					t.Errorf("expected %d after pop but got %d", expectPeekAfterPop, peek)
				}
			}
		})
	}
}

package datastructs

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		id                  int
		vals                []int
		push                []int
		popTimes            int
		expectPeekAfterPush int
		expectPeekAfterPop  int
		expectNoOKAfterPush bool
		expectNoOKAfterPop  bool
	}{
		{
			id:                  1,
			vals:                []int{5, 3, 10, 15},
			push:                []int{1, 2},
			popTimes:            3,
			expectPeekAfterPush: 2,
			expectPeekAfterPop:  10,
		},
		{
			id:                  2,
			vals:                []int{1, 1, 1},
			push:                []int{1, 0, 2},
			popTimes:            2,
			expectPeekAfterPush: 2,
			expectPeekAfterPop:  1,
		},
		{
			id:                  3,
			vals:                []int{0, 0, 0},
			push:                []int{0, 0},
			popTimes:            4,
			expectPeekAfterPush: 0,
			expectPeekAfterPop:  0,
		},
		{
			id:                  4,
			vals:                []int{0},
			push:                []int{0},
			popTimes:            3,
			expectPeekAfterPush: 0,
			expectNoOKAfterPop:  true,
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
		t.Run(fmt.Sprintf("stack test #%d", test.id), func(t *testing.T) {
			t.Parallel()
			s := NewStack(test.vals...)
			for _, v := range test.push {
				s.Push(v)
			}
			peek, ok := s.Peek()
			if test.expectNoOKAfterPush && ok {
				t.Errorf("got an unexpected OK after push")
			}
			if !test.expectNoOKAfterPush && !ok {
				t.Errorf("did not get an expected OK after push")
			}
			if peek != test.expectPeekAfterPush {
				t.Errorf("expected %d after push but got %d", test.expectPeekAfterPush, peek)
			}
			for i := 0; i < test.popTimes; i++ {
				s.Pop()
			}
			peek, ok = s.Peek()
			if test.expectNoOKAfterPop && ok {
				t.Errorf("got an unexpected OK after pop")
			}
			if !test.expectNoOKAfterPop && !ok {
				t.Errorf("did not get an expected ok after pop")
			}
			if peek != test.expectPeekAfterPop {
				t.Errorf("expected %d after pop but got %d", test.expectPeekAfterPop, peek)
			}
		})
	}
}

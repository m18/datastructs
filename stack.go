package datastructs

import "fmt"

// Stack is a representation of the Stack data structure
type Stack struct {
	sl []int
}

// NewStack returns a new stack
func NewStack(vals ...int) *Stack {
	return &Stack{sl: vals}
}

// Push adds a value to the stack
func (s *Stack) Push(n int) {
	s.sl = append(s.sl, n)
}

// Peek returns the value from the top of the stack; it also returns a bool indicating whether the stack is empty or not
func (s *Stack) Peek() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	return s.sl[len(s.sl)-1], true
}

// Pop removes the value from the top of the stack and returns it; it also returns a bool indicating whether the stack is empty or not
func (s *Stack) Pop() (int, bool) {
	v, ok := s.Peek()
	if !ok {
		return 0, false
	}
	s.sl = s.sl[:len(s.sl)-1]
	return v, true
}

// Size returns the number of values on the stack
func (s *Stack) Size() int {
	return len(s.sl)
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.sl)
}

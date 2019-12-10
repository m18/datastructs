package datastructs

import "fmt"

type Stack struct {
	sl []int
}

func NewStack(vals ...int) *Stack {
	return &Stack{sl: vals}
}

func (s *Stack) Push(n int) {
	s.sl = append(s.sl, n)
}

func (s *Stack) Peek() int {
	return s.sl[len(s.sl)-1]
}

func (s *Stack) Pop() int {
	v := s.Peek()
	s.sl = s.sl[:len(s.sl)-1]
	return v
}

func (s *Stack) Size() int {
	return len(s.sl)
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.sl)
}

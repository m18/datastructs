package datastructs

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Errorf("expected size to be 0 but it's %d", s.Size())
	}
	s.Push(5)
	if s.Size() != 1 {
		t.Errorf("expected size to be 1 but it's %d", s.Size())
	}
	s.Push(10)
	v := s.Peek()
	if v != 10 || s.Size() != 2 {
		t.Errorf("expected size to be 2 & value to be 10 but got %d and %d", s.Size(), v)
	}
	v = s.Pop()
	if v != 10 || s.Size() != 1 {
		t.Errorf("expected size to be 1 & value to be 10 but got %d and %d", s.Size(), v)
	}
}

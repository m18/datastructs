package datastructs

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tests := []struct {
		desc   string
		input  []interface{}
		isErr  bool
		verify func(*BinaryTreeNode) error
	}{
		{
			desc:  "no nils",
			input: []interface{}{1, 2, 3, 4},
			verify: func(bt *BinaryTreeNode) error {
				if bt != nil && bt.Value == 1 {
					return nil
				}
				return errors.New("verification error")
			},
		},
		{
			desc:  "nil root",
			input: []interface{}{nil, 1},
			isErr: true,
		},
		{
			desc:  "nil parent",
			input: []interface{}{1, 2, 3, nil, 4, 5, 6, nil, 7, 8},
			isErr: true,
		},
		{
			desc:  "nil tree",
			input: []interface{}{nil, nil},
			verify: func(bt *BinaryTreeNode) error {
				if bt != nil {
					return errors.New("expected nil")
				}
				return nil
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			bt, err := NewBinaryTree(test.input...)
			switch {
			case test.isErr && err == nil:
				t.Errorf("expected an error but there was none")
			case !test.isErr && err != nil:
				t.Errorf("unxpected error: %v", err)
			case test.verify != nil:
				if err := test.verify(bt); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func TestBinaryTreeStat(t *testing.T) {
	tests := []struct {
		input  []interface{}
		depth  int
		maxVal int
	}{
		{
			input:  []interface{}{1, 2, 4, 3},
			depth:  3,
			maxVal: 4,
		},
		{
			input:  []interface{}{nil},
			depth:  0,
			maxVal: MinInt,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			t.Parallel()
			bt, _ := NewBinaryTree(test.input...)
			d, v := bt.stat()
			if d != test.depth {
				t.Errorf("expected depth to be %d but it was %d", test.depth, d)
			}
			if v != test.maxVal {
				t.Errorf("expected maxVal to be %d but it was %d", test.maxVal, v)
			}
		})
	}
}

func TestBinaryTreeBFT(t *testing.T) {
	bt, _ := NewBinaryTree(4, nil, 6, nil, nil, nil, 10)
	expect := []struct {
		empty bool
		value int
		index int
		level int
	}{
		{empty: false, value: 4, index: 0, level: 1},
		{empty: true, index: 0, level: 2},
		{empty: false, value: 6, index: 1, level: 2},
		{empty: true, index: 0, level: 3},
		{empty: true, index: 1, level: 3},
		{empty: true, index: 2, level: 3},
		{empty: false, value: 10, index: 3, level: 3},
	}
	expectedCount := len(expect)
	i := 0
	bt.bft(func(node *BinaryTreeNode, level, index int) {
		e := expect[i]
		i++
		if ((node == nil) != e.empty) || (node != nil && node.Value != e.value) || index != e.index || level != e.level {
			t.Errorf("expectations have not been met for %v", e)
		}
	})
	if i != expectedCount {
		t.Errorf("expected exec count to be %d but it was %d", expectedCount, i)
	}
}

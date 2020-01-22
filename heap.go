package datastructs

import (
	"math"
)

// Heap is a binary tree-based implementation of the Heap data structure
type Heap struct {
	items   []int
	compare func(n, to int) bool
}

// NewMinHeap returns a new min-heap
func NewMinHeap(vals ...int) *Heap {
	return newHeap(func(n, to int) bool { return n < to }, vals...)
}

// NewMaxHeap creates a new max-heap
func NewMaxHeap(vals ...int) *Heap {
	return newHeap(func(n, to int) bool { return n > to }, vals...)
}

func newHeap(c func(n, to int) bool, vals ...int) *Heap {
	res := &Heap{compare: c}
	for _, n := range vals {
		res.Push(n)
	}
	return res
}

// Size returns the number of values in the heap
func (h *Heap) Size() int {
	return len(h.items)
}

// Peek returns the item from the top of the heap; it also returns a bool indicating whether the heap is empty or not
func (h *Heap) Peek() (int, bool) {
	if h.Size() == 0 {
		return 0, false
	}
	return h.items[0], true
}

// Pop removes the value from the top of the heap and returns it; it also returns a bool indicating whether the heap is empty or not
func (h *Heap) Pop() (int, bool) {
	res, ok := h.Peek()
	if !ok {
		return 0, false
	}
	h.items[0] = h.items[h.Size()-1]
	h.items = h.items[:h.Size()-1]
	h.heapifyDown()
	return res, true
}

// Push adds a value to the heap
func (h *Heap) Push(n int) {
	h.items = append(h.items, n)
	h.heapifyUp()
}

func (h *Heap) heapifyUp() {
	if h.Size() == 0 {
		return
	}
	idx := h.Size() - 1
	for {
		parentIdx := h.parentIndex(idx)
		if parentIdx == -1 || !h.checkSwap(idx, parentIdx, false) {
			break
		}
		idx = parentIdx
	}
}

func (h *Heap) heapifyDown() {
	if h.Size() == 0 {
		return
	}
	idx := 0
	for {
		lChildIdx, rChildIdx := h.childIndices(idx)
		if lChildIdx > h.Size()-1 {
			break
		}
		childIdx := lChildIdx
		if rChildIdx <= h.Size()-1 && h.compare(h.items[rChildIdx], h.items[lChildIdx]) {
			childIdx = rChildIdx
		}
		if !h.checkSwap(idx, childIdx, true) {
			break
		}
		idx = childIdx
	}
}

// checkSwap checks if two nodes need to be swapped, and performs the operation if needed, in which case true is returned
func (h *Heap) checkSwap(idx1, idx2 int, reverse bool) bool {
	if reverse == h.compare(h.items[idx1], h.items[idx2]) {
		return false
	}
	h.items[idx1], h.items[idx2] = h.items[idx2], h.items[idx1]
	return true
}

// parentIndex returns a slice index that corresponds to the parent node of a binary tree node represented by the element of the slice at index idx
func (h *Heap) parentIndex(idx int) int {
	level, idxInLevel := h.nodePosition(idx)
	if level <= 0 {
		return -1
	}
	prevLevelOffset := 1<<(level-1) - 1
	parentIndex := prevLevelOffset + idxInLevel/2
	return parentIndex
}

func (h *Heap) childIndices(idx int) (int, int) {
	level, idxInLevel := h.nodePosition(idx)
	if level < 0 {
		return -1, -1
	}
	rChildIdxInLevel := (idxInLevel+1)*2 - 1
	lChildIdxInLevel := rChildIdxInLevel - 1
	nextLevelOffset := 1<<(level+1) - 1
	return nextLevelOffset + lChildIdxInLevel, nextLevelOffset + rChildIdxInLevel
}

func (h *Heap) nodePosition(idx int) (int, int) {
	if idx < 0 {
		return -1, -1
	}
	if idx == 0 {
		return 0, 0
	}
	// level is zero-based
	level := int(math.Sqrt(float64(idx + 1)))
	idxInLevel := idx - (1<<level - 1)
	return level, idxInLevel
}

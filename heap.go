package datastructs

type Heap struct {
	capacity int
	items    []int
	compare  func(i, j int) bool
}

// NewMinHeap returns a new min-heap. `capacity` == 0 means unbounded.
func NewMinHeap(capacity int, items ...int) *Heap {
	return newHeap(
		capacity,
		func(x, y int) bool { return x < y },
		items...,
	)
}

// NewMaxHeap returns a new max-heap. `capacity` == 0 means unbounded.
func NewMaxHeap(capacity int, items ...int) *Heap {
	return newHeap(
		capacity,
		func(x, y int) bool { return x > y },
		items...,
	)
}

// Size returns the number of items on the heap.
func (h *Heap) Size() int {
	return len(h.items)
}

// Push adds item onto the heap.
func (h *Heap) Push(item int) {
	h.items = append(h.items, item)
	h.heapifyUp()
	if h.capacity > 0 && len(h.items) > h.capacity {
		h.shrink()
	}
}

// Peek returns the top element of the heap. It will panic if the heap is empty.
func (h *Heap) Peek() int {
	return h.items[0]
}

// Pop removes the top element of the heap and returns it. It will panic if the heap is empty.
func (h *Heap) Pop() int {
	res := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.shrink()
	h.heapifyDown(0)
	return res
}

func newHeap(capacity int, compare func(x, y int) bool, items ...int) *Heap {
	res := &Heap{
		capacity: capacity,
	}
	res.compare = func(i, j int) bool {
		return compare(res.items[i], res.items[j])
	}
	for _, item := range items {
		res.Push(item)
	}
	return res
}

func (h *Heap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *Heap) leftChildIndex(i int) int {
	return i*2 + 1
}

func (h *Heap) rightChildIndex(i int) int {
	return i*2 + 2
}

func (h *Heap) isLeaf(i int) bool {
	return i+1 > len(h.items)/2
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) shrink() {
	h.items = h.items[:len(h.items)-1]
}

func (h *Heap) heapifyUp() {
	i := len(h.items) - 1
	for pi := h.parentIndex(i); i > 0 && h.compare(i, pi); pi = h.parentIndex(i) {
		h.swap(i, pi)
		i = pi
	}
}

func (h *Heap) heapifyDown(i int) {
	if h.isLeaf(i) {
		return
	}
	li := h.leftChildIndex(i)
	ri := h.rightChildIndex(i)
	// no right child:
	// because heap is a `complete` binary tree,
	// non-leaf nodes will always have the left child,
	// but not necessarily the right one (there will be 1 such node)
	if ri > len(h.items)-1 {
		ri = -1
	}
	if h.compare(i, li) && (ri == -1 || h.compare(i, ri)) {
		return
	}
	if ri == -1 || h.compare(li, ri) {
		h.swap(i, li)
		h.heapifyDown(li)
	} else {
		h.swap(i, ri)
		h.heapifyDown(ri)
	}
}

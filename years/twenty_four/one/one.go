package one

import (
	"container/heap"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func RunPartOne(input string) int {
	lines := strings.Split(input, "\n")
	leftHeap := &MinHeap{}
	rightHeap := &MinHeap{}

	heap.Init(leftHeap)
	heap.Init(rightHeap)

	for _, line := range lines {
		print(line)
	}

	return 1
}

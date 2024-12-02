package one

import (
	"container/heap"
	"math"
	"strconv"
	"strings"
	"text/scanner"
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

func MeasureDistance(first int, second int) int {
	difference := math.Abs(float64(first) - float64(second))
	return int(difference)
}

func RunPartOne(input string) int {
	leftHeap := &MinHeap{}
	rightHeap := &MinHeap{}

	heap.Init(leftHeap)
	heap.Init(rightHeap)

	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<' ' | 1<<'\t' | 1<<'\n' | 1<<'.'

	isLeft := true

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		num, err := strconv.Atoi(s.TokenText())
		if err != nil {
			continue
		}

		if isLeft {
			heap.Push(leftHeap, num)
		} else {
			heap.Push(rightHeap, num)
		}

		isLeft = !isLeft
	}

	if leftHeap.Len() != rightHeap.Len() {
		panic("Left and right lists are not of the same length")
	}

	sum := 0

	for leftHeap.Len() > 0 {
		leftValue := heap.Pop(leftHeap).(int)
		rightValue := heap.Pop(rightHeap).(int)

		sum += MeasureDistance(leftValue, rightValue)
	}

	return sum
}

func RunPartTwo(input string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<' ' | 1<<'\t' | 1<<'\n' | 1<<'.'

	isLeft := true

	leftCount := map[int]int{}
	var leftValues []int
	var rightValues []int

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		num, err := strconv.Atoi(s.TokenText())
		if err != nil {
			continue
		}

		if isLeft {
			leftCount[num] = 0
			leftValues = append(leftValues, num)
		} else {
			rightValues = append(rightValues, num)
		}

		isLeft = !isLeft
	}

	for i := 0; i < len(rightValues); i++ {
		val, ok := leftCount[rightValues[i]]
		if ok {
			leftCount[rightValues[i]] = val + 1
		}
	}

	sum := 0
	for _, v := range leftValues {
		frequency := leftCount[v]
		sum += frequency * v
	}

	return sum
}

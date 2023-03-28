package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	B := make([]int, N)
	sumall := 0
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		sumall += A[i]
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &B[i])
		sumall -= B[i]
	}
	sum := 0
	ans := 0
	pq := &Int64Heap{}
	for i := 0; i < N; i++ {
		sum += A[i] - B[i]
		cur := sumall - sum
		ans += abs64(cur)
		if cur <= 0 {
			heap.Push(pq, 0)
			continue
		}
		if cur > sumall {
			cur = sumall
		}
		heap.Push(pq, -cur)
		heap.Push(pq, -cur)
		ans += heap.Pop(pq).(int)
	}
	fmt.Println(ans)
}

type Int64Heap []int

func (h Int64Heap) Len() int {
	return len(h)
}

func (h Int64Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Int64Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Int64Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Int64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h Int64Heap) Top() int {
	return h[0]
}

func abs64(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

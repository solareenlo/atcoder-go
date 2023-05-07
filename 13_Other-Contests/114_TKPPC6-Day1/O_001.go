package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 1 << 60

	var N int
	fmt.Fscan(in, &N)

	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &L[i], &R[i])
	}

	data := make([]pair, 0)
	for i := 0; i < N; i++ {
		data = append(data, pair{L[i], R[i]})
	}
	sortPair(data)
	data = append(data, pair{INF, INF})

	now := -INF
	que := &Heap{}
	for i := 0; i < N+1; i++ {
		for now < data[i].x && que.Len() > 0 {
			r := heap.Pop(que).(int)
			if r < now {
				fmt.Println(-1)
				return
			}
			now++
		}
		now = data[i].x
		heap.Push(que, data[i].y)
	}

	left := 0
	right := (1 << 32)
	for left+1 < right {
		mid := (left + right) / 2
		data1 := make([]pair, 0)
		for i := 0; i < N; i++ {
			data1 = append(data1, pair{L[i], R[i]})
		}
		sortPair(data1)
		data1 = append(data1, pair{INF, INF})
		mini := INF
		for i := range R {
			if mini > R[i] {
				mini = R[i]
			}
		}
		now := mini - mid
		idx := 0
		flag := true
		que1 := &Heap{}
		for idx != N {
			for now+mid >= data1[idx].x {
				heap.Push(que1, data1[idx].y)
				idx++
			}
			if que1.Len() == 0 {
				flag = false
				break
			}
			now = min(now+mid, heap.Pop(que1).(int))
		}
		if flag {
			right = mid
		} else {
			left = mid
		}
	}
	fmt.Println(right)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

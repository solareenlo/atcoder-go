package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)

	var t [200001]int
	t[1] = 0
	for i := 1; i < N-1; i++ {
		fmt.Fscan(in, &t[i+1])
	}
	var to, ne [400002]int
	var X, Y, he [200001]int
	k := 2
	for i := 1; i < M+1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		to[k] = b
		ne[k] = he[a]
		he[a] = k
		k++
		to[k] = a
		ne[k] = he[b]
		he[b] = k
		k++
		fmt.Fscan(in, &X[i], &Y[i])
	}

	var D [200001]int
	for i := 1; i < N+1; i++ {
		D[i] = int(1e18)
	}
	D[1] = 0

	que := &Heap{}
	heap.Push(que, 1)

	m := (1 << 20) - 1
	for que.Len() > 0 {
		p := heap.Pop(que).(int)
		i := p & m
		if D[i] != p>>20 {
			continue
		}
		if D[i] > K {
			break
		}
		d := D[i] + t[i]
		for q := he[i]; q > 0; q = ne[q] {
			j := to[q]
			x := X[q>>1]
			y := Y[q>>1]
			d2 := int((float64(d)+float64(y)-0.5)/float64(y))*y + x
			if D[j] > d2 {
				D[j] = d2
				heap.Push(que, (d2<<20)+j)
			}
		}
	}
	if D[N] <= K {
		fmt.Println(D[N])
	} else {
		fmt.Println(-1)
	}
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

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var A [20]int
	var E [202020][]pair
	var dist, mask [101010]int
	var p2 [202020]int
	var num [1 << 20]int

	var N, M, K, D int
	fmt.Fscan(in, &N, &M, &K, &D)
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < M; i++ {
		var x, y, r int
		fmt.Fscan(in, &x, &y, &r)
		E[x] = append(E[x], pair{y, r})
		E[y] = append(E[y], pair{x, r})
	}

	for i := 0; i < K; i++ {
		for x := 0; x < N; x++ {
			dist[x] = 1 << 60
		}
		Q := make(HeapPair, 0)
		heap.Init(&Q)
		dist[A[i]] = 0
		heap.Push(&Q, pair{0, A[i]})
		for Q.Len() > 0 {
			tmp := heap.Pop(&Q).(pair)
			co := -tmp.x
			cur := tmp.y
			if co != dist[cur] {
				continue
			}
			if co > D {
				break
			}
			mask[cur] |= 1 << i
			for _, e := range E[cur] {
				if dist[e.x] > co+e.y {
					dist[e.x] = co + e.y
					heap.Push(&Q, pair{-dist[e.x], e.x})
				}
			}
		}
	}

	p2[0] = 1
	for i := 0; i < N; i++ {
		num[mask[i]]++
		p2[i+1] = p2[i] * 2 % MOD
	}
	for i := 0; i < K; i++ {
		for x := 0; x < 1<<K; x++ {
			if (x & (1 << i)) != 0 {
				num[x^(1<<i)] += num[x]
			}
		}
	}
	for x := 0; x < 1<<K; x++ {
		num[x] = p2[num[x]] - 1
	}
	for i := 0; i < K; i++ {
		for x := 0; x < 1<<K; x++ {
			if (x & (1 << i)) == 0 {
				num[x] = (num[x] + MOD - num[x^(1<<i)]) % MOD
			}
		}
	}

	ret := 0
	for x := 0; x < 1<<K; x++ {
		ret ^= num[x] % MOD
	}
	fmt.Println(ret)
}

type pair struct {
	x, y int
}

type HeapPair []pair

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

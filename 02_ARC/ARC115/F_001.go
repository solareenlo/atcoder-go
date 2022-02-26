package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 2020

var (
	n, k int
	e    = make([][]int, N)
	h    = [N]int{}
	to   = [N]int{}
	w    = [N]int{}
	s    = [N]int{}
	t    = [N]int{}
	hs   = &Heap{}
	ht   = &Heap{}
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
	}
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	fmt.Fscan(in, &k)
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &s[i], &t[i])
	}

	for i := 1; i <= n; i++ {
		id := -1
		mn := 1 << 60 / 2
		var dfs func(u, fa, x int)
		dfs = func(u, fa, x int) {
			if (h[u] < h[i] || (h[u] == h[i] && (u < i))) && x < mn {
				mn = x
				id = u
			}
			for _, v := range e[u] {
				if v != fa {
					dfs(v, u, max(x, h[v]))
				}
			}
			return
		}
		dfs(i, 0, h[i])
		to[i] = id
		w[i] = mn - h[i]
	}

	cnt := 0
	S := 0
	T := 0
	for i := 1; i <= k; i++ {
		if s[i] != t[i] {
			cnt++
		}
		S += h[s[i]]
		T += h[t[i]]
		if to[s[i]] != -1 {
			heap.Push(hs, pair{-w[s[i]], i})
		}
		if to[t[i]] != -1 {
			heap.Push(ht, pair{-w[t[i]], i})
		}
	}

	ans := max(S, T)
	for cnt > 0 {
		if hs.Len() > 0 && (ht.Len() == 0 || S-(*hs)[0].x < T-(*ht)[0].x) {
			x := (*hs)[0].x
			i := (*hs)[0].y
			heap.Pop(hs)
			ans = max(ans, S-x)
			S -= h[s[i]]
			if s[i] != t[i] {
				cnt--
			}
			s[i] = to[s[i]]
			if s[i] != t[i] {
				cnt++
			}
			S += h[s[i]]
			if to[s[i]] != -1 {
				heap.Push(hs, pair{-w[s[i]], i})
			}
		} else {
			x := (*ht)[0].x
			i := (*ht)[0].y
			heap.Pop(ht)
			ans = max(ans, T-x)
			if t[i] >= 0 {
				T -= h[t[i]]
			}
			if s[i] != t[i] {
				cnt--
			}
			if t[i] >= 0 {
				t[i] = to[t[i]]
			}
			if s[i] != t[i] {
				cnt++
			}
			if t[i] >= 0 {
				T += h[t[i]]
			}
			if t[i] >= 0 && to[t[i]] != -1 {
				heap.Push(ht, pair{-w[t[i]], i})
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair struct{ x, y int }

type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

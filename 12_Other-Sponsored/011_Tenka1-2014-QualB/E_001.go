package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type pairIF struct {
	a int
	b float64
}

var L, N, M int
var A, cnt [100009]int
var K, T [100009]float64
var D [100009][]pairIF
var ans []pair

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &L, &N, &M)
	for i := 1; i < L+1; i++ {
		var d float64
		fmt.Fscan(in, &d)
		K[i] = d
	}
	for i := 1; i < N+1; i++ {
		var d float64
		fmt.Fscan(in, &A[i], &d)
		T[i] = d
		D[A[i]] = append(D[A[i]], pairIF{i, T[i]})
	}
	for i := 1; i < L+1; i++ {
		sum := 0.0
		for j := 0; j < len(D[i]); j++ {
			sum += K[i] - D[i][j].b
			if sum != 0.0 {
				add(i, State{0, 0, -1.0 * float64(j+1), sum, D[i][j].a})
			}
		}
	}
	Q := &HeapTuple{}
	for i := 1; i < L+1; i++ {
		Z := get_value(i, 1)
		heap.Push(Q, tuple{Z.a + K[i], Z.b, i})
	}
	for i := 0; i < M; i++ {
		W := heap.Pop(Q).(tuple)
		ans = append(ans, pair{W.y, W.z})
		cnt[W.z]++
		Z := get_value(W.z, float64(cnt[W.z]+1))
		heap.Push(Q, tuple{Z.a + K[W.z], Z.b, W.z})
	}
	sortPair(ans)
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%d %d\n", ans[i].a, ans[i].b)
	}
}

type pairFI struct {
	a float64
	b int
}

func get_value(p int, pos float64) pairFI {
	if len(S[p]) == 0 {
		return pairFI{0.0, 0}
	}
	pos1 := lowerBound(S[p], State{pos, pos, 0.0, 0.0, 0})
	pos1--
	T := pos - S[p][pos1].A
	return pairFI{S[p][pos1].B / T, S[p][pos1].id}
}

var S [1000009][]State

type State struct {
	L, R, A, B float64
	id         int
}

func lowerBound(a []State, x State) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].L >= x.L
	})
	return idx
}

func add(p int, L State) {
	if len(S[p]) == 0 {
		S[p] = append(S[p], State{-1e9, 1e9, L.A, L.B, L.id})
		return
	}

	G := 1e9
	for {
		F := cross(S[p][len(S[p])-1], L)
		if F < -1e8 {
			if G < 1e8 {
				S[p] = append(S[p], State{G, 1e9, L.A, L.B, L.id})
			}
			break
		}
		if F > S[p][len(S[p])-1].L {
			S[p][len(S[p])-1].R = F
			S[p] = append(S[p], State{F, 1e9, L.A, L.B, L.id})
			break
		}
		G = S[p][len(S[p])-1].L
		S[p] = S[p][:len(S[p])-1]
	}
}

func cross(C1, C2 State) float64 {
	if C1.A > C2.A {
		C1, C2 = C2, C1
	}
	if C2.B-C1.B >= -1e-9 {
		return -1e9
	}
	E1 := C1.B*C2.A - C2.B*C1.A
	E2 := C1.B - C2.B
	return E1 / E2
}

type pair struct {
	a, b int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].a == tmp[j].a {
			return tmp[i].b < tmp[j].b
		}
		return tmp[i].a < tmp[j].a
	})
}

type tuple struct {
	x    float64
	y, z int
}

type HeapTuple []tuple

func (h HeapTuple) Len() int            { return len(h) }
func (h HeapTuple) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h HeapTuple) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapTuple) Push(x interface{}) { *h = append(*h, x.(tuple)) }

func (h *HeapTuple) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

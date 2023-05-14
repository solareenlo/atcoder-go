package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N+1)
	B := make([]int, N+1)
	for i := 2; i <= N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 2; i <= N; i++ {
		fmt.Fscan(in, &B[i])
	}
	A_ := make([]int, N+1)
	B_ := make([]int, N+1)
	copy(A_, A)
	copy(B_, B)
	sort.Ints(A_)
	sort.Ints(B_)
	for i := 0; i < N+1; i++ {
		if A_[i] != B_[i] {
			fmt.Println(-1)
			return
		}
	}
	mndv := pf(N)
	a := make([]int, 0)
	b := make([]int, 0)
	t := make([]int, 0)
	for i := 2; i <= N; i++ {
		if N < 2*i && mndv[i] == i {
			if A[i] != B[i] {
				fmt.Println(-1)
				return
			}
		} else {
			t = append(t, i)
			a = append(a, A[i])
			b = append(b, B[i])
		}
	}
	idx := make([]*Heap, 50001)
	for i := range idx {
		idx[i] = &Heap{}
	}
	for i := len(a) - 1; i >= 0; i-- {
		heap.Push(idx[a[i]], i)
	}

	ans := make([]pair, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			heap.Pop(idx[a[i]])
			continue
		}
		j := (*idx[b[i]])[0]
		heap.Pop(idx[a[i]])
		heap.Pop(idx[a[j]])
		heap.Push(idx[a[i]], j)
		swp(t[i], t[j], &ans, mndv)
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(len(ans))
	for k := 0; k < len(ans); k++ {
		i := ans[k].x
		j := ans[k].y
		if i > j {
			i, j = j, i
		}
		fmt.Println(i, j)
	}
}

func pf(N int) []int {
	p := make([]int, N+1)
	for i := range p {
		p[i] = -1
	}
	for i := 2; i <= N; i++ {
		if p[i] != -1 {
			continue
		}
		for j := i; j <= N; j += i {
			if p[j] == -1 {
				p[j] = i
			}
		}
	}
	return p
}

func swp(i, j int, ans *[]pair, mndv []int) {
	p := mndv[i]
	q := mndv[j]
	v := []int{i, p, 2 * p, 2, 2 * q, q, j}
	for {
		f := false
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				if v[i] == v[j] {
					f = true
					v = eraseSlice(v, i, j)
					break
				}
			}
		}
		if f == false {
			break
		}
	}
	for i := 0; i < len(v)-1; i++ {
		*ans = append(*ans, pair{v[i], v[i+1]})
	}
	for i := len(v) - 3; i >= 0; i-- {
		*ans = append(*ans, pair{v[i], v[i+1]})
	}
}

func eraseSlice(s []int, start int, end int) []int {
	// スライスの一部を削除する
	copy(s[start:], s[end:])
	// 新しいスライスを返す
	return s[:len(s)-(end-start)]
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

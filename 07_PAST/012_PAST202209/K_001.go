package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	var N, M int
	fmt.Fscan(in, &N, &M)
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i], &B[i])
		A[i]--
		B[i]--
	}
	var Q int
	fmt.Fscan(in, &Q)
	t := make([]int, Q)
	x := make([]int, Q)
	y := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &t[i], &x[i], &y[i])
		x[i]--
		y[i]--
	}
	ans := make([]int, Q)
	for i := range ans {
		ans[i] = -1
	}
	S := make(map[pair]bool)
	for i := 0; i < M; i++ {
		S[pair{A[i], B[i]}] = true
	}
	for i := 0; i < Q; i++ {
		if t[i] == 1 {
			S[pair{x[i], y[i]}] = true
		} else if t[i] == 2 {
			delete(S, pair{x[i], y[i]})
		}
	}
	uf := NewDsu(N)
	for key, _ := range S {
		uf.Merge(key.x, key.y)
	}
	for qq := Q - 1; qq >= 0; qq-- {
		if t[qq] == 1 {
			delete(S, pair{x[qq], y[qq]})
			ufn := NewDsu(N)
			for key, _ := range S {
				ufn.Merge(key.x, key.y)
			}
			uf = ufn
		} else if t[qq] == 2 {
			S[pair{x[qq], y[qq]}] = true
			uf.Merge(x[qq], y[qq])
		} else {
			if uf.Same(x[qq], y[qq]) {
				ans[qq] = 1
			} else {
				ans[qq] = 0
			}
		}
	}
	for i := 0; i < Q; i++ {
		if ans[i] == -1 {
			continue
		}
		if ans[i] != 0 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

type dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.parentOrSize = make([]int, d.n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *dsu) Groups() [][]int {
	leaderBuf := make([]int, d.n)
	groupSize := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B, T, X, Y [1 << 17]int
	var Z [1 << 17]bool

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &A[i], &B[i])
		A[i]--
		B[i]--
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &T[i], &X[i])
		X[i]--
		if T[i] == 1 {
			Z[X[i]] = true
		} else {
			fmt.Fscan(in, &Y[i])
			Y[i]--
		}
	}
	uf := NewDsu(n)
	for i := 0; i < m; i++ {
		if !Z[i] {
			uf.Merge(A[i], B[i])
		}
	}
	for i := q - 1; i >= 0; i-- {
		if T[i] == 1 {
			uf.Merge(A[X[i]], B[X[i]])
		} else {
			Z[i] = uf.Same(X[i], Y[i])
		}
	}
	for i := 0; i < q; i++ {
		if T[i] == 2 {
			if Z[i] {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
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

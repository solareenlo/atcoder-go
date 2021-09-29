package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, R int
	fmt.Scan(&n, &m, &R)
	r := make([]int, R)
	for i := range r {
		fmt.Scan(&r[i])
		r[i]--
	}
	sort.Ints(r)

	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = int(1e9 + 7)
		}
	}

	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b, &c)
		a--
		b--
		d[a][b] = c
		d[b][a] = c
	}

	for i := 0; i < n; i++ {
		d[i][i] = 0
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	res := int(1e9 + 7)
	for NextPermutation(sort.IntSlice(r)) {
		sum := 0
		for i := 0; i < R-1; i++ {
			sum += d[r[i]][r[i+1]]
		}
		res = min(res, sum)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

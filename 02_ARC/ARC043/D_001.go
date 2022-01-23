package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	l := 0
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
		l += a[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	p := make([]int, 50005)
	q := make([]int, 50005)
	for k := 1; k <= 50000; k++ {
		p[k] = -1 << 60
	}

	for i, j := 0, 0; i < m; i++ {
		for k := 0; k <= 50000; k++ {
			q[k] = p[k] + (j-k)*(l-j+k)
			if k >= a[i] {
				q[k] = max(q[k], p[k-a[i]]+(k-a[i])*(l-k+a[i]))
			}
		}
		j += a[i]
		p, q = q, p
	}

	res := 0
	for k := 0; k+k <= l; k++ {
		res = max(res, p[k]+k*(l-k)*(n-m+1))
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

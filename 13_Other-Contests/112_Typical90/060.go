package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	l := make([]int, n+1)
	for i := range l {
		l[i] = INF
	}
	l[0] = 0
	r := make([]int, n+1)
	for i := range r {
		r[i] = INF
	}
	r[0] = 0

	L := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		L[i] = lowerBound(l, a[i])
		l[L[i]] = a[i]
	}

	m := 0
	for i := n - 1; i >= 0; i-- {
		R := lowerBound(r, a[i])
		r[R] = a[i]
		m = max(m, L[i]+R-1)
	}
	fmt.Println(m)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

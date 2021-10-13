package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	p := make([]int, m)
	y := make([]int, m)
	num := make([][]int, n+1)
	for i := 0; i < m; i++ {
		fmt.Scan(&p[i], &y[i])
		num[p[i]] = append(num[p[i]], y[i])
	}

	for i := 0; i < n; i++ {
		sort.Ints(num[i+1])
	}

	for i := 0; i < m; i++ {
		index := sort.SearchInts(num[p[i]], y[i])
		fmt.Printf("%06d%06d\n", p[i], index+1)
	}
}

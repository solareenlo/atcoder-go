package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][3]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	res := -1 << 60
	for bit := 0; bit < 8; bit++ {
		vec := make([]int, n)
		for i := 0; i < n; i++ {
			for j := 0; j < 3; j++ {
				if (bit>>j)&1 != 0 {
					vec[i] += a[i][j]
				} else {
					vec[i] -= a[i][j]
				}
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(vec)))
		sum := 0
		for i := 0; i < m; i++ {
			sum += vec[i]
		}
		res = max(res, sum)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, m+2)
	a[0] = 0
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i+1])
	}
	sort.Ints(a[:m+1])
	a[m+1] = n + 1

	mini := n
	for i := 0; i < m+1; i++ {
		if a[i+1]-a[i]-1 != 0 {
			mini = min(mini, a[i+1]-a[i]-1)
		}
	}

	res := 0
	for i := 0; i < m+1; i++ {
		res += (a[i+1] - a[i] - 1) / mini
		if (a[i+1]-a[i]-1)%mini != 0 {
			res++
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

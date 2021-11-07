package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	dp := make([]int, 1000003)

	for _, x := range a {
		if dp[x] != 0 {
			dp[x] = 2
			continue
		}
		for i := x; i < 1000003; i += x {
			dp[i]++
		}
	}

	res := 0
	for _, x := range a {
		if dp[x] == 1 {
			res++
		}
	}

	fmt.Println(res)
}

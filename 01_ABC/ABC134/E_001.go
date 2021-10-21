package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 10
	}

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		index := sort.SearchInts(dp, -a+1)
		dp[index] = -a
	}

	fmt.Println(sort.SearchInts(dp, 1))
}

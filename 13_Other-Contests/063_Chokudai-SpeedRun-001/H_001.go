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
		dp[i] = 1000000007
	}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		dp[sort.SearchInts(dp, a)] = a
	}
	fmt.Println(sort.SearchInts(dp, 1000000007))
}

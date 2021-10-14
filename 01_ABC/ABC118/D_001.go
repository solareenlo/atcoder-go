package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, m)
	for i := range a {
		fmt.Scan(&a[i])
	}

	d := []int{0, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	dp := make([]string, 101010)
	for i := 0; i < n; i++ {
		if i != 0 && dp[i] == "" {
			continue
		}
		for j := range a {
			dp[i+d[a[j]]] = max(dp[i+d[a[j]]], strconv.Itoa(a[j])+dp[i])
		}
	}
	fmt.Println(dp[n])
}

func max(a, b string) string {
	if len(a) > len(b) || (len(a) == len(b)) && a > b {
		return a
	}
	return b
}

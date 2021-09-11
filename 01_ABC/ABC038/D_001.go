package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][2]int, n)
	var w, h int
	for i := 0; i < n; i++ {
		fmt.Scan(&w, &h)
		a[i] = [2]int{w, h}
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] > a[j][1]
		}
		return a[i][0] < a[j][0]
	})

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 << 60
	}

	res := 0
	for i := 0; i < n; i++ {
		idx := sort.SearchInts(dp, a[i][1])
		dp[idx] = a[i][1]
		if res < idx {
			res = idx
		}
	}
	fmt.Println(res + 1)
}

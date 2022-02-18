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
		a[i] -= i + 1
	}
	sort.Ints(a)

	med := a[n/2]

	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(a[i] - med)
	}
	fmt.Println(sum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

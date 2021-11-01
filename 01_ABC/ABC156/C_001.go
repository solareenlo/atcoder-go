package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}
	sort.Ints(x)

	mini := 1 << 60
	for i := x[0]; i <= x[n-1]; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += (x[j] - i) * (x[j] - i)
		}
		mini = min(mini, sum)
	}

	fmt.Println(mini)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

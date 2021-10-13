package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	h := make([]int, n)
	for i := range h {
		fmt.Scan(&h[i])
	}
	sort.Ints(h)

	sum := make([]int, n)
	for i := 0; i < n-1; i++ {
		sum[i+1] = sum[i] + abs(h[i]-h[i+1])
	}

	mini, index := 1<<60, 0
	for i := k - 1; i < n; i++ {
		diff := h[i] - h[i-(k-1)]
		if mini > diff {
			mini = diff
			index = i
		}
	}
	fmt.Println(h[index] - h[index-(k-1)])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	x := make([]int, m)
	for i := range x {
		fmt.Scan(&x[i])
	}
	sort.Ints(x)

	if n >= m {
		fmt.Println(0)
	} else {
		diff := make([]int, m-1)
		for i := 0; i < m-1; i++ {
			diff[i] = abs(x[i] - x[i+1])
		}
		sort.Ints(diff)

		sum := 0
		for i := 0; i < m-n; i++ {
			sum += diff[i]
		}
		fmt.Println(sum)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

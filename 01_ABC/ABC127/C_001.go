package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	l := make([]int, m)
	r := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&l[i], &r[i])
	}

	mini, maxi := 1<<60, 0
	for i := 0; i < m; i++ {
		maxi = max(maxi, l[i])
		mini = min(mini, r[i])
	}

	if mini-maxi < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(mini - maxi + 1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

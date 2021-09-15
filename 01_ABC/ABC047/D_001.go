package main

import "fmt"

func main() {
	var n, t, a int
	fmt.Scan(&n, &t)

	mini, maxi := int(1e12), 0
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		m[a-mini]++
		maxi = max(maxi, a-mini)
		mini = min(mini, a)
	}
	fmt.Println(m[maxi])
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

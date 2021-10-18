package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	mini := sum
	pre_sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		pre_sum += a[i]
		mini = min(mini, abs(pre_sum-(sum-pre_sum)))
	}

	fmt.Println(mini)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

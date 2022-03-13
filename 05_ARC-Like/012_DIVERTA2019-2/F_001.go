package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := [10]int{0, 1, 2, 4, 7, 12, 20, 29, 38, 52}
	f := make([]int, 10)
	f[2] = 1
	for i := 3; i < n; i++ {
		f[i] = f[i-1] * (a[i-1] + a[i-2] + 1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Print(0, " ")
			} else {
				fmt.Print(f[max(i, j)-1]*a[min(i, j)]+1, " ")
			}
		}
		fmt.Println()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

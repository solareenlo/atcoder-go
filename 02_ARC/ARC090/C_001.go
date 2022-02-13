package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, 2)
	for i := 0; i < 2; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	maxi := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n-i; j++ {
			sum += a[0][j]
		}
		for j := 0; j < i+1; j++ {
			sum += a[1][n-1-j]
		}
		maxi = max(maxi, sum)
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	maxi := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			maxi = max(maxi, (x[i]-x[j])*(x[i]-x[j])+(y[i]-y[j])*(y[i]-y[j]))
		}
	}

	fmt.Println(math.Sqrt(float64(maxi)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

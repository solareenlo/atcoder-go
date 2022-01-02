package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	maxi := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := math.Hypot(x[i]-x[j], y[i]-y[j])
			maxi = max(maxi, dist)
		}
	}

	fmt.Println(maxi)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

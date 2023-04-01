package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	x1 := make([]float64, n)
	y1 := make([]float64, n)
	r1 := make([]float64, n)
	x2 := make([]float64, m)
	y2 := make([]float64, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&x1[i], &y1[i], &r1[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&x2[i], &y2[i])
	}

	ans := 1e9
	for i := 0; i < n; i++ {
		ans = math.Min(ans, r1[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = math.Min(ans, math.Hypot(x1[i]-x2[j], y1[i]-y2[j])-r1[i])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			ans = math.Min(ans, math.Hypot(x2[i]-x2[j], y2[i]-y2[j])/2)
		}
	}
	fmt.Println(ans)
}

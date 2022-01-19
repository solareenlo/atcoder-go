package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([][5]float64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&p[i][j])
		}
	}

	maxi := 0.0
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < 4; j++ {
			sum += p[i][j]
		}
		sum += p[i][4] * 110.0 / 900.0
		maxi = max(maxi, sum)
	}

	fmt.Println(maxi)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

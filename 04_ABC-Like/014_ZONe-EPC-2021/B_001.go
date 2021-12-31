package main

import "fmt"

func main() {
	var n int
	var D, H float64
	fmt.Scan(&n, &D, &H)

	maxi := 0.0
	for i := 0; i < n; i++ {
		var d, h float64
		fmt.Scan(&d, &h)
		maxi = max(maxi, (H-h)/(D-d)*(-d)+h)
	}

	fmt.Println(maxi)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

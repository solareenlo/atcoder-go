package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		x[i] = abs(t)
		sum += x[i]
	}

	fmt.Println(sum)
	y := 0
	for i := 0; i < n; i++ {
		y += pow(x[i], 2)
	}
	fmt.Println(math.Sqrt(float64(y)))

	maxi := 0
	for i := 0; i < n; i++ {
		maxi = max(maxi, x[i])
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

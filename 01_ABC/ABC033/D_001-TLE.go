package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	var dist func(i, j int) int
	dist = func(i, j int) int {
		return (x[i]-x[j])*(x[i]-x[j]) + (y[i]-y[j])*(y[i]-y[j])
	}

	a1, a2, a3 := 0, 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				d1 := dist(i, j)
				d2 := dist(i, k)
				d3 := dist(j, k)
				sum := d1 + d2 + d3
				maxi := max(d1, max(d2, d3))
				switch {
				case sum-maxi > maxi:
					a1++
				case sum-maxi == maxi:
					a2++
				case sum-maxi < maxi:
					a3++
				}
			}
		}
	}
	fmt.Println(a1, a2, a3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

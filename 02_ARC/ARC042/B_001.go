package main

import (
	"fmt"
	"math"
)

func main() {
	var X, Y float64
	var n int
	fmt.Scan(&X, &Y, &n)

	x := make([]float64, n+1)
	y := make([]float64, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}
	x[n] = x[0]
	y[n] = y[0]

	dist := 1e10
	for i := 0; i < n; i++ {
		a := y[i+1] - y[i]
		b := x[i] - x[i+1]
		c := x[i]*(y[i]-y[i+1]) + y[i]*(x[i+1]-x[i])
		d := math.Abs(a*X+b*Y+c) / math.Hypot(a, b)
		dist = math.Min(dist, d)
	}

	fmt.Println(dist)
}

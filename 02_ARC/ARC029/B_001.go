package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var n int
	fmt.Scan(&a, &b, &n)

	r := math.Sqrt(a*a + b*b)
	if a > b {
		a, b = b, a
	}

	const EPS = 1e-8
	for i := 0; i < n; i++ {
		var c, d float64
		fmt.Scan(&c, &d)
		if c > d {
			c, d = d, c
		}
		if a < c+EPS && b < d+EPS || a < c+EPS && c < r+EPS && math.Cos(math.Asin(c/r)-math.Atan(a/b))*b*2-math.Sqrt(r*r-c*c) < d {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

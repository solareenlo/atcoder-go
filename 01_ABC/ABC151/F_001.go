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
	cx, cy := 0.0, 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
		cx += x[i]
		cy += y[i]
	}

	cx /= float64(n)
	cy /= float64(n)

	res := 0.0
	for p := 1.0; p >= 1e-8; p *= 0.999 {
		idx := -1
		dist := -1.0
		for i := 0; i < n; i++ {
			d := math.Hypot(x[i]-cx, y[i]-cy)
			if dist < d {
				dist = d
				idx = i
			}
		}
		cx += (x[idx] - cx) * p
		cy += (y[idx] - cy) * p
		res = dist
	}

	fmt.Println(res)
}

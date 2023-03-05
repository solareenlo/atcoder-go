package main

import (
	"fmt"
	"math"
)

func main() {
	var x0, y0, vx, vy, vh float64
	fmt.Scan(&x0, &y0, &vx, &vy, &vh)

	var solve func(float64, float64, float64) []float64
	solve = func(a, b, c float64) []float64 {
		if math.Abs(a) < 1e-10 && math.Abs(b) < 1e-10 {
			return []float64{}
		} else if math.Abs(a) < 1e-10 {
			return []float64{-c / b}
		}
		D := b*b - 4.0*a*c
		if D < 0 {
			return []float64{}
		}
		return []float64{(-b + math.Sqrt(D)) / (2 * a), (-b - math.Sqrt(D)) / (2 * a)}
	}

	ans := solve(vx*vx+vy*vy-vh*vh, 2*(vx*x0+vy*y0), x0*x0+y0*y0)
	t := -1.0
	for _, a := range ans {
		if a >= 0 {
			if t == -1 || t > a {
				t = a
			}
		}
	}
	if t < 0 {
		fmt.Println("IMPOSSIBLE")
	} else {
		fmt.Println(t)
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	l, r := 0.0, 2000.0
	for r-l > 1e-12 {
		m := (r + l) / 2
		if a*m+b*math.Sin(c*m*math.Pi) < 100 {
			l = m
		} else {
			r = m
		}
	}
	fmt.Printf("%.18f\n", l)
}

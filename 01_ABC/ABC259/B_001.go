package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, cta float64
	fmt.Scan(&a, &b, &cta)

	cta *= math.Acos(-1) / 180
	fmt.Printf("%.9f %.9f\n", math.Cos(cta)*a-math.Sin(cta)*b, math.Cos(cta)*b+math.Sin(cta)*a)
}

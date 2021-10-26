package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, x float64
	fmt.Scan(&a, &b, &x)

	if x == a*a*b {
		fmt.Printf("%.10f\n", 0.0)
	} else if x < a*a*b/2.0 {
		t := 2.0 * x / b / a
		angle := math.Atan2(t, b) * 180.0 / math.Pi
		fmt.Printf("%.10f\n", 90.0-angle)
	} else if x == a*a*b/2.0 {
		fmt.Printf("%.10f\n", 45.0)
	} else {
		t := (a*a*b - x) * 2.0 / a / a
		angle := math.Atan2(t, a) * 180.0 / math.Pi
		fmt.Printf("%.10f\n", angle)
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	fmt.Printf("%.3f\n", math.Round(b*1000/a)/1000)
}

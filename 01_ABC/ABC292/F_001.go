package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a > b {
		a, b = b, a
	}
	if b >= 2*a/math.Sqrt(3) {
		fmt.Println(2 * a / math.Sqrt(3))
	} else {
		fmt.Println(b / math.Cos(math.Atan(2*a/b-math.Sqrt(3))))
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		sum += a * b
	}

	fmt.Println(math.Floor(float64(sum) * 1.05))
}

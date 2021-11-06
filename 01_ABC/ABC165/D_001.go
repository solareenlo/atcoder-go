package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	var x int
	if n < b-1 {
		x = n
	} else {
		x = b - 1
	}

	fmt.Println(math.Floor(float64(a * (x % b) / b)))
}

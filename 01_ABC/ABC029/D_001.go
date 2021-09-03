package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i := 10; i <= int(1e9); i *= 10 {
		x := i / 10
		res += n/i*x + int(math.Min(math.Max(0.0, float64(n%i-x+1)), float64(x)))
	}
	fmt.Println(res)
}

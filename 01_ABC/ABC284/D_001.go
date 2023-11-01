package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var n, x int64
		fmt.Scan(&n)

		for x = 2; n%x != 0; x++ {
		}

		if n%(x*x) == 0 {
			fmt.Printf("%d %d\n", x, n/(x*x))
		} else {
			fmt.Printf("%d %d\n", int64(math.Sqrt(float64(n/x))), x)
		}
	}
}

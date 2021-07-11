package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a int
	fmt.Scan(&n)
	res := 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		res += math.Max(0, float64(a-10))
	}
	fmt.Println(res)
}

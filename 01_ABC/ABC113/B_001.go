package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var t, a float64
	fmt.Scan(&n, &t, &a)

	var index int
	h, mini := 0.0, 1e16
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		w := math.Abs(a - (t - h*0.006))
		if w < mini {
			index = i
			mini = w
		}
	}
	fmt.Println(index + 1)
}

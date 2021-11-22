package main

import (
	"fmt"
	"math"
)

func main() {
	var r, x, y float64
	fmt.Scan(&r, &x, &y)

	d := math.Sqrt(x*x + y*y)
	if d < r {
		fmt.Println(2)
	} else {
		fmt.Println(math.Ceil(d / r))
	}
}

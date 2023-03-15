package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scan(&x)
	t := int(math.Sqrt(float64(x) * 2))
	if t*(t+1) == x*2 {
		fmt.Println(t)
	} else {
		fmt.Println(-1)
	}
}

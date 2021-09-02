package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, e float64
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Println(math.Max(a+d+e, b+c+e))
}

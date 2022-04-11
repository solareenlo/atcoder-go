package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	c = math.Sqrt(a*a + b*b)
	fmt.Println(a/c, b/c)
}

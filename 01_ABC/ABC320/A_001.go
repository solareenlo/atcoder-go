package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(int(math.Pow(a, b)) + int(math.Pow(b, a)))
}

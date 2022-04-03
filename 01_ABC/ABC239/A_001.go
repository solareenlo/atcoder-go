package main

import (
	"fmt"
	"math"
)

func main() {
	var h int
	fmt.Scan(&h)

	fmt.Println(math.Sqrt(float64(h * (12800000 + h))))
}

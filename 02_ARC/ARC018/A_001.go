package main

import (
	"fmt"
	"math"
)

func main() {
	var h, bmi float64
	fmt.Scan(&h, &bmi)

	h /= 100.0
	fmt.Println(bmi * math.Pow(h, 2))
}

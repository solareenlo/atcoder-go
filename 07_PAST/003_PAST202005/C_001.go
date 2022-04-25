package main

import (
	"fmt"
	"math"
)

func main() {
	var a, r, n float64
	fmt.Scan(&a, &r, &n)

	if math.Log10(a)+(n-1)*math.Log10(r) > 9 {
		fmt.Println("large")
	} else {
		fmt.Println(int(a * math.Pow(r, n-1)))
	}
}

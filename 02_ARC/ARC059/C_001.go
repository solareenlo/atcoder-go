package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	v := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
		v += a * a
	}

	S := int(math.Round(float64(sum) / float64(n)))
	fmt.Println(v - 2*S*sum + S*S*n)
}

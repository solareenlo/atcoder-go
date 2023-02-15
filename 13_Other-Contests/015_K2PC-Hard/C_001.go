package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := int(math.Sqrt(float64(n))) - 1
	if n == 2 {
		fmt.Println(0)
	} else {
		if s*s+s+s+1 != n {
			fmt.Println(-1)
		} else {
			fmt.Println((s*s*s-s)/3 + n - 1)
		}
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var d int
	fmt.Scan(&d)
	ans := int(1e9)
	for i := 1; i*i <= d; i++ {
		x := int(math.Sqrt(float64(d - i*i)))
		ans = min(ans, d-i*i-x*x)
		ans = min(ans, abs(d-i*i-(x+1)*(x+1)))
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

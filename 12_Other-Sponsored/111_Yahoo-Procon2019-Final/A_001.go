package main

import (
	"fmt"
	"math"
)

func main() {
	var h, w, a, b float64
	fmt.Scan(&h, &w, &a, &b)
	t := math.Min(a, h-a)
	ans := 2 * t * (h*a - a*a + t*(t/3-h/2)) / (h - a) / (h - a)
	t = math.Min(b, w-b)
	ans *= 2 * t * (w*b - b*b + t*(t/3-w/2)) / (w - b) / (w - b)
	fmt.Println(ans)
}

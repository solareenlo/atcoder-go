package main

import (
	"fmt"
	"math"
)

func main() {
	var xa, ya, xb, yb, t, v, n int
	fmt.Scan(&xa, &ya, &xb, &yb, &t, &v, &n)

	maxDist := float64(t * v)
	var x, y int
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		dist := math.Hypot(float64(x-xa), float64(y-ya)) + math.Hypot(float64(x-xb), float64(y-yb))
		if dist <= maxDist {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

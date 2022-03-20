package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, r1, x2, y2, r2 float64
	fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2)

	dist := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	res := 5
	if dist < math.Abs(r1-r2) {
		res = 1
	} else if dist == math.Abs(r1-r2) {
		res = 2
	} else if dist < r1+r2 {
		res = 3
	} else if dist == r1+r2 {
		res = 4
	}
	fmt.Println(res)
}

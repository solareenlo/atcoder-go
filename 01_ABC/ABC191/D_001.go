package main

import (
	"fmt"
	"math"
)

func main() {
	var X, Y, R float64
	fmt.Scan(&X, &Y, &R)

	const m = 1e4
	x := int(math.Round(X*m))%m + 2e9
	y := int(math.Round(Y*m))%m + 2e9
	r := int(math.Round(R * m))

	res := 0
	x0, x1 := ((x-r-1)/m+1)*m, (x+r)/m*m
	for i := x0; i <= x1; i += m {
		d := r*r - (x-i)*(x-i)
		dy := int(math.Sqrt(float64(d)))
		for dy*dy > d {
			dy--
		}
		res += (y+dy)/m - (y-dy-1)/m
	}
	fmt.Println(res)
}

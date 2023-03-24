package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var h, w float64
		fmt.Scan(&h, &w)
		if w < h {
			w, h = h, w
		}
		var res float64
		if w/h < 2/math.Sqrt(3) {
			res = 2 * math.Sqrt(h*h-math.Sqrt(3)*w*h+w*w)
		} else {
			res = math.Sqrt(h*h + w*w/4.0)
		}
		fmt.Printf("%.10f\n", res)
		t--
	}
}

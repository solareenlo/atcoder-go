package main

import (
	"fmt"
	"math"
)

func main() {
	var h, w float64
	fmt.Scan(&h, &w)
	if h >= w {
		h, w = w, h
	}
	fmt.Println(math.Min(h/2, (h+w-math.Sqrt(2*h*w))/2))
}
